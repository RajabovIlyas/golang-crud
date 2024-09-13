package models

import (
	"github.com/google/uuid"
	"time"
)

type ResponseToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type GenerateTokenModel struct {
	ID             uuid.UUID
	AccessTokenKey uuid.UUID
}

type RefreshTokenModel struct {
	RefreshToken string `json:"refresh_token"`
}

type Tokens struct {
	ID             uuid.UUID `json:"id" gorm:"primary_key;type:uuid;default:gen_random_uuid()"`
	AccessTokenKey uuid.UUID `json:"access_token_key" gorm:"type:uuid;default:gen_random_uuid()"`
	UserID         uuid.UUID `json:"user_id" gorm:"not null"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
}
