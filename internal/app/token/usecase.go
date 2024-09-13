package token

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/google/uuid"
)

type UseCase interface {
	GenerateToken(context.Context, uuid.UUID) (models.ResponseToken, error)
	GenerateTokenByToken(models.GenerateTokenModel) (models.ResponseToken, error)
	FindTokenById(context.Context, string) (models.Tokens, error)
	FindTokenByAccessKey(context.Context, string) (models.Tokens, error)
	CreateToken(context.Context, uuid.UUID) (models.Tokens, error)
	UpdateToken(context.Context, uuid.UUID) (models.Tokens, error)
	DeleteTokenById(context.Context, uuid.UUID) error
	DeleteTokenByAccessKey(context.Context, uuid.UUID) error
	DeleteOldTokens(context.Context) error
}
