package model

import (
	"errors"
	"time"
)

type Task struct {
	ID          int       `db:"id"`
	TemplateID  *int      `db:"template_id"`
	UserID      int       `db:"user_id"`
	Title       string    `db:"title"`
	Description *string   `db:"description"`
	DueDate     time.Time `db:"due_date"`
	IsCompleted bool      `db:"is_completed"`
	Priority    string    `db:"priority"` // "low", "medium", "high"
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`

	Tags []Tag `db:"-" json:"tags"` // 태그는 Task와 N:M 관계를 가짐
}

type CreateTaskRequest struct {
	Title       string    `json:"title"`
	Description *string   `json:"description"`
	DueDate     time.Time `json:"due_date"`
	IsCompleted bool      `json:"is_completed"`
	Priority    string    `json:"priority"` // "low", "medium", "high"
}

// CheckValidCreateTaskRequest는 CreateTaskRequest의 유효성을 검사하는 메서드
func (r CreateTaskRequest) CheckValidCreateTaskRequest() error {
	if r.Title == "" {
		return errors.New("title is required")
	}
	if r.DueDate.IsZero() {
		return errors.New("due date is required")
	}
	if r.Priority != "low" && r.Priority != "medium" && r.Priority != "high" {
		return errors.New("priority must be 'low', 'medium', or 'high'")
	}
	return nil
}

// ToTask는 CreateTaskRequest를 Task로 변환하는 메서드
func (r CreateTaskRequest) ToTask(userID int) *Task {
	return &Task{
		UserID:      userID,
		Title:       r.Title,
		Description: r.Description,
		DueDate:     r.DueDate,
		IsCompleted: r.IsCompleted,
		Priority:    r.Priority,
	}
}

// toTaskTemplate는 CreateTaskRequest를 Task로 변환하는 메서드
func (r CreateTaskRequest) ToTaskTemplate(templateID int, userID int) *Task {
	return &Task{
		TemplateID:  &templateID,
		UserID:      userID,
		Title:       r.Title,
		Description: r.Description,
		DueDate:     r.DueDate,
		IsCompleted: r.IsCompleted,
		Priority:    r.Priority,
	}
}
