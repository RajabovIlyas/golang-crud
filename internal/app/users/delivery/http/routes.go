package http

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/users"
	"github.com/gin-gonic/gin"
)

func MapUsersRoutes(userGroups *gin.RouterGroup, uh users.Handlers) {
	userGroups.GET("/", uh.GetUsers)
	userGroups.POST("/password/:userID", uh.UpdateUserPassword)
	userGroups.GET("/:userID", uh.GetUser)
	userGroups.PUT("/:userID", uh.UpdateUser)
	userGroups.DELETE("/:userID", uh.DeleteUser)
}
