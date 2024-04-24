-- name: GetUserByKey :one
SELECT * FROM users
WHERE api_key = $1;

-- name: AuthenticateUser :one
Select * from users
WHERE username = $1;

-- name: CreateUser :one
INSERT INTO users (id, username, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: UpdateUserKey :exec
UPDATE users SET api_key = $2
WHERE id = $1;