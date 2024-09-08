package user

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type Repository interface {
	Find(context.Context) ([]database.FindUsersRow, error)
	FindByID(context.Context, uuid.UUID) (database.FindUserByIdRow, error)
	Create(context.Context, database.CreateUserParams) (database.CreateUserRow, error)
	UpdateByID(context.Context, database.UpdateUserByIdParams) (database.UpdateUserByIdRow, error)
	DeleteByID(context.Context, uuid.UUID) error
	FindByUsername(context.Context, string) (database.FindUserByUsernameRow, error)
	UpdatePasswordById(context.Context, database.UpdateUserPasswordByIdParams) (database.UpdateUserPasswordByIdRow, error)
}
