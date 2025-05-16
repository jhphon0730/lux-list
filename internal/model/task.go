package model

import "time"

type Task struct {
	ID          int       `db:"id"`
	TemplateID  *int      `db:"template_id"`
	UserID      int       `db:"user_id"`
	Title       string    `db:"title"`
	Description *string   `db:"description"`
	DueDate     time.Time `db:"due_date"`
	IsCompleted bool      `db:"is_completed"`
	Priority    string    `db:"priority"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
