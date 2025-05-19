package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id" db:"id" example:"550e8400-e29b-41d4-a716-446655440000"`
	TelegramID int64     `json:"telegram_id" db:"telegram_id" example:"123456789"`
	Username   string    `json:"username" db:"username" example:"johndoe"`
	CreatedAt  time.Time `json:"created_at" db:"created_at" example:"2025-04-01T12:00:00Z"`
}
