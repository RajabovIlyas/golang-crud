package authService

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type Service interface {
	Login(*http.Request, models.UserLogin) (models.ResponseToken, error)
	Logout(*http.Request, string) error
	Register(*http.Request, models.CreateUser) (models.ResponseToken, error)
	Refresh(*http.Request, uuid.UUID) (models.ResponseToken, error)
	AuthMe(*http.Request, uuid.UUID) (database.FindUserByIdRow, error)
}
