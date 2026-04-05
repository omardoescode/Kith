package auth

import (
	"time"

	"github.com/google/uuid"
)

func CreateUser() (*User, error) {
	password, err := HashPassword("TestPassword")
	if err != nil {
		return nil, err
	}

	id, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	return &User{
		id:            id,
		name:          "Omar Mohamed",
		email:         "Omar@gmail.com",
		password_hash: password,
		created_at:    time.Now(),
		last_login:    nil,
	}, nil
}
