package token

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type Repository interface {
	FindByID(context.Context, uuid.UUID) (database.FindTokenByIdRow, error)
	FindByAccessKey(context.Context, uuid.UUID) (database.FindTokenByAccessKeyRow, error)
	Create(context.Context, uuid.UUID) (database.CreateTokenRow, error)
	UpdateByID(context.Context, uuid.UUID) (database.UpdateTokenByIdRow, error)
	DeleteByID(context.Context, uuid.UUID) (database.DeleteTokenByIdRow, error)
	DeleteByAccessKey(context.Context, uuid.UUID) (database.DeleteTokenByAccessKeyRow, error)
	DeleteOldTokens(context.Context) ([]database.DeleteOldTokensRow, error)
}
