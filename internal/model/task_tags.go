package model

type TaskTag struct {
	TaskID int `db:"task_id"`
	TagID  int `db:"tag_id"`
}
