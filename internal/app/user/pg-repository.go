package user

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/google/uuid"
)

type Repository interface {
	Find() ([]models.Users, error)
	FindByID(uuid.UUID) (models.Users, error)
	Create(models.CreateUser) (models.Users, error)
	UpdateByID(models.UpdateUser) (models.Users, error)
	DeleteByID(uuid.UUID) error
	FindByUsername(string) (models.Users, error)
	UpdatePasswordById(models.UpdatePassword) (models.Users, error)
}
