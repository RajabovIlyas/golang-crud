package migration

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {

	if check := db.Migrator().HasTable(&models.Users{}); check == false {
		_ = db.Migrator().CreateTable(&models.Users{})
	}

	if check := db.Migrator().HasTable(&models.Tokens{}); check == false {
		_ = db.Migrator().CreateTable(&models.Tokens{})
	}

	if check := db.Migrator().HasTable(&models.Files{}); check == false {
		_ = db.Migrator().CreateTable(&models.Files{})
	}
}
