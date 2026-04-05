package auth

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	id            uuid.UUID
	name          string
	email         string
	password_hash string
	created_at    time.Time
	last_login    *time.Time
}
