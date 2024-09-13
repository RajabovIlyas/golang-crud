package migration

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"gorm.io/gorm"
)

func ResetMigrations(db *gorm.DB) {

	if check := db.Migrator().HasTable(&models.Users{}); check != false {
		_ = db.Migrator().DropTable(&models.Users{})
	}

	if check := db.Migrator().HasTable(&models.Tokens{}); check != false {
		_ = db.Migrator().DropTable(&models.Tokens{})
	}

	if check := db.Migrator().HasTable(&models.Files{}); check != false {
		_ = db.Migrator().DropTable(&models.Files{})
	}
}
