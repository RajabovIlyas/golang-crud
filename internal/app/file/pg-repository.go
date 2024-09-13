package file

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/google/uuid"
)

type Repository interface {
	Create(models.CreateFile) (models.Files, error)
	FindByID(uuid.UUID) (models.Files, error)
	FindByFileName(string) (models.Files, error)
	Delete(uuid.UUID) error
}
