-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE TABLE IF NOT EXISTS history(
    id serial PRIMARY KEY,
    source VARCHAR(255),
    destination VARCHAR(255),
    original VARCHAR(255),
    translation VARCHAR(255)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS history;
-- +goose StatementEnd