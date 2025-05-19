package entity

import (
	"time"

	"github.com/google/uuid"
)

type Note struct {
	ID        uuid.UUID `json:"id" db:"id" example:"550e8400-e29b-41d4-a716-446655440001"`
	UserID    uuid.UUID `json:"user_id" db:"user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	Title     string    `json:"title" db:"title" example:"My First Note"`
	Content   string    `json:"content" db:"content" example:"This is the content of the note."`
	CreatedAt time.Time `json:"created_at" db:"created_at" example:"2025-04-01T12:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at" example:"2025-04-01T12:00:00Z"`
}
