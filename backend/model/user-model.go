package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID uuid.UUID `json:"user_id"`
	Username string `json:"username"`
	Password string `json:"password"`

	CreatedAt *time.Time	`json:"created_at"`
}


