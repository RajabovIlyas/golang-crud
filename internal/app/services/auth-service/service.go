package authService

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type Service interface {
	Login(*http.Request, models.UserLogin) (models.TokenModel, error)
	Logout(*http.Request, string) error
	Register(*http.Request, models.CreateUser) (models.TokenModel, error)
	Refresh(*http.Request, uuid.UUID) (models.TokenModel, error)
	AuthMe(*http.Request, uuid.UUID) (database.FindUserByIdRow, error)
}
