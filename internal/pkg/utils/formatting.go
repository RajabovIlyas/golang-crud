package utils

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
)

func DatabaseUserModelToUserModel(dbUser models.UserModel) models.UserModel {
	return models.UserModel{
		ID:       dbUser.ID,
		Username: dbUser.Username,
	}
}

func DatabaseUserModelsToUserModels(dbUsers []database.FindUsersRow) []models.UserModel {
	users := make([]models.UserModel, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = DatabaseUserModelToUserModel(models.UserModel(u))
	}
	return users
}
