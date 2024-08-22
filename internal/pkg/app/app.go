package app

import (
	"database/sql"
	"github.com/RajabovIlyas/golang-crud/internal/app/common"
	"github.com/RajabovIlyas/golang-crud/internal/app/middleware"
	"github.com/RajabovIlyas/golang-crud/internal/app/routes"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

type App struct {
	g *gin.Engine
	r *routes.Route
	c common.Config
}

func New() (*App, error) {

	a := &App{}

	a.c = common.GetConfig()

	conn, err := sql.Open("postgres", a.c.UrlDB)
	if err != nil {
		log.Fatal().Msg("db connection error:" + err.Error())
	}

	common.Logger()

	a.g = gin.Default()

	a.g.Use(middleware.Logger)

	a.r = routes.New(a.g, database.New(conn))

	a.r.PaveRoutes()

	return a, nil
}

func (a *App) Run() error {
	log.Info().Msg("Server started")

	err := a.g.Run(a.c.Port)
	return err
}
