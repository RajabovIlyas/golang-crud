package user

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.UserModel, error)
	SetUserCtx(ctx context.Context, key string, user *models.UserModel) error
	DeleteUserCtx(ctx context.Context, key string) error
}
