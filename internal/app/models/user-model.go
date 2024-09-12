package models

import (
	"github.com/google/uuid"
	"time"
)

type CreateUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UpdateUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type UpdateUserReq struct {
	ID       string `json:"id"`
	Username string `json:"username" binding:"required"`
}

type UpdatePassword struct {
	ID       uuid.UUID `json:"id"`
	Password string    `json:"password"`
}

type UpdatePasswordReq struct {
	ID       string `json:"id"`
	Password string `json:"password" binding:"required"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseUser struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type Users struct {
	ID        uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	Username  string    `json:"username" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Password  string    `json:"password" gorm:"not null"`
}
