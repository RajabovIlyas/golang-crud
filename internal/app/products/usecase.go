package products

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type UseCase interface {
	Find(context.Context) ([]models.Products, error)
	FindById(context.Context, string) (models.Products, error)
	Create(context.Context, models.CreateUser) (models.Products, error)
	Update(context.Context, models.UpdateUserReq) (models.Products, error)
	Delete(context.Context, string) error
	FindByUsername(context.Context, string) (models.Products, error)
	UpdatePasswordById(context.Context, models.UpdatePasswordReq) (models.Products, error)
}
