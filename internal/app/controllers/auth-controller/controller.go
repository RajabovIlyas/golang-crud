package authController

import "github.com/gin-gonic/gin"

type Controller interface {
	Login(c *gin.Context)
	Registration(c *gin.Context)
	LogoutMe(c *gin.Context)
	AuthMe(c *gin.Context)
	RefreshToken(c *gin.Context)
}
