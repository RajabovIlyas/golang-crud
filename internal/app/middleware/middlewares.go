package middleware

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/tokens"
	"github.com/RajabovIlyas/golang-crud/internal/app/users"
	"github.com/rs/zerolog"
)

type MiddlewareManager struct {
	cfg     *config.Config
	userUC  users.UseCase
	tokenUC tokens.UseCase
	logger  zerolog.Logger
}

func NewMiddlewareManager(cfg *config.Config, userUC users.UseCase, tokenUC tokens.UseCase, logger zerolog.Logger) *MiddlewareManager {
	return &MiddlewareManager{cfg, userUC, tokenUC, logger}
}
