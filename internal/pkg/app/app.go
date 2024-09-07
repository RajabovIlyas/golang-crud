package app

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/server"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/db/postgres"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/db/redis"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/logger"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func Run() error {

	loadConfig, err := config.LoadConfig(constants.CONFIG_FILE_PATH)

	if err != nil {
		log.Fatal().Msg("load config error:" + err.Error())
		return err
	}

	cfg, _ := config.ParseConfig(loadConfig)

	db, cdb, err := postgres.NewPsqlDB(cfg)

	if err != nil {
		log.Fatal().Msg("db connection error:" + err.Error())
		return err
	}

	redisClient := redis.NewRedisClient(cfg)

	logger.Logger()

	g := gin.Default()

	s := server.NewServer(g, cfg, db, redisClient)

	log.Info().Msg("Server started")

	defer postgres.DisconnectPsqlDB(cdb)
	defer redis.DisconnectRedis(redisClient)

	if err = s.Run(); err != nil {
		log.Fatal().Msg(err.Error())
	}

	return err
}
