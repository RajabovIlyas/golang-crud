package usecase

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
)

type userUC struct {
	cfg      *config.Config
	userRepo user.Repository
}

func NewUserUseCase(cfg *config.Config, userRepo user.Repository) user.UseCase {
	return &userUC{cfg: cfg, userRepo: userRepo}
}

func (u userUC) Find(ctx context.Context) ([]database.FindUsersRow, error) {
	return u.userRepo.Find(ctx)
}

func (u userUC) FindById(ctx context.Context, userIDStr string) (database.FindUserByIdRow, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return database.FindUserByIdRow{}, err
	}
	return u.userRepo.FindByID(ctx, userID)
}

func (u userUC) Create(ctx context.Context, createUser models.CreateUser) (database.CreateUserRow, error) {
	return u.userRepo.Create(ctx, database.CreateUserParams{Username: createUser.Username, Password: createUser.Password})
}

func (u userUC) Update(ctx context.Context, userIDStr string, updateUser models.UpdateUser) (database.UpdateUserByIdRow, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return database.UpdateUserByIdRow{}, err
	}
	return u.userRepo.UpdateByID(ctx, database.UpdateUserByIdParams{
		ID:       userID,
		Username: updateUser.Username,
	})
}

func (u userUC) Delete(ctx context.Context, userIDStr string) error {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return err
	}
	return u.userRepo.DeleteByID(ctx, userID)
}

func (u userUC) FindByUsername(ctx context.Context, username string) (database.FindUserByUsernameRow, error) {
	return u.userRepo.FindByUsername(ctx, username)
}

func (u userUC) UpdatePasswordById(ctx context.Context, updateUser models.UpdatePassword) (database.UpdateUserPasswordByIdRow, error) {
	newUserPassword, _ := utils.HashPassword(updateUser.Password)
	userID, err := uuid.Parse(updateUser.ID)
	if err != nil {
		return database.UpdateUserPasswordByIdRow{}, err
	}
	return u.userRepo.UpdatePasswordById(ctx, database.UpdateUserPasswordByIdParams{
		Password: newUserPassword,
		ID:       userID,
	})
}
