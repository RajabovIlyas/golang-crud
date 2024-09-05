package usersController

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/users-service"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type UsersController struct {
	us *usersService.UsersService
}

func NewUsersController(p *models.DBConfigParam) *UsersController {
	return &UsersController{us: usersService.NewUsersService(p)}
}

// GetUsers List : Return list of users
//
//	@Summary		List all users
//	@Description	List all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.ResponseUser
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/ [get]
func (uc *UsersController) GetUsers(c *gin.Context) {
	users, err := uc.us.FindUsers(c.Request.Context())

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.DatabaseResponseUsersToResponseUsers(users))
}

// GetUser Get : Return user by id
//
//	@Summary		Return user by id
//	@Description	Return user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.ResponseUser
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/{id} [get]
func (uc *UsersController) GetUser(c *gin.Context) {
	userIDStr := c.Param("userID")

	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	user, err := uc.us.FindUserById(c.Request.Context(), userID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.DatabaseResponseUserToResponseUser(models.ResponseUser(user)))
}

// CreateUser Create : Creates new  user.
//
//	@Summary		Create new user record in Library
//	@Description	Create new user record in Library
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.CreateUser	true	"Add user"
//	@Success		200		{object}	models.ResponseUser
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/users/ [post]
func (uc *UsersController) CreateUser(c *gin.Context) {
	var newUser models.CreateUser
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	user, err := uc.us.CreateUser(c.Request.Context(), newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.DatabaseResponseUserToResponseUser(models.ResponseUser(user)))
}

// UpdateUser Update : Update user by id
//
//	@Summary		Update user details
//	@Description	Update user details
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"id"
//	@Param			user	body		models.UpdateUser	true	"Update user"
//	@Success		200		{object}	models.ResponseUser
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/users/{id} [put]
func (uc *UsersController) UpdateUser(c *gin.Context) {
	userID := c.Param("userID")

	var updateUser models.UpdateUser
	if err := c.BindJSON(&updateUser); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	user, err := uc.us.UpdateUser(c.Request.Context(), userID, updateUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.DatabaseResponseUserToResponseUser(models.ResponseUser(user)))
}

// DeleteUser Delete : Delete user by id
//
//	@Summary		Delete user record by id
//	@Description	Delete user record by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.Message
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/{id} [delete]
func (uc *UsersController) DeleteUser(c *gin.Context) {
	userID := c.Param("userID")

	err := uc.us.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Message{Message: "Delete user successfully!"})
}

// UpdateUserPassword Update : Update user password by id
//
//	@Summary		Delete user record by id
//	@Description	Delete user record by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.ResponseUser
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/password/{id} [delete]
func (uc *UsersController) UpdateUserPassword(c *gin.Context) {
	var updatePassword models.UpdatePassword
	if err := c.BindJSON(&updatePassword); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	updatePassword.ID = c.Param("userID")

	user, err := uc.us.UpdateUserPasswordById(c.Request.Context(), updatePassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, utils.DatabaseResponseUserToResponseUser(models.ResponseUser(user)))
}
