package app

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/server"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/db/postgres"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/db/redis"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func Run(logger zerolog.Logger) error {

	loadConfig, err := config.LoadConfig(constants.CONFIG_FILE_PATH)

	if err != nil {
		logger.Fatal().Msg("load config error:" + err.Error())
		return err
	}

	cfg, _ := config.ParseConfig(loadConfig, logger)

	db, cdb, err := postgres.NewPsqlDB(cfg)

	if err != nil {
		logger.Fatal().Msg("db connection error:" + err.Error())
		return err
	}

	redisClient := redis.NewRedisClient(cfg)

	g := gin.Default()

	s := server.NewServer(g, cfg, db, redisClient, logger)

	logger.Info().Msg("Server started")

	defer postgres.DisconnectPsqlDB(cdb, logger)
	defer redis.DisconnectRedis(redisClient, logger)

	if err = s.Run(); err != nil {
		logger.Fatal().Msg(err.Error())
	}

	return err
}
