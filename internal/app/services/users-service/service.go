package usersService

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type Service interface {
	FindUsers(context.Context) ([]database.FindUsersRow, error)
	FindUserById(context.Context, uuid.UUID) (database.FindUserByIdRow, error)
	CreateUser(context.Context, models.CreateUser) (database.CreateUserRow, error)
	UpdateUser(context.Context, string, models.UpdateUser) (database.UpdateUserByIdRow, error)
	DeleteUser(context.Context, string) error
	FindUserByUserName(context.Context, string) (database.FindUserByUsernameRow, error)
	UpdateUserPasswordById(context.Context, models.UpdatePassword) (database.UpdateUserPasswordByIdRow, error)
}
