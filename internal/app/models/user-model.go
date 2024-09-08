package models

import (
	"github.com/google/uuid"
)

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
