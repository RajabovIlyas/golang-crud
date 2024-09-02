package usersController

import "github.com/gin-gonic/gin"

type Controller interface {
	GetUsers(c *gin.Context)
	GetUser(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	UpdateUserPassword(c *gin.Context)
}
