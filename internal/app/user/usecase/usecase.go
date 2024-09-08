package usecase

import (
	"context"
	"fmt"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type userUC struct {
	cfg       *config.Config
	userRepo  user.Repository
	redisRepo user.RedisRepository
	logger    zerolog.Logger
}

const (
	basePrefix    = "api-user:"
	cacheDuration = 3600
)

func NewUserUseCase(cfg *config.Config, userRepo user.Repository, redisRepo user.RedisRepository, logger zerolog.Logger) user.UseCase {
	return &userUC{cfg: cfg, userRepo: userRepo, redisRepo: redisRepo, logger: logger}
}

func (u userUC) Find(ctx context.Context) ([]database.FindUsersRow, error) {
	foundUsers, err := u.userRepo.Find(ctx)
	if err != nil {
		u.logger.Error().Err(err).Msg("userUC.Find: find users error")
		return nil, err
	}
	return foundUsers, nil
}

func (u userUC) FindById(ctx context.Context, userIDStr string) (models.ResponseUser, error) {

	cachedUser, err := u.redisRepo.GetByIDCtx(ctx, u.GenerateUserKey(userIDStr))
	if err != nil {
		u.logger.Error().Err(err).Msgf("authUC.FindById.GetByIDCtx: %v", err)
	}
	if cachedUser != nil {
		return *cachedUser, nil
	}

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.FindById: invalid userID = %s", userIDStr)
		return models.ResponseUser{}, err
	}
	foundUser, err := u.userRepo.FindByID(ctx, userID)

	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.FindById: find by id = %v", userID)
		return models.ResponseUser{}, err
	}

	responseUser := models.ResponseUser(foundUser)

	if err = u.redisRepo.SetUserCtx(ctx, u.GenerateUserKey(userID.String()), cacheDuration, &responseUser); err != nil {
		u.logger.Error().Err(err).Msgf("authUC.FindById.SetUserCtx: %v", err)
	}

	return responseUser, nil
}

func (u userUC) Create(ctx context.Context, createUser models.CreateUser) (models.ResponseUser, error) {
	createdUser, err := u.userRepo.Create(ctx, database.CreateUserParams{Username: createUser.Username, Password: createUser.Password})
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.Create: error create user")
		return models.ResponseUser{}, err
	}
	return models.ResponseUser(createdUser), nil
}

func (u userUC) Update(ctx context.Context, userIDStr string, updateUser models.UpdateUser) (models.ResponseUser, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.Update: invalid userID = %s", userIDStr)
		return models.ResponseUser{}, err
	}
	updatedUser, err := u.userRepo.UpdateByID(ctx, database.UpdateUserByIdParams{
		ID:       userID,
		Username: updateUser.Username,
	})
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.Update: error update user by id = %v", userID)
		return models.ResponseUser{}, err
	}

	if err = u.redisRepo.DeleteUserCtx(ctx, u.GenerateUserKey(updatedUser.ID.String())); err != nil {
		u.logger.Error().Err(err).Msgf("userUC.Update.DeleteUserCtx: %s", err)
	}

	return models.ResponseUser(updatedUser), nil
}

func (u userUC) Delete(ctx context.Context, userIDStr string) error {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.Delete: invalid userID = %s", userIDStr)
		return err
	}
	err = u.userRepo.DeleteByID(ctx, userID)
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.Delete: error delete user by id = %v", userID)
		return err
	}

	if err = u.redisRepo.DeleteUserCtx(ctx, u.GenerateUserKey(userIDStr)); err != nil {
		u.logger.Error().Err(err).Msgf("userUC.Delete.DeleteUserCtx: %s", err)
	}

	return nil
}

func (u userUC) FindByUsername(ctx context.Context, username string) (database.FindUserByUsernameRow, error) {
	foundUser, err := u.userRepo.FindByUsername(ctx, username)
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.FindByUsername: error find by username = %s", username)
		return database.FindUserByUsernameRow{}, err
	}

	return foundUser, nil
}

func (u userUC) UpdatePasswordById(ctx context.Context, updateUser models.UpdatePassword) (database.UpdateUserPasswordByIdRow, error) {
	newUserPassword, _ := utils.HashPassword(updateUser.Password)
	userID, err := uuid.Parse(updateUser.ID)
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.UpdatePasswordById: invalid userId = %s", updateUser.ID)
		return database.UpdateUserPasswordByIdRow{}, err
	}
	updatedUser, err := u.userRepo.UpdatePasswordById(ctx, database.UpdateUserPasswordByIdParams{
		Password: newUserPassword,
		ID:       userID,
	})
	if err != nil {
		u.logger.Error().Err(err).Msgf("userUC.UpdatePasswordById: error update password by userID = %s", userID)
		return database.UpdateUserPasswordByIdRow{}, err
	}

	return updatedUser, nil
}

func (u *userUC) GenerateUserKey(userID string) string {
	return fmt.Sprintf("%s: %s", basePrefix, userID)
}
