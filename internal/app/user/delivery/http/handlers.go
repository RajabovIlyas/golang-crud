package http

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandlers struct {
	cfg    *config.Config
	userUC user.UseCase
}

func NewUserHandlers(cfg *config.Config, userUC user.UseCase) user.Handlers {
	return &userHandlers{cfg: cfg, userUC: userUC}
}

// GetUsers List : Return list of users
//
//	@Summary		List all users
//	@Description	List all users
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Users
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/ [get]
func (u userHandlers) GetUsers(g *gin.Context) {
	users, err := u.userUC.Find(context.Background())

	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, users)
}

// GetUser Get : Return user by id
//
//	@Summary		Return user by id
//	@Description	Return user by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.Users
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/{id} [get]
func (u userHandlers) GetUser(g *gin.Context) {
	userIDStr := g.Param("userID")

	foundUser, err := u.userUC.FindById(context.Background(), userIDStr)

	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, foundUser)
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
//	@Success		200		{object}	models.Users
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/users/{id} [put]
func (u userHandlers) UpdateUser(g *gin.Context) {
	userID := g.Param("userID")

	var updateUser models.UpdateUserReq
	if err := g.BindJSON(&updateUser); err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: "Error parsing request body"})
		return
	}
	updateUser.ID = userID

	updatedUser, err := u.userUC.Update(context.Background(), updateUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, updatedUser)
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
func (u userHandlers) DeleteUser(g *gin.Context) {
	userID := g.Param("userID")

	err := u.userUC.Delete(context.Background(), userID)
	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, models.Message{Message: "Delete user successfully!"})
}

// UpdateUserPassword Update : Update user password by id
//
//	@Summary		Delete user record by id
//	@Description	Delete user record by id
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id	path		string	true	"id"
//	@Success		200	{object}	models.Users
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/users/password/{id} [delete]
func (u userHandlers) UpdateUserPassword(g *gin.Context) {
	var updatePassword models.UpdatePasswordReq
	if err := g.BindJSON(&updatePassword); err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: "error parsing request body"})
		return
	}

	updatePassword.ID = g.Param("userID")

	updatedUser, err := u.userUC.UpdatePasswordById(context.Background(), updatePassword)
	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, updatedUser)
}
