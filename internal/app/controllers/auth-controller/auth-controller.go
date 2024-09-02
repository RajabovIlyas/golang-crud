package authController

import (
	"fmt"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/auth-service"
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
func (ac *AuthController) LogoutMe(c *gin.Context) {
	accessTokenKey, _ := c.Get("accessTokenKey")
	err := ac.as.Logout(c.Request, fmt.Sprintf("%v", accessTokenKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "logged out"})
}

func (ac *AuthController) AuthMe(c *gin.Context) {

	userID, _ := c.Get("userID")

	fmt.Println("userID", userID)

	user, err := ac.as.AuthMe(c.Request, fmt.Sprintf("%v", userID))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

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
