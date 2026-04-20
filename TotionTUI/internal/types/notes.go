package types

import "time"

type Notes struct {
	Id          int64     `json:"id"`
	UserId      int64     `json:"user_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type NotesRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
