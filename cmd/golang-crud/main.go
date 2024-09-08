package main

import (
	_ "github.com/RajabovIlyas/golang-crud/docs"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/app"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/logger"
	"github.com/rs/zerolog/log"
)

//	@title			Go Example REST API
//	@version		1.0v
//	@description	Example Golang REST API
//	@contact.name	Raj Ilyas
//	@contact.url	https://github.com/RajabovIlyas
//	@contact.email	rajabowilyas@gmail.com

//	@securityDefinitions.apikey	ApiKeyAuth
//	@in							header
//	@name 						Authorization

//	@host		localhost:3000
//	@basePath	/api/v1

func main() {
	appLogger, err := logger.InitLogger()

	if err != nil {
		log.Fatal().Err(err).Msg("Error initializing logger")
		return
	}

	err = app.Run(appLogger)

	if err != nil {
		appLogger.Fatal().Msg(err.Error())
	}
}
