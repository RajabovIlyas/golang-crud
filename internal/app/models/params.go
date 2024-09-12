package models

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"gorm.io/gorm"
)

type DBConfigParam struct {
	DB *gorm.DB
	C  *config.Config
}
