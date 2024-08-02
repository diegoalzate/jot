-- name: CreateTask :one
INSERT INTO tasks (id, name, description, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, name, description, created_at, updated_at;