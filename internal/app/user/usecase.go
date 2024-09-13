package user

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type UseCase interface {
	Find(context.Context) ([]models.Users, error)
	FindById(context.Context, string) (models.Users, error)
	Create(context.Context, models.CreateUser) (models.Users, error)
	Update(context.Context, models.UpdateUserReq) (models.Users, error)
	Delete(context.Context, string) error
	FindByUsername(context.Context, string) (models.Users, error)
	UpdatePasswordById(context.Context, models.UpdatePasswordReq) (models.Users, error)
}
