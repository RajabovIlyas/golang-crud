package http

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/users"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/httpErrors"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/httpResponse"
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandlers struct {
	cfg    *config.Config
	userUC users.UseCase
}

func NewUserHandlers(cfg *config.Config, userUC users.UseCase) users.Handlers {
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
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(httpResponse.SuccessResponse(users))
}

// GetUser Get : Return users by id
//
//	@Summary		Return users by id
//	@Description	Return users by id
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
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(httpResponse.SuccessResponse(foundUser))
}

// UpdateUser Update : Update users by id
//
//	@Summary		Update users details
//	@Description	Update users details
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			id		path		string				true	"id"
//	@Param			users	body		models.UpdateUser	true	"Update users"
//	@Success		200		{object}	models.Users
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/users/{id} [put]
func (u userHandlers) UpdateUser(g *gin.Context) {
	userID := g.Param("userID")

	var updateUser models.UpdateUserReq
	if err := g.BindJSON(&updateUser); err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}
	updateUser.ID = userID

	updatedUser, err := u.userUC.Update(context.Background(), updateUser)
	if err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(httpResponse.SuccessResponse(updatedUser))
}

// DeleteUser Delete : Delete users by id
//
//	@Summary		Delete users record by id
//	@Description	Delete users record by id
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
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(httpResponse.NoContentResponse("Deleted users " + userID))
}

// UpdateUserPassword Update : Update users password by id
//
//	@Summary		Delete users record by id
//	@Description	Delete users record by id
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
		g.JSON(http.StatusBadRequest, httpErrors.BadRequest)
		return
	}

	updatePassword.ID = g.Param("userID")

	updatedUser, err := u.userUC.UpdatePasswordById(context.Background(), updatePassword)
	if err != nil {
		g.JSON(httpErrors.ErrorResponse(err))
		return
	}

	g.JSON(httpResponse.SuccessResponse(updatedUser))
}
