// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package query

import (
	"context"
	"database/sql"
	"time"
)

const createTask = `-- name: CreateTask :one
INSERT INTO tasks (id, name, description, created_at, updated_at)
VALUES (?, ?, ?, ?, ?)
RETURNING id, name, description, created_at, updated_at
`

type CreateTaskParams struct {
	ID          string
	Name        string
	Description sql.NullString
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (q *Queries) CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error) {
	row := q.db.QueryRowContext(ctx, createTask,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i Task
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
