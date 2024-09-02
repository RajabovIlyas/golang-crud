package authController

import (
	"fmt"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/auth-service"
	"github.com/RajabovIlyas/golang-crud/internal/app/utils"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	as *authService.AuthService
}

func NewAuthController(db *database.Queries) *AuthController {
	return &AuthController{authService.NewAuthService(db)}
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
func (ac *AuthController) Login(c *gin.Context) {
	var newUser models.UserLogin
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := ac.as.Login(c.Request, newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"tokens": token})
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
func (ac *AuthController) Registration(c *gin.Context) {
	var newUser models.CreateUser
	if err := c.BindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := ac.as.Register(c.Request, newUser)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

// LogoutMe .
//
// @Security ApiKeyAuth
//
//	@Summary		Logout user
//	@Description	Logout user
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.Message
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/auth/logout [post]
func (ac *AuthController) LogoutMe(c *gin.Context) {
	accessTokenKey, _ := c.Get("accessTokenKey")
	err := ac.as.Logout(c.Request, fmt.Sprintf("%v", accessTokenKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

// AuthMe auth : Auth me
//
// @Security ApiKeyAuth
//
//	@Summary		Auth me
//	@Description	Auth me
//	@Tags			auth
//	@Accept			json
//	@Produce		json
//	@Success		200		{object}	models.ResponseUser
//	@Failure		500		{object}	models.ErrorModel
//	@Router			/auth/auth-me [get]
func (ac *AuthController) AuthMe(c *gin.Context) {

	userID, _ := c.Get("userID")

	user, err := ac.as.AuthMe(c.Request, fmt.Sprintf("%v", userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": utils.DatabaseResponseUserToResponseUser(models.ResponseUser(user))})
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
func (ac *AuthController) RefreshToken(c *gin.Context) {
	var refreshToken models.RefreshTokenModel
	if err := c.BindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newToken, err := ac.as.Refresh(c.Request, refreshToken.RefreshToken)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})
}
