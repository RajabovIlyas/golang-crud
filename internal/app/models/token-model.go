package models

import "github.com/google/uuid"

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

type TokenModel struct {
	ID             uuid.UUID `json:"id"`
	AccessTokenKey uuid.UUID `json:"access_token_key"`
	UserID         uuid.UUID `json:"user_id"`
}
