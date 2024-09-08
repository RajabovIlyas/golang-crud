package user

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.ResponseUser, error)
	SetUserCtx(ctx context.Context, key string, seconds int, user *models.ResponseUser) error
	DeleteUserCtx(ctx context.Context, key string) error
}
