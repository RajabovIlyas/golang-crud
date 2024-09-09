package auth

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
)

type UseCase interface {
	Login(context.Context, models.UserLogin) (models.ResponseToken, error)
	Logout(context.Context, string) error
	Register(context.Context, models.CreateUser) (models.ResponseToken, error)
	Refresh(context.Context, string) (models.ResponseToken, error)
	AuthMe(context.Context, string) (models.UserModel, error)
}
