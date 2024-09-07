package middleware

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/gin-gonic/gin"
)

type MiddlewareManager interface {
	AuthSessionMiddleware() gin.HandlerFunc
}

type middlewareManager struct {
	cfg     *config.Config
	userUC  user.UseCase
	tokenUC token.UseCase
}

func NewMiddlewareManager(cfg *config.Config, userUC user.UseCase, tokenUC token.UseCase) MiddlewareManager {
	return &middlewareManager{cfg, userUC, tokenUC}
}
