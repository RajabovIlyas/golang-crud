package redis

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"time"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	redisHost := cfg.Redis.RedisAddr

	if redisHost == "" {
		redisHost = ":6379"
	}

	client := redis.NewClient(&redis.Options{
		Addr: redisHost,

		PoolTimeout: time.Duration(cfg.Redis.PoolTimeout) * time.Second,
		Password:    cfg.Redis.Password, // no password set
		DB:          cfg.Redis.DB,       // use default DB
	})

	return client
}

func DisconnectRedis(client *redis.Client, logger zerolog.Logger) {
	err := client.Close()
	if err != nil {
		logger.Fatal().Msg("Error to disconnect Redis: " + err.Error())
	}
}
