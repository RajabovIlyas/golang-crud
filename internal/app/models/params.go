package models

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/database"
)

type DBConfigParam struct {
	DB *database.Queries
	C  *config.Config
}
