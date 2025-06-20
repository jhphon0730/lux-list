package repository

import (
	"database/sql"
	"lux-list/internal/model"
)

const (
	EXIST_TAG_IN_TASK_QUERY    = "SELECT EXISTS(SELECT 1 FROM task_tags WHERE task_id = $1 AND tag_id = $2)"
	ADD_TAG_TO_TASK_QUERY      = "INSERT INTO task_tags (task_id, tag_id) VALUES ($1, $2)"
	REMOVE_TAG_FROM_TASK_QUERY = "DELETE FROM task_tags WHERE task_id = $1 AND tag_id = $2"
)

// TaskTagRepository는 작업과 태그 간의 관계를 관리하는 인터페이스
type TaskTagRepository interface {
	AddTagToTask(taskID int, tagID int) error
	RemoveTagFromTask(taskID int, tagID int) error
	GetTagsByTaskID(taskID int) ([]model.Tag, error)
}

// taskTagRepository는 TaskTagRepository 인터페이스를 구현하는 구조체
type taskTagRepository struct {
	db *sql.DB
}

// NewTaskTagRepository는 TaskTagRepository의 인스턴스를 생성하는 함수
func NewTaskTagRepository(db *sql.DB) TaskTagRepository {
	return &taskTagRepository{
		db: db,
	}
}

// AddTagToTask는 작업에 태그를 추가하는 메서드
func (r *taskTagRepository) AddTagToTask(taskID int, tagID int) error {
	// 이미 연결되어 있는지 확인하는 쿼리 실행
	var exists bool
	err := r.db.QueryRow(EXIST_TAG_IN_TASK_QUERY, taskID, tagID).Scan(&exists)
	if err != nil {
		return err
	}

	// 이미 연결되어 있다면 sql.ErrNoRows 대신 의미 있는 에러를 반환
	if exists {
		return sql.ErrTxDone
	}

	// 연결되어 있지 않으면 태그 추가
	_, err = r.db.Exec(ADD_TAG_TO_TASK_QUERY, taskID, tagID)
	if err != nil {
		return err
	}

	return nil
}

// RemoveTagFromTask는 작업에서 태그를 제거하는 메서드
func (r *taskTagRepository) RemoveTagFromTask(taskID int, tagID int) error {
	result, err := r.db.Exec(REMOVE_TAG_FROM_TASK_QUERY, taskID, tagID)
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

// GetTagsByTaskID는 특정 작업에 연결된 태그를 조회하는 메서드
func (r *taskTagRepository) GetTagsByTaskID(taskID int) ([]model.Tag, error) {
	rows, err := r.db.Query(GET_TAGS_BY_TASK_ID_QUERY, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []model.Tag
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.ID, &tag.UserID, &tag.Name, &tag.Color, &tag.CreatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}
