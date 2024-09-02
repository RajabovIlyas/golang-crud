package usersService

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/common"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/utils"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
	"net/http"
)

type UsersService struct {
	db *database.Queries
	c  *common.Config
}

func NewUsersService(db *database.Queries) *UsersService {
	c, _ := common.GetConfig(".")
	return &UsersService{db, &c}
}

func (us *UsersService) FindUsers(r *http.Request) ([]database.FindUsersRow, error) {
	return us.db.FindUsers(r.Context())
}

func (us *UsersService) FindUserById(r *http.Request, userID uuid.UUID) (database.FindUserByIdRow, error) {
	return us.db.FindUserById(r.Context(), userID)
}

func (us *UsersService) CreateUser(r *http.Request, newUser models.CreateUser) (database.CreateUserRow, error) {
	return us.db.CreateUser(r.Context(), database.CreateUserParams{Username: newUser.Username, Password: newUser.Password})
}

func (us *UsersService) UpdateUser(r *http.Request, userIDStr string, changeUser models.UpdateUser) (database.UpdateUserByIdRow, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return database.UpdateUserByIdRow{}, err
	}
	return us.db.UpdateUserById(r.Context(), database.UpdateUserByIdParams{
		ID:       userID,
		Username: changeUser.Username,
	})
}

func (us *UsersService) DeleteUser(r *http.Request, userIDStr string) error {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return err
	}
	return us.db.DeleteUserById(r.Context(), userID)
}

func (us *UsersService) FindUserByUserName(r *http.Request, username string) (database.FindUserByUsernameRow, error) {
	return us.db.FindUserByUsername(r.Context(), username)
}

func (us *UsersService) UpdateUserPasswordById(r *http.Request, updateUser models.UpdatePassword) (database.UpdateUserPasswordByIdRow, error) {
	newUserPassword, _ := utils.HashPassword(updateUser.Password)
	userID, err := uuid.Parse(updateUser.ID)
	if err != nil {
		return database.UpdateUserPasswordByIdRow{}, err
	}
	return us.db.UpdateUserPasswordById(r.Context(), database.UpdateUserPasswordByIdParams{
		Password: newUserPassword,
		ID:       userID,
	})
}

func (us *UsersService) Logout(r *http.Request) {}

func (us *UsersService) Register(r *http.Request, user models.CreateUser) (database.CreateUserRow, error) {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	return us.CreateUser(r, user)
}
