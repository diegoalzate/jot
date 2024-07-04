-- name: CreateUser :one
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, email, created_at, updated_at;

-- name: GetUserById :one
SELECT id, username, email, created_at, updated_at
FROM users
WHERE id = $1;

-- name: GetIdentity :one
SELECT id, user_id, provider
FROM identities
WHERE provider_id = $1 AND provider = $2;

-- name: CreateIdentity :one
INSERT INTO identities (id, user_id, provider, provider_id, identity_data, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING id, user_id, provider, provider_id, identity_data, created_at, updated_at;