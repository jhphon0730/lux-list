package repository

import (
	"database/sql"
	"lux-list/internal/model"
)

// TaskRepository는 작업 관련 데이터베이스 작업을 정의하는 인터페이스
type TaskRepository interface {
	GetTasks(userID int) ([]model.Task, error)
	GetTaskByID(taskID int) (*model.Task, error)
	CreateTask(userID int, task *model.Task) (*model.Task, error)
	DeleteTask(userID int, taskID int) error
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
func (r *taskRepository) GetTasks(userID int) ([]model.Task, error) {
	var tasks []model.Task
	query := "SELECT id, title, description, due_date, is_completed, priority, created_at, updated_at FROM tasks WHERE user_id = $1 ORDER BY due_date DESC"
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var task model.Task
		if err := rows.Scan(&task.ID, &task.Title, &task.Description, &task.DueDate, &task.IsCompleted, &task.Priority, &task.CreatedAt, &task.UpdatedAt); err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}

	return tasks, nil
}

// GetTaskByID는 작업 ID로 작업을 조회하는 메서드
func (r *taskRepository) GetTaskByID(taskID int) (*model.Task, error) {
	var task model.Task
	return &task, nil
}

// CreateTask는 새로운 작업을 생성하는 메서드
func (r *taskRepository) CreateTask(userID int, task *model.Task) (*model.Task, error) {
	query := "INSERT INTO tasks (user_id, title, description, due_date, is_completed, priority) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, created_at, updated_at"
	row := r.db.QueryRow(query, userID, task.Title, task.Description, task.DueDate, task.IsCompleted, task.Priority)
	if err := row.Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt); err != nil {
		return nil, err
	}
	task.UserID = userID

	return task, nil
}

// DeleteTask는 작업을 삭제하는 메서드
func (r *taskRepository) DeleteTask(userID int, taskID int) error {
	query := "DELETE FROM tasks WHERE id = $1 AND user_id = $2"
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
