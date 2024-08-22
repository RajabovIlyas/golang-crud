package common

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/database"
)

func DatabaseUserToUser(dbUser database.User) models.User {
	return models.User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

func DatabaseUsersToUsers(dbUsers []database.User) []models.User {
	users := make([]models.User, len(dbUsers))
	for i, u := range dbUsers {
		users[i] = DatabaseUserToUser(u)
	}
	return users
}
