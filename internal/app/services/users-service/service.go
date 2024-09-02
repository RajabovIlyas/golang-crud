package usersService

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type Service interface {
	FindUsers(*http.Request) ([]database.FindUsersRow, error)
	FindUserById(*http.Request, uuid.UUID) (database.FindUserByIdRow, error)
	CreateUser(*http.Request, models.CreateUser) (database.CreateUserRow, error)
	UpdateUser(*http.Request, string, models.UpdateUser) (database.UpdateUserByIdRow, error)
	DeleteUser(*http.Request, string) error
	FindUserByUserName(*http.Request, string) (database.FindUserByUsernameRow, error)
	UpdateUserPasswordById(*http.Request, models.UpdatePassword) (database.UpdateUserPasswordByIdRow, error)
}
