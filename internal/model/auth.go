package model

import "time"

type User struct {
	ID        int       `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
}

type LoginRequest struct {
	Name string `json:"name"`
}
