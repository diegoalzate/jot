-- name: CreateTask :one
INSERT INTO tasks (id, name, description, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING id, name, description, created_at, updated_at;

-- name: ListTasks :many
SELECT id, name, description, created_at, updated_at
FROM tasks
LIMIT 100;