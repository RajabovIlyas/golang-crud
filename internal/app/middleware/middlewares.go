package middleware

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/rs/zerolog"
)

type MiddlewareManager struct {
	cfg     *config.Config
	userUC  user.UseCase
	tokenUC token.UseCase
	logger  zerolog.Logger
}

func NewMiddlewareManager(cfg *config.Config, userUC user.UseCase, tokenUC token.UseCase, logger zerolog.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg, userUC, tokenUC, logger}
}
