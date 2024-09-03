package tokensService

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type Service interface {
	FindTokenById(context.Context, string) (database.FindTokenByIdRow, error)
	FindTokenByAccessKey(context.Context, string) (database.FindTokenByAccessKeyRow, error)
	CreateToken(context.Context, uuid.UUID) (database.CreateTokenRow, error)
	UpdateToken(context.Context, uuid.UUID) (database.UpdateTokenByIdRow, error)
	DeleteTokenById(context.Context, uuid.UUID) error
	DeleteTokenByAccessKey(context.Context, uuid.UUID) error
	DeleteOldTokens(context.Context) error
}
