-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- Вставка пользователей
INSERT INTO users (telegram_id, username)
VALUES 
    (123456789, 'user1'),
    (987654321, 'user2'),
    (555555555, 'user3');

-- Вставка тегов
INSERT INTO tags (name)
VALUES 
    ('Important'),
    ('Work'),
    ('Personal'),
    ('Urgent');

-- Вставка заметок
INSERT INTO notes (user_id, title, content)
VALUES 
    ((SELECT id FROM users WHERE telegram_id = 123456789), 'My first note', 'This is the content of the first note'),
    ((SELECT id FROM users WHERE telegram_id = 123456789), 'Important note', 'This is a note with important information'),
    ((SELECT id FROM users WHERE telegram_id = 987654321), 'Work tasks', 'List of tasks for work'),
    ((SELECT id FROM users WHERE telegram_id = 555555555), 'Personal note', 'Personal thoughts and ideas');

-- Связь заметок с тегами
INSERT INTO note_tags (note_id, tag_id)
VALUES 
    ((SELECT id FROM notes WHERE title = 'My first note'), (SELECT id FROM tags WHERE name = 'Personal')),
    ((SELECT id FROM notes WHERE title = 'Important note'), (SELECT id FROM tags WHERE name = 'Important')),
    ((SELECT id FROM notes WHERE title = 'Work tasks'), (SELECT id FROM tags WHERE name = 'Work')),
    ((SELECT id FROM notes WHERE title = 'Personal note'), (SELECT id FROM tags WHERE name = 'Personal')),
    ((SELECT id FROM notes WHERE title = 'Work tasks'), (SELECT id FROM tags WHERE name = 'Urgent'));

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
