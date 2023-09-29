package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	PostID uuid.UUID `json:"post_id"`
	AuthorID uuid.UUID `json:"author_id"`
	PostTitle string `json:"post_title"`
	PostDescription string	`json:"post_description"`

	CreatedAt *time.Time	`json:"created_at"`
}


