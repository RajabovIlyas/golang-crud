package services

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type UsersService struct {
	db *database.Queries
}

func NewUsersService(db *database.Queries) *UsersService {
	return &UsersService{db}
}

func (us *UsersService) GetUsers(r *http.Request) ([]database.User, error) {
	return us.db.GetUsers(r.Context())
}

func (us *UsersService) GetUser(r *http.Request, userIDStr string) (database.User, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return database.User{}, err
	}

	return us.db.GetUserById(r.Context(), userID)
}

func (us *UsersService) CreateUser(r *http.Request, newUser models.CreateOrChangeUser) (database.User, error) {
	return us.db.CreateUser(r.Context(), newUser.Name)
}

func (us *UsersService) UpdateUser(r *http.Request, userIDStr string, changeUser models.CreateOrChangeUser) (database.User, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return database.User{}, err
	}
	return us.db.ChangeUserById(r.Context(), database.ChangeUserByIdParams{
		ID:   userID,
		Name: changeUser.Name,
	})
}

func (us *UsersService) DeleteUser(r *http.Request, userIDStr string) error {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return err
	}
	return us.db.DeleteUserById(r.Context(), userID)
}
