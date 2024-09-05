package app

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/app/cron"
	"github.com/RajabovIlyas/golang-crud/internal/app/middleware"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/routes"
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

	c, _ := config.ParseConfig(loadConfig)

	db, cdb, err := postgres.NewPsqlDB(c)

	if err != nil {
		log.Fatal().Msg("db connection error:" + err.Error())
		return err
	}

	p := &models.DBConfigParam{db, c}

	rc := redis.NewRedisClient(c)

	logger.Logger()

	g := gin.Default()

	g.Use(middleware.Logger)

	cs := cron.NewCronService(p)

	cs.DeleteAllToken()

	r := routes.New(g, rc, p)

	r.PaveRoutes()

	log.Info().Msg("Server started")

	defer postgres.DisconnectPsqlDB(cdb)
	defer redis.DisconnectRedis(rc)

	return g.Run(c.Server.Port)
}
