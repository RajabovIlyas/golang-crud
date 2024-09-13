package migration

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/constants"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/db/postgres"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

func MigrateConfig() (*gorm.DB, error) {
	loadConfig, err := config.LoadConfig(constants.CONFIG_FILE_PATH)

	if err != nil {

		return nil, err
	}

	cfg, err := config.ParseConfig(loadConfig, log.Logger)

	if err != nil {
		return nil, err
	}

	return postgres.NewPsqlDB(cfg)
}
