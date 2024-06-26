-- name: CreateUser :one
INSERT INTO users (id, username, email, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, username, email, created_at, updated_at;

-- name: GetIdentity :one
SELECT id, user_id, provider
FROM identities
WHERE provider_id = $1 AND provider = $2;