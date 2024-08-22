package common

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"os"
)

type Config struct {
	UrlDB string
	Port  string
}

func GetConfig() Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn().Msg(err.Error())
	}

	urlDB := os.Getenv("DB_URL")
	if urlDB == "" {
		log.Fatal().Msg("$DB_URL must be set")
	}

	port := os.Getenv("PORT")

	if port == "" {
		port = constants.Port
	}
	return Config{urlDB, ":" + port}
}
