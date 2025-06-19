package repository

import (
	"database/sql"
	"lux-list/internal/model"
)

const (
	GET_TAGS_QUERY            = "SELECT id, name, color, created_at FROM tags WHERE user_id = $1 ORDER BY created_at DESC"
	GET_TAGS_BY_TAG_ID_QUERY  = "SELECT id, name, color, created_at FROM tags WHERE user_id = $1 AND id = $2"
	GET_TAGS_BY_USER_ID_QUERY = "SELECT id, name, color, created_at FROM tags WHERE user_id = $1"
	GET_TAGS_BY_TASK_ID_QUERY = "SELECT id, name, color, created_at FROM tags WHERE id IN (SELECT tag_id FROM task_tags WHERE task_id = $1)"
)

// TagRepository는 태그 관련 데이터베이스 작업을 정의하는 인터페이스
type TagRepository interface {
	GetTags(userID int) ([]model.Tag, error)
	GetTagsByTagID(userID int, tagID int) (*model.Tag, error)
	GetTagsByUserID(userID int) ([]model.Tag, error)
	GetTagsByTaskID(userID int, taskID int) ([]model.Tag, error)
}

// tagRepository는 TagRepository 인터페이스를 구현하는 구조체
type tagRepository struct {
	db *sql.DB
}

// NewTagRepository는 TagRepository의 인스턴스를 생성하는 함수
func NewTagRepository(db *sql.DB) TagRepository {
	return &tagRepository{
		db: db,
	}
}

// GetTags는 사용자의 모든 태그를 조회하는 메서드
func (r *tagRepository) GetTags(userID int) ([]model.Tag, error) {
	rows, err := r.db.Query(GET_TAGS_QUERY, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []model.Tag
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Color, &tag.CreatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// GetTagsByUserID는 사용자의 모든 태그를 조회하는 메서드
func (r *tagRepository) GetTagsByTagID(userID int, tagID int) (*model.Tag, error) {
	row := r.db.QueryRow(GET_TAGS_BY_TAG_ID_QUERY, userID, tagID)
	var tag model.Tag
	if err := row.Scan(&tag.ID, &tag.Name, &tag.Color, &tag.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &tag, nil
}

// GetTagsByUserID는 사용자의 모든 태그를 조회하는 메서드
func (r *tagRepository) GetTagsByUserID(userID int) ([]model.Tag, error) {
	rows, err := r.db.Query(GET_TAGS_BY_USER_ID_QUERY, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []model.Tag
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Color, &tag.CreatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}

// GetTagsByTaskID는 작업 ID로 태그를 조회하는 메서드
func (r *tagRepository) GetTagsByTaskID(userID int, taskID int) ([]model.Tag, error) {
	rows, err := r.db.Query(GET_TAGS_BY_TASK_ID_QUERY, taskID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []model.Tag
	for rows.Next() {
		var tag model.Tag
		if err := rows.Scan(&tag.ID, &tag.Name, &tag.Color, &tag.CreatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}
	return tags, nil
}
