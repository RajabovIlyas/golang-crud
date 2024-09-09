package repository

import (
	"context"
	"encoding/json"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	basePrefix    = "api-user:"
	cacheDuration = 3600
)

type userRedisRepo struct {
	redisClient *redis.Client
}

func NewUserRedisRepo(redisClient *redis.Client) user.RedisRepository {
	return &userRedisRepo{redisClient: redisClient}
}

func (u *userRedisRepo) GetByIDCtx(ctx context.Context, key string) (*models.UserModel, error) {

	userBytes, err := u.redisClient.Get(ctx, u.generateUserKey(key)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "userRedisRepo.GetByIDCtx.redisClient.Get")
	}
	userModel := &models.UserModel{}
	if err = json.Unmarshal(userBytes, userModel); err != nil {
		return nil, errors.Wrap(err, "userRedisRepo.GetByIDCtx.json.Unmarshal")
	}
	return userModel, nil
}

func (u *userRedisRepo) SetUserCtx(ctx context.Context, key string, user *models.UserModel) error {

	userBytes, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "userRedisRepo.SetUserCtx.json.Unmarshal")
	}
	if err = u.redisClient.Set(ctx, u.generateUserKey(key), userBytes, time.Second*time.Duration(cacheDuration)).Err(); err != nil {
		return errors.Wrap(err, "userRedisRepo.SetUserCtx.redisClient.Set")
	}
	return nil
}

func (u *userRedisRepo) DeleteUserCtx(ctx context.Context, key string) error {

	if err := u.redisClient.Del(ctx, u.generateUserKey(key)).Err(); err != nil {
		return errors.Wrap(err, "userRedisRepo.DeleteUserCtx.redisClient.Del")
	}
	return nil
}

func (u *userRedisRepo) generateUserKey(key string) string {
	return basePrefix + key
}
