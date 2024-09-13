package migration

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&models.Users{}, &models.Tokens{}, &models.Files{})
}
