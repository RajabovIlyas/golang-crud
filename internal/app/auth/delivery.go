package auth

import "github.com/gin-gonic/gin"

type Handlers interface {
	Login(g *gin.Context)
	Registration(g *gin.Context)
	LogoutMe(g *gin.Context)
	AuthMe(g *gin.Context)
	RefreshToken(g *gin.Context)
}
