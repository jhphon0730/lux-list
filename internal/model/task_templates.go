package database

import "time"

type TaskTemplate struct {
	ID          int        `db:"id"`
	UserID      int        `db:"user_id"`
	Title       string     `db:"title"`
	Description *string    `db:"description"`
	RepeatType  string     `db:"repeat_type"`
	RepeatDays  *string    `db:"repeat_days"`
	StartDate   time.Time  `db:"start_date"`
	EndDate     *time.Time `db:"end_date"`
	CreatedAt   time.Time  `db:"created_at"`
	UpdatedAt   time.Time  `db:"updated_at"`
}
