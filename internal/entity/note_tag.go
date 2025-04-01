package entity

import "github.com/google/uuid"

type NoteTag struct {
	NoteID uuid.UUID `json:"note_id" db:"note_id" example:"550e8400-e29b-41d4-a716-446655440000"`
	TagID  int       `json:"tag_id" db:"tag_id" example:"1"`
}
