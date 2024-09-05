package utils

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
)

func DatabaseResponseUserToResponseUser(dbUser models.ResponseUser) models.ResponseUser {
	return models.ResponseUser{
		ID:       dbUser.ID,
		Username: dbUser.Username,
	}
}

func DatabaseResponseUsersToResponseUsers(dbUsers []database.FindUsersRow) []models.ResponseUser {
	users := make([]models.ResponseUser, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = DatabaseResponseUserToResponseUser(models.ResponseUser(u))
	}
	return users
}
