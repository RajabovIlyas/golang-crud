package token

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.TokenModel, error)
	SetTokenCtx(ctx context.Context, key string, user *models.TokenModel) error
	DeleteTokenCtx(ctx context.Context, key string) error
}
