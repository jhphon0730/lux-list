package repository

import (
	"database/sql"

	"lux-list/internal/model"
	"lux-list/pkg/utils"

	sq "github.com/Masterminds/squirrel"
)

const (
	// Query
	FIND_ALL_TASKS_QUERY            = "SELECT id, user_id, title, description, due_date, is_completed, priority, created_at, updated_at FROM tasks WHERE user_id = $1 ORDER BY due_date DESC"
	FIND_ALL_TASKS_QUERY_BY_TASK_ID = "SELECT id, user_id, title, description, due_date, is_completed, priority, created_at, updated_at FROM tasks WHERE id = $1 AND user_id = $2"
	INSERT_TASKS_QUERY              = "INSERT INTO tasks (user_id, title, description, due_date, is_completed, priority) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	DELETE_TASKS_QUERY              = "DELETE FROM tasks WHERE id = $1 AND user_id = $2"
	UPDATE_TASKS_QUERY              = "UPDATE tasks SET title = $1, description = $2, due_date = $3, is_completed = $4, priority = $5, updated_at = NOW() WHERE id = $6 AND user_id = $7 RETURNING updated_at"
)

// TaskRepository는 작업 관련 데이터베이스 작업을 정의하는 인터페이스
type TaskRepository interface {
	GetTasks(userID int, search_query map[string]interface{}) (*model.TaskListResult, error)
	GetTasksByTaskID(userID int, taskID int) (*model.Task, error)
	CreateTasks(userID int, task *model.Task) (*model.Task, error)
	DeleteTasks(userID int, taskID int) error
	UpdateTasks(userID int, taskID int, task *model.Task) (*model.Task, error)
}

// taskRepository는 TaskRepository 인터페이스를 구현하는 구조체
type taskRepository struct {
	db *sql.DB
}

// NewTaskRepository는 TaskRepository의 인스턴스를 생성하는 함수
func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepository{
		db: db,
	}
}

// GetTasks는 사용자의 모든 작업을 조회하는 메서드
func (r *taskRepository) GetTasks(userID int, search_query map[string]interface{}) (*model.TaskListResult, error) {
	limit, page := utils.CreatePaginationQuery(search_query)
	orderBy := utils.CreateOrderByQuery(search_query)

	queryBuilder := sq.Select(
		"id", "user_id", "title", "description", "due_date", "is_completed", "priority", "created_at", "updated_at",
		"COUNT(*) OVER() AS total_count", // 전체 작업 수를 가져오기 위한 서브쿼리
	).
		From("tasks").
		Where(sq.Eq{"user_id": userID}).
		Limit(uint64(limit)).
		Offset(uint64((page - 1) * limit)).
		OrderBy(orderBy)

	// 검색 쿼리 처리
	for key, value := range search_query {
		switch key {
		case "title":
			queryBuilder = queryBuilder.Where(sq.Like{"title": "%" + value.(string) + "%"})
		case "is_completed":
			queryBuilder = queryBuilder.Where(sq.Eq{"is_completed": value})
		case "priority":
			queryBuilder = queryBuilder.Where(sq.Eq{"priority": value.(string)})
		case "due_date":
			queryBuilder = queryBuilder.Where(sq.Eq{"due_date": value.(string)})
		}
	}

	query, args, err := queryBuilder.PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	totalCount := 0
	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.DueDate, &task.IsCompleted, &task.Priority, &task.CreatedAt, &task.UpdatedAt, &totalCount); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return &model.TaskListResult{
		Tasks:      tasks,
		TotalCount: totalCount,
	}, nil
}

// GetTasksByTaskID는 작업 ID로 작업을 조회하는 메서드
func (r *taskRepository) GetTasksByTaskID(userID int, taskID int) (*model.Task, error) {
	var task model.Task
	query := FIND_ALL_TASKS_QUERY_BY_TASK_ID
	row := r.db.QueryRow(query, taskID, userID)
	if err := row.Scan(&task.ID, &task.UserID, &task.Title, &task.Description, &task.DueDate, &task.IsCompleted, &task.Priority, &task.CreatedAt, &task.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	return &task, nil
}

// CreateTasks는 새로운 작업을 생성하는 메서드
func (r *taskRepository) CreateTasks(userID int, task *model.Task) (*model.Task, error) {
	query := INSERT_TASKS_QUERY
	row := r.db.QueryRow(query, userID, task.Title, task.Description, task.DueDate, task.IsCompleted, task.Priority)
	if err := row.Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}
	task.UserID = userID

	return task, nil
}

// DeleteTask는 작업을 삭제하는 메서드
func (r *taskRepository) DeleteTasks(userID int, taskID int) error {
	query := DELETE_TASKS_QUERY
	result, err := r.db.Exec(query, taskID, userID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows
	}

	return nil
}

// UpdateTasks는 작업을 업데이트하는 메서드
func (r *taskRepository) UpdateTasks(userID int, taskID int, task *model.Task) (*model.Task, error) {
	query := UPDATE_TASKS_QUERY
	row := r.db.QueryRow(query, task.Title, task.Description, task.DueDate, task.IsCompleted, task.Priority, taskID, userID)
	if err := row.Scan(&task.UpdatedAt); err != nil {
		return nil, err
	}
	task.UserID = userID

	return task, nil
}
