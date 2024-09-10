package http

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/auth"
	"github.com/RajabovIlyas/golang-crud/internal/app/middleware"
	"github.com/gin-gonic/gin"
)

func MapAuthRoutes(authGroups *gin.RouterGroup, uh auth.Handlers, m *middleware.MiddlewareManager) {
	authGroups.POST("/login", uh.Login)
	authGroups.POST("/registration", uh.Registration)
	authGroups.POST("/logout", m.AuthSessionMiddleware(), uh.LogoutMe)
	authGroups.GET("/auth-me", m.AuthSessionMiddleware(), uh.AuthMe)
	authGroups.POST("/refresh-token", uh.RefreshToken)
}
