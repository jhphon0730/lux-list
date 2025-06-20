package model

import (
	"errors"
	"time"
)

type Tag struct {
	ID        int       `db:"id"`
	UserID    int       `db:"user_id"`
	Name      string    `db:"name"`
	Color     string    `db:"color"`
	CreatedAt time.Time `db:"created_at"`
}

// CreateTagRequest는 태그 생성을 위한 요청 구조체입니다.
type CreateTagRequest struct {
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}

// CheckValidCreateTagRequest는 CreateTagRequest의 유효성을 검사하는 메서드
func (r *CreateTagRequest) CheckValidCreateTagRequest() error {
	if r.Name == "" {
		return errors.New("name is required")
	}
	if r.Color == "" {
		return errors.New("color is required")
	}
	if len(r.Color) != 7 || r.Color[0] != '#' {
		return errors.New("color must be a valid hex code (e.g., #FFFFFF)")
	}
	return nil
}

// ToTag는 CreateTagRequest를 Tag 모델로 변환하는 메서드
func (r *CreateTagRequest) ToTag(userID int) *Tag {
	return &Tag{
		UserID: userID,
		Name:   r.Name,
		Color:  r.Color,
	}
}
