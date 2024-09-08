package user

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
)

type UseCase interface {
	Find(context.Context) ([]database.FindUsersRow, error)
	FindById(context.Context, string) (models.ResponseUser, error)
	Create(context.Context, models.CreateUser) (models.ResponseUser, error)
	Update(context.Context, string, models.UpdateUser) (models.ResponseUser, error)
	Delete(context.Context, string) error
	FindByUsername(context.Context, string) (database.FindUserByUsernameRow, error)
	UpdatePasswordById(context.Context, models.UpdatePassword) (database.UpdateUserPasswordByIdRow, error)
}
