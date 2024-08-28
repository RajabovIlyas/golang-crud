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

// GetUsers List : Return list of users
//
//	@Summary		List all users
//	@Description	List all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.User
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/ [get]
func (uc *UserController) GetUsers(c *gin.Context) {
	users, err := uc.us.GetUsers(c.Request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, gin.H{"users": common.DatabaseUsersToUsers(users)})
}

// GetUser Get : Return user by id
//
//	@Summary		Return user by id
//	@Description	Return user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.User
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/{id} [get]
func (uc *UserController) GetUser(c *gin.Context) {
	userID := c.Param("userID")

	user, err := uc.us.GetUser(c.Request, userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": common.DatabaseUserToUser(user)})
}

// CreateUser Create : Creates new  user.
//
//	@Summary		Create new user record in Library
//	@Description	Create new user record in Library
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.CreateOrChangeUser	true	"Add user"
//	@Success		200		{object}	models.User
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/users/ [post]
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

	c.JSON(http.StatusOK, gin.H{"data": common.DatabaseUserToUser(user)})
}

// ChangeUser Update : Update user by id
//
//	@Summary		Update user details
//	@Description	Update user details
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string						true	"id"
//	@Param			user	body		models.CreateOrChangeUser	true	"Update user"
//	@Success		200		{object}	models.User
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/users/{id} [put]
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

	c.JSON(http.StatusOK, gin.H{"data": common.DatabaseUserToUser(user)})
}

// DeleteUser Delete : Delete user by id
//
//	@Summary		Delete user record by id
//	@Description	Delete user record by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.User
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	userID := c.Param("userID")

	err := uc.us.DeleteUser(c.Request, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Delete user successfully!"})
}
