// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package query

import (
	"context"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

const createIdentity = `-- name: CreateIdentity :one
INSERT INTO identities (id, user_id, provider, provider_id, identity_data, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, provider, provider_id, identity_data, created_at, updated_at
`

type CreateIdentityParams struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Provider     string
	ProviderID   string
	IdentityData json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type CreateIdentityRow struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	Provider     string
	ProviderID   string
	IdentityData json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (q *Queries) CreateIdentity(ctx context.Context, arg CreateIdentityParams) (CreateIdentityRow, error) {
	row := q.db.QueryRowContext(ctx, createIdentity,
		arg.ID,
		arg.UserID,
		arg.Provider,
		arg.ProviderID,
		arg.IdentityData,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i CreateIdentityRow
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Provider,
		&i.ProviderID,
		&i.IdentityData,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createUser = `-- name: CreateUser :one
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, email, created_at, updated_at
`

type CreateUserParams struct {
	ID        uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.CreatedAt,
		arg.UpdatedAt,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getIdentity = `-- name: GetIdentity :one
SELECT id, user_id, provider
FROM identities
WHERE provider_id = $1 AND provider = $2
`

type GetIdentityParams struct {
	ProviderID string
	Provider   string
}

type GetIdentityRow struct {
	ID       uuid.UUID
	UserID   uuid.UUID
	Provider string
}

func (q *Queries) GetIdentity(ctx context.Context, arg GetIdentityParams) (GetIdentityRow, error) {
	row := q.db.QueryRowContext(ctx, getIdentity, arg.ProviderID, arg.Provider)
	var i GetIdentityRow
	err := row.Scan(&i.ID, &i.UserID, &i.Provider)
	return i, err
}

const getUserById = `-- name: GetUserById :one
SELECT id, username, email, created_at, updated_at
FROM users
WHERE id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
