package tokensService

import (
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type Service interface {
	FindTokenById(*http.Request, string) (database.FindTokenByIdRow, error)
	FindTokenByAccessKey(*http.Request, string) (database.FindTokenByAccessKeyRow, error)
	CreateToken(*http.Request, uuid.UUID) (database.CreateTokenRow, error)
	UpdateToken(*http.Request, uuid.UUID) (database.UpdateTokenByIdRow, error)
	DeleteTokenById(*http.Request, uuid.UUID) error
	DeleteTokenByAccessKey(r *http.Request, accessKey uuid.UUID) error
	DeleteOldTokens(r *http.Request) error
}
