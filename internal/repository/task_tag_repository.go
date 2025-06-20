package repository

import (
	"database/sql"
	"lux-list/internal/model"
)

const (
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
	_, err := r.db.Exec(ADD_TAG_TO_TASK_QUERY, taskID, tagID)
	return err
}

// RemoveTagFromTask는 작업에서 태그를 제거하는 메서드
func (r *taskTagRepository) RemoveTagFromTask(taskID int, tagID int) error {
	_, err := r.db.Exec(REMOVE_TAG_FROM_TASK_QUERY, taskID, tagID)
	return err
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
