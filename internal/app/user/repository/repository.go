package repository

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type userRepo struct {
	db *database.Queries
}

func NewUserRepository(db *database.Queries) user.Repository {
	return &userRepo{db}
}

func (u userRepo) Find(ctx context.Context) ([]database.FindUsersRow, error) {
	return u.db.FindUsers(ctx)
}

func (u userRepo) FindByID(ctx context.Context, uuid uuid.UUID) (database.FindUserByIdRow, error) {
	return u.db.FindUserById(ctx, uuid)
}

func (u userRepo) Create(ctx context.Context, createUser database.CreateUserParams) (database.CreateUserRow, error) {
	return u.db.CreateUser(ctx, createUser)
}

func (u userRepo) UpdateByID(ctx context.Context, updateUser database.UpdateUserByIdParams) (database.UpdateUserByIdRow, error) {
	return u.db.UpdateUserById(ctx, updateUser)
}

func (u userRepo) DeleteByID(ctx context.Context, id uuid.UUID) error {
	return u.db.DeleteUserById(ctx, id)
}

func (u userRepo) FindByUsername(ctx context.Context, username string) (database.FindUserByUsernameRow, error) {
	return u.db.FindUserByUsername(ctx, username)
}

func (u userRepo) UpdatePasswordById(ctx context.Context, updatePassword database.UpdateUserPasswordByIdParams) (database.UpdateUserPasswordByIdRow, error) {
	return u.db.UpdateUserPasswordById(ctx, updatePassword)
}
