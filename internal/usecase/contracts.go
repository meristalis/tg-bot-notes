// Package usecase implements application business logic. Each logic group in own file.
package usecase

import (
	"context"

	"github.com/meristalis/tg-bot-notes/internal/entity"
)

//go:generate mockgen -source=interfaces.go -destination=./mocks_usecase_test.go -package=usecase_test

type (
	// Translation -.
	Translation interface {
		Translate(context.Context, entity.Translation) (entity.Translation, error)
		History(context.Context) ([]entity.Translation, error)
	}
	Note interface {
		GetAllNotes(context.Context) ([]entity.Note, error)
		AddNote(context.Context, entity.Note) (entity.Note, error)
	}
)
