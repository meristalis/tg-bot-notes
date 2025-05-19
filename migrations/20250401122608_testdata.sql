-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Вставка тегов
INSERT INTO tags (name)
VALUES 
    ('Important'),
    ('Work'),
    ('Personal'),
    ('Urgent');

-- Вставка заметок с возвратом ID
WITH 
note1 AS (
    INSERT INTO notes (user_id, title, content)
    VALUES ('94f2e7c9-5e25-4dcf-9b12-1d2a837b100b', 'My first note', 'This is the content of the first note')
    RETURNING id
),
note2 AS (
    INSERT INTO notes (user_id, title, content)
    VALUES ('94f2e7c9-5e25-4dcf-9b12-1d2a837b100b', 'Important note', 'This is a note with important information')
    RETURNING id
),
note3 AS (
    INSERT INTO notes (user_id, title, content)
    VALUES ('94f2e7c9-5e25-4dcf-9b12-1d2a837b100b', 'Work tasks', 'List of tasks for work')
    RETURNING id
),
note4 AS (
    INSERT INTO notes (user_id, title, content)
    VALUES ('94f2e7c9-5e25-4dcf-9b12-1d2a837b100b', 'Personal note', 'Personal thoughts and ideas')
    RETURNING id
)

-- Связь заметок с тегами
INSERT INTO note_tags (note_id, tag_id)
SELECT note1.id, tags.id FROM note1, tags WHERE tags.name = 'Personal'
UNION ALL
SELECT note2.id, tags.id FROM note2, tags WHERE tags.name = 'Important'
UNION ALL
SELECT note3.id, tags.id FROM note3, tags WHERE tags.name = 'Work'
UNION ALL
SELECT note4.id, tags.id FROM note4, tags WHERE tags.name = 'Personal'
UNION ALL
SELECT note3.id, tags.id FROM note3, tags WHERE tags.name = 'Urgent';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Удаление связей note_tags
DELETE FROM note_tags
WHERE note_id IN (
    SELECT id FROM notes WHERE title IN (
        'My first note', 'Important note', 'Work tasks', 'Personal note'
    )
);

-- Удаление заметок
DELETE FROM notes
WHERE title IN ('My first note', 'Important note', 'Work tasks', 'Personal note');

-- Удаление тегов
DELETE FROM tags
WHERE name IN ('Important', 'Work', 'Personal', 'Urgent');

-- +goose StatementEnd
