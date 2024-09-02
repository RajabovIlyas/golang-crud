package models

import "github.com/google/uuid"

type TokenModel struct {
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
