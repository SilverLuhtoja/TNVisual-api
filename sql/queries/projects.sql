-- name: GetProjectsByName :one
SELECT * FROM projects
WHERE name = $1;

-- name: GetProjects :many
Select * from projects;

-- name: CreateProjects :one
INSERT INTO projects (id, name, content, description)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteProject :exec
DELETE FROM projects
WHERE name = $1;
