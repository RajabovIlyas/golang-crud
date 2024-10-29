package products

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/google/uuid"
)

type Repository interface {
	Find() ([]models.Products, error)
	FindByID(uuid.UUID) (models.Products, error)
	Create(models.CreateUser) (models.Products, error)
	UpdateByID(models.UpdateUser) (models.Products, error)
	DeleteByID(uuid.UUID) error
	UpdateByName(string) (models.Products, error)
}
