package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateUser struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateUser struct {
	ID       string `json:"id"`
	Username string `json:"username"`
}

type UpdatePassword struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
