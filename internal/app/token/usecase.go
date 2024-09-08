package token

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type UseCase interface {
	GenerateToken(context.Context, uuid.UUID) (models.ResponseToken, error)
	GenerateTokenByToken(models.GenerateTokenModel) (models.ResponseToken, error)
	FindTokenById(context.Context, string) (database.FindTokenByIdRow, error)
	FindTokenByAccessKey(context.Context, string) (database.FindTokenByAccessKeyRow, error)
	CreateToken(context.Context, uuid.UUID) (database.CreateTokenRow, error)
	UpdateToken(context.Context, uuid.UUID) (database.UpdateTokenByIdRow, error)
	DeleteTokenById(context.Context, uuid.UUID) error
	DeleteTokenByAccessKey(context.Context, uuid.UUID) error
	DeleteOldTokens(context.Context) error
}
