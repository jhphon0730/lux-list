package repository

import (
	"database/sql"
	"lux-list/internal/model"
)

const (
	GET_TAGS_QUERY            = ""
	GET_TAGS_BY_TAG_ID_QUERY  = ""
	GET_TAGS_BY_USER_ID_QUERY = ""
	GET_TAGS_BY_TASK_ID_QUERY = ""
)

// TagRepository는 태그 관련 데이터베이스 작업을 정의하는 인터페이스
type TagRepository interface {
	GetTags(userID int) ([]*model.Tag, error)
	GetTagsByTagID(userID int, tagID int) (*model.Tag, error)
	GetTagsByUserID(userID int) ([]*model.Tag, error)
	GetTagsByTaskID(userID int, taskID int) ([]*model.Tag, error)
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
