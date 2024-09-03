package authService

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type Service interface {
	Login(context.Context, models.UserLogin) (models.ResponseToken, error)
	Logout(context.Context, string) error
	Register(context.Context, models.CreateUser) (models.ResponseToken, error)
	Refresh(context.Context, uuid.UUID) (models.ResponseToken, error)
	AuthMe(context.Context, uuid.UUID) (database.FindUserByIdRow, error)
}
