package persistent

import (
	"context"
	"fmt"

	"github.com/meristalis/tg-bot-notes/internal/entity"
	"github.com/meristalis/tg-bot-notes/pkg/postgres"
)

// NoteRepo - репозиторий для работы с заметками.
type NoteRepo struct {
	*postgres.Postgres
}

// NewNoteRepo - конструктор для NoteRepo.
func NewNoteRepo(pg *postgres.Postgres) *NoteRepo {
	return &NoteRepo{pg}
}

// GetAllNotes - получение всех заметок.
func (r *NoteRepo) GetAllNotes(ctx context.Context) ([]entity.Note, error) {
	sql, _, err := r.Builder.
		Select("id, user_id, title, content, created_at, updated_at").
		From("notes").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("NoteRepo - GetAllNotes - r.Builder: %w", err)
	}

	rows, err := r.Pool.Query(ctx, sql)
	if err != nil {
		return nil, fmt.Errorf("NoteRepo - GetAllNotes - r.Pool.Query: %w", err)
	}
	defer rows.Close()

	entities := make([]entity.Note, 0, _defaultEntityCap)

	for rows.Next() {
		e := entity.Note{}

		err = rows.Scan(&e.ID, &e.UserID, &e.Title, &e.Content, &e.CreatedAt, &e.UpdatedAt)
		if err != nil {
			return nil, fmt.Errorf("NoteRepo - GetAllNotes - rows.Scan: %w", err)
		}

		entities = append(entities, e)
	}

	return entities, nil
}
func (r *NoteRepo) Store(ctx context.Context, note entity.Note) error {
	// Формирование SQL-запроса на вставку новой заметки.
	sql, args, err := r.Builder.
		Insert("notes").
		Columns("user_id, title, content").
		Values(note.UserID, note.Title, note.Content).
		ToSql()
	if err != nil {
		return fmt.Errorf("NoteRepo - Store - r.Builder: %w", err)
	}

	// Выполнение SQL-запроса.
	_, err = r.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("NoteRepo - Store - r.Pool.Exec: %w", err)
	}

	return nil
}
