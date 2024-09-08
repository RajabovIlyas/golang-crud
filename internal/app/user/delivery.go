package user

import "github.com/gin-gonic/gin"

type Handlers interface {
	GetUsers(g *gin.Context)
	GetUser(g *gin.Context)
	UpdateUser(g *gin.Context)
	DeleteUser(g *gin.Context)
	UpdateUserPassword(g *gin.Context)
}
