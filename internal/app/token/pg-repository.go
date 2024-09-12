package token

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/google/uuid"
)

type Repository interface {
	FindByID(uuid.UUID) (models.Tokens, error)
	FindByAccessKey(uuid.UUID) (models.Tokens, error)
	Create(uuid.UUID) (models.Tokens, error)
	UpdateByID(uuid.UUID) (models.Tokens, error)
	DeleteByID(uuid.UUID) (models.Tokens, error)
	DeleteByAccessKey(uuid.UUID) (models.Tokens, error)
	DeleteOldTokens() ([]models.Tokens, error)
}
