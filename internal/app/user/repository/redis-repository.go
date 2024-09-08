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

type userRedisRepo struct {
	redisClient *redis.Client
}

func NewUserRedisRepo(redisClient *redis.Client) user.RedisRepository {
	return &userRedisRepo{redisClient: redisClient}
}

func (a *userRedisRepo) GetByIDCtx(ctx context.Context, key string) (*models.ResponseUser, error) {

	userBytes, err := a.redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "userRedisRepo.GetByIDCtx.redisClient.Get")
	}
	responseUser := &models.ResponseUser{}
	if err = json.Unmarshal(userBytes, responseUser); err != nil {
		return nil, errors.Wrap(err, "userRedisRepo.GetByIDCtx.json.Unmarshal")
	}
	return responseUser, nil
}

func (a *userRedisRepo) SetUserCtx(ctx context.Context, key string, seconds int, user *models.ResponseUser) error {

	userBytes, err := json.Marshal(user)
	if err != nil {
		return errors.Wrap(err, "userRedisRepo.SetUserCtx.json.Unmarshal")
	}
	if err = a.redisClient.Set(ctx, key, userBytes, time.Second*time.Duration(seconds)).Err(); err != nil {
		return errors.Wrap(err, "userRedisRepo.SetUserCtx.redisClient.Set")
	}
	return nil
}

func (a *userRedisRepo) DeleteUserCtx(ctx context.Context, key string) error {

	if err := a.redisClient.Del(ctx, key).Err(); err != nil {
		return errors.Wrap(err, "userRedisRepo.DeleteUserCtx.redisClient.Del")
	}
	return nil
}
