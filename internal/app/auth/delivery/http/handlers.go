package http

import (
	"context"
	"fmt"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/auth"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authHandlers struct {
	cfg    *config.Config
	authUC auth.UseCase
}

func NewAuthHandlers(cfg *config.Config, authUC auth.UseCase) auth.Handlers {
	return &authHandlers{cfg: cfg, authUC: authUC}
}

// Login
//
//	@Summary		Login user
//	@Description	Login user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.CreateUser	true	"Login user"
//	@Success		200		{object}	models.ResponseToken
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/auth/login [post]
func (a authHandlers) Login(g *gin.Context) {
	var newUser models.UserLogin
	if err := g.BindJSON(&newUser); err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	token, err := a.authUC.Login(context.Background(), newUser)
	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, token)
}

// Registration
//
//	@Summary		Registration user
//	@Description	Registration user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.CreateUser	true	"Add user"
//	@Success		200		{object}	models.ResponseToken
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/auth/registration [post]
func (a authHandlers) Registration(g *gin.Context) {
	var createUser models.CreateUser
	if err := g.BindJSON(&createUser); err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	token, err := a.authUC.Register(context.Background(), createUser)

	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, token)
}

// LogoutMe
//
//	@Security		ApiKeyAuth
//
//	@Summary		Logout user
//	@Description	Logout user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Message
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/auth/logout [post]
func (a authHandlers) LogoutMe(g *gin.Context) {
	accessTokenKey, _ := g.Get("accessTokenKey")
	err := a.authUC.Logout(context.Background(), fmt.Sprintf("%v", accessTokenKey))
	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, models.Message{Message: "logged out"})
}

// AuthMe auth : Auth me
//
//	@Security		ApiKeyAuth
//
//	@Summary		Auth me
//	@Description	Auth me
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Users
//	@Failure		500	{object}	models.ErrorModel
//	@Router			/auth/auth-me [get]
func (a authHandlers) AuthMe(g *gin.Context) {
	userID, _ := g.Get("userID")

	user, err := a.authUC.AuthMe(context.Background(), fmt.Sprintf("%v", userID))

	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, user)
}

// RefreshToken Refresh : Refresh token.
//
//	@Summary		Refresh Token
//	@Description	Refresh Token
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		models.RefreshTokenModel	true	"Refresh token"
//	@Success		200		{object}	models.ResponseToken
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/auth/refresh-token [post]
func (a authHandlers) RefreshToken(g *gin.Context) {
	var refreshToken models.RefreshTokenModel
	if err := g.BindJSON(&refreshToken); err != nil {
		g.JSON(http.StatusBadRequest, models.ErrorModel{Error: err.Error()})
		return
	}

	newToken, err := a.authUC.Refresh(context.Background(), refreshToken.RefreshToken)

	if err != nil {
		g.JSON(http.StatusInternalServerError, models.ErrorModel{Error: err.Error()})
		return
	}

	g.JSON(http.StatusOK, newToken)
}
