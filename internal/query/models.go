// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package query

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type Identity struct {
	ID           uuid.UUID
	UserID       uuid.UUID
	ProviderID   string
	Provider     string
	IdentityData json.RawMessage
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type User struct {
	ID        uuid.UUID
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
