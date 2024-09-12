package repository

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepo{db}
}

func (u userRepo) Find() ([]models.Users, error) {
	var users []models.Users
	result := u.db.Find(&users)
	return users, result.Error
}

func (u userRepo) FindByID(id uuid.UUID) (models.Users, error) {
	var findUser models.Users
	result := u.db.First(&findUser, id)
	return findUser, result.Error
}

func (u userRepo) Create(createUser models.CreateUser) (models.Users, error) {
	var saveUser models.Users
	saveUser.Username = createUser.Username
	saveUser.Password = createUser.Password
	result := u.db.Create(&saveUser)
	return saveUser, result.Error
}

func (u userRepo) UpdateByID(updateUser models.UpdateUser) (models.Users, error) {
	var updatedUser models.Users
	result := u.db.First(&updateUser, updateUser.ID)
	if result.Error != nil {
		return updatedUser, result.Error
	}
	updatedUser.Username = updateUser.Username
	result = u.db.Save(&updatedUser)
	return updatedUser, result.Error
}

func (u userRepo) DeleteByID(id uuid.UUID) error {
	result := u.db.Delete(&models.Users{}, id)
	return result.Error
}

func (u userRepo) FindByUsername(username string) (models.Users, error) {
	var findUser models.Users
	result := u.db.Where("username = ?", username).First(&findUser)
	return findUser, result.Error
}

func (u userRepo) UpdatePasswordById(updatePassword models.UpdatePassword) (models.Users, error) {
	var updateUser models.Users
	result := u.db.First(&updateUser, updatePassword.ID)
	if result.Error != nil {
		return updateUser, result.Error
	}
	updateUser.Password = updatePassword.Password
	result = u.db.Save(&updateUser)
	return updateUser, result.Error
}
