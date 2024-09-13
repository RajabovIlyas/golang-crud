package token

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.Tokens, error)
	SetTokenCtx(ctx context.Context, key string, user *models.Tokens) error
	DeleteTokenCtx(ctx context.Context, key string) error
}
