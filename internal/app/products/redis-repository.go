package products

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type RedisRepository interface {
	GetByIDCtx(ctx context.Context, key string) (*models.Products, error)
	SetUserCtx(ctx context.Context, key string, user *models.Products) error
	DeleteProductsCtx(ctx context.Context, key string) error
}
