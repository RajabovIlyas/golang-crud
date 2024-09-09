package repository

import (
	"context"
	"encoding/json"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	basePrefix    = "api-token:"
	cacheDuration = 3600
)

type tokenRedisRepo struct {
	redisClient *redis.Client
}

func NewTokenRedisRepo(redisClient *redis.Client) token.RedisRepository {
	return &tokenRedisRepo{redisClient: redisClient}
}

func (t *tokenRedisRepo) GetByIDCtx(ctx context.Context, key string) (*models.TokenModel, error) {

	tokenBytes, err := t.redisClient.Get(ctx, t.generateTokenKey(key)).Bytes()
	if err != nil {
		return nil, errors.Wrap(err, "tokenRedisRepo.GetByIDCtx.redisClient.Get")
	}
	tokenModel := &models.TokenModel{}
	if err = json.Unmarshal(tokenBytes, tokenModel); err != nil {
		return nil, errors.Wrap(err, "tokenRedisRepo.GetByIDCtx.json.Unmarshal")
	}
	return tokenModel, nil
}

func (t *tokenRedisRepo) SetTokenCtx(ctx context.Context, key string, token *models.TokenModel) error {

	tokenBytes, err := json.Marshal(token)
	if err != nil {
		return errors.Wrap(err, "tokenRedisRepo.SetTokenCtx.json.Unmarshal")
	}
	if err = t.redisClient.Set(ctx, t.generateTokenKey(key), tokenBytes, time.Second*time.Duration(cacheDuration)).Err(); err != nil {
		return errors.Wrap(err, "tokenRedisRepo.SetTokenCtx.redisClient.Set")
	}
	return nil
}

func (t *tokenRedisRepo) DeleteTokenCtx(ctx context.Context, key string) error {

	if err := t.redisClient.Del(ctx, t.generateTokenKey(key)).Err(); err != nil {
		return errors.Wrap(err, "tokenRedisRepo.DeleteTokenCtx.redisClient.Del")
	}
	return nil
}

func (t *tokenRedisRepo) generateTokenKey(key string) string {
	return basePrefix + key
}
