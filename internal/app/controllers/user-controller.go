package controllers

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/common"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/services"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	us *services.UsersService
}

func NewUserController(db *database.Queries) *UserController {
	return &UserController{us: services.NewUsersService(db)}
}

func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.us.GetUsers(c.Request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"users": common.DatabaseUsersToUsers(users)})
}

func (uc *UserController) GetUser(c *gin.Context) {
	userID := c.Param("userID")

	user, err := uc.us.GetUser(c.Request, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": common.DatabaseUserToUser(user)})
}

func (uc *UserController) CreateUser(c *gin.Context) {
	var newUser models.CreateOrChangeUser
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.us.CreateUser(c.Request, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": common.DatabaseUserToUser(user)})
}

func (uc *UserController) ChangeUser(c *gin.Context) {
	userID := c.Param("userID")

	var changeUser models.CreateOrChangeUser
	if err := c.BindJSON(&changeUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	user, err := uc.us.UpdateUser(c.Request, userID, changeUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": common.DatabaseUserToUser(user)})
}

func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("userID")

	err := uc.us.DeleteUser(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete user successfully!"})
}
