package models

import (
	"time"
)

type Post struct {
	ID             int       `json:"id"`
	Title          string    `json:"title"`
	Content        string    `json:"content"`
	AuthorID       int       `json:"author_id"`
	AuthorUsername string    `json:"author_username"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// TableName specifies the table name for the Post model
func (Post) TableName() string {
	return "posts"
} 