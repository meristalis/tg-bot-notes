// Package repo implements application outer layer logic. Each logic group in own file.
package repo

import (
	"context"

	"github.com/meristalis/tg-bot-notes/internal/entity"
)

//go:generate mockgen -source=contracts.go -destination=../usecase/mocks_repo_test.go -package=usecase_test

type (
	// TranslationRepo -.
	TranslationRepo interface {
		Store(context.Context, entity.Translation) error
		GetHistory(context.Context) ([]entity.Translation, error)
	}

	// TranslationWebAPI -.
	TranslationWebAPI interface {
		Translate(entity.Translation) (entity.Translation, error)
	}

	// NoteRepo
	NoteRepo interface {
		Store(context.Context, entity.Note) error
		GetAllNotes(context.Context) ([]entity.Note, error)
	}
)
