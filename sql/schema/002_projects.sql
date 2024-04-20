-- +goose Up
CREATE TABLE projects(
    id UUID PRIMARY KEY,
    name TEXT UNIQUE NOT NULL,
    content JSON NOT NULL,
    description TEXT DEFAULT NULL
);

-- +goose Down
DROP TABLE projects;