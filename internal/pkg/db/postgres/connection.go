package postgres

import (
	"fmt"
	"github.com/RajabovIlyas/golang-crud/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPsqlDB(cfg *config.Config) (*gorm.DB, error) {

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Postgres.PostgresqlHost,
		cfg.Postgres.PostgresqlPort,
		cfg.Postgres.PostgresqlUser,
		cfg.Postgres.PostgresqlPassword,
		cfg.Postgres.PostgresqlDBName)
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})

	return db, err
}
