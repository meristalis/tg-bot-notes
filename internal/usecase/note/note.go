package note

import (
	"context"
	"fmt"

	"github.com/meristalis/tg-bot-notes/internal/entity"
	"github.com/meristalis/tg-bot-notes/internal/repo"
)

// UseCase - структура для бизнес-логики работы с заметками.
type UseCase struct {
	repo repo.NoteRepo
}

// New - конструктор для UseCase.
func New(r repo.NoteRepo) *UseCase {
	return &UseCase{
		repo: r,
	}
}

// GetAllNotes - получение всех заметок.
func (uc *UseCase) GetAllNotes(ctx context.Context) ([]entity.Note, error) {
	notes, err := uc.repo.GetAllNotes(ctx)
	if err != nil {
		return nil, fmt.Errorf("NoteUseCase - GetAllNotes - s.repo.GetAllNotes: %w", err)
	}

	return notes, nil
}

// AddNote - добавление новой заметки.
func (uc *UseCase) AddNote(ctx context.Context, note entity.Note) (entity.Note, error) {

	// Сохранение новой заметки в базе данных
	err := uc.repo.Store(ctx, note)
	if err != nil {
		return entity.Note{}, fmt.Errorf("NoteUseCase - AddNote - s.repo.Store: %w", err)
	}

	return note, nil
}
