package usersService

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
)

type UsersService struct {
	db *database.Queries
	c  *config.Config
}

func NewUsersService(p *models.DBConfigParam) *UsersService {

	return &UsersService{db: p.DB, c: p.C}
}

func (us *UsersService) FindUsers(c context.Context) ([]database.FindUsersRow, error) {
	return us.db.FindUsers(c)
}

func (us *UsersService) FindUserById(c context.Context, userID uuid.UUID) (database.FindUserByIdRow, error) {
	return us.db.FindUserById(c, userID)
}

func (us *UsersService) CreateUser(c context.Context, newUser models.CreateUser) (database.CreateUserRow, error) {
	return us.db.CreateUser(c, database.CreateUserParams{Username: newUser.Username, Password: newUser.Password})
}

func (us *UsersService) UpdateUser(c context.Context, userIDStr string, changeUser models.UpdateUser) (database.UpdateUserByIdRow, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return database.UpdateUserByIdRow{}, err
	}
	return us.db.UpdateUserById(c, database.UpdateUserByIdParams{
		ID:       userID,
		Username: changeUser.Username,
	})
}

func (us *UsersService) DeleteUser(c context.Context, userIDStr string) error {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return err
	}
	return us.db.DeleteUserById(c, userID)
}

func (us *UsersService) FindUserByUserName(c context.Context, username string) (database.FindUserByUsernameRow, error) {
	return us.db.FindUserByUsername(c, username)
}

func (us *UsersService) UpdateUserPasswordById(c context.Context, updateUser models.UpdatePassword) (database.UpdateUserPasswordByIdRow, error) {
	newUserPassword, _ := utils.HashPassword(updateUser.Password)
	userID, err := uuid.Parse(updateUser.ID)
	if err != nil {
		return database.UpdateUserPasswordByIdRow{}, err
	}
	return us.db.UpdateUserPasswordById(c, database.UpdateUserPasswordByIdParams{
		Password: newUserPassword,
		ID:       userID,
	})
}

func (us *UsersService) Logout(c context.Context) {}

func (us *UsersService) Register(c context.Context, user models.CreateUser) (database.CreateUserRow, error) {
	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	return us.CreateUser(c, user)
}
