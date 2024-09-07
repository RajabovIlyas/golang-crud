package usecase

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	utils2 "github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
)

type tokenUS struct {
	cfg       *config.Config
	tokenRepo token.Repository
}

func NewTokenUseCase(cfg *config.Config, tokenRepo token.Repository) token.UseCase {
	return &tokenUS{cfg, tokenRepo}
}

func (t tokenUS) GenerateToken(ctx context.Context, userID uuid.UUID) (models.ResponseToken, error) {
	newToken, err := t.tokenRepo.Create(ctx, userID)

	if err != nil {
		return models.ResponseToken{}, err
	}

	return t.GenerateTokenByToken(models.GenerateTokenModel{ID: newToken.ID, AccessTokenKey: newToken.AccessTokenKey})
}

func (t tokenUS) GenerateTokenByToken(generateToken models.GenerateTokenModel) (models.ResponseToken, error) {
	at, err := utils2.GenerateToken[uuid.UUID](generateToken.AccessTokenKey, t.cfg.Server.JwtSecretKey, &t.cfg.Server.TokenExpiresIn)

	if err != nil {
		return models.ResponseToken{}, err
	}

	rt, err := utils2.GenerateToken[uuid.UUID](generateToken.ID, t.cfg.Server.JwtSecretKey, &t.cfg.Server.TokenRefreshExpiresIn)

	if err != nil {
		return models.ResponseToken{}, err
	}

	return models.ResponseToken{AccessToken: at, RefreshToken: rt}, nil
}

func (t tokenUS) FindTokenById(ctx context.Context, tokenIDStr string) (database.FindTokenByIdRow, error) {
	tokenID, err := uuid.Parse(tokenIDStr)
	if err != nil {
		return database.FindTokenByIdRow{}, err
	}
	return t.tokenRepo.FindByID(ctx, tokenID)
}

func (t tokenUS) FindTokenByAccessKey(ctx context.Context, accessKeyStr string) (database.FindTokenByAccessKeyRow, error) {
	accessKey, err := uuid.Parse(accessKeyStr)
	if err != nil {
		return database.FindTokenByAccessKeyRow{}, err
	}
	return t.tokenRepo.FindByAccessKey(ctx, accessKey)
}

func (t tokenUS) CreateToken(ctx context.Context, userID uuid.UUID) (database.CreateTokenRow, error) {
	return t.tokenRepo.Create(ctx, userID)
}

func (t tokenUS) UpdateToken(ctx context.Context, tokenID uuid.UUID) (database.UpdateTokenByIdRow, error) {
	return t.tokenRepo.UpdateByID(ctx, tokenID)
}

func (t tokenUS) DeleteTokenById(ctx context.Context, tokenID uuid.UUID) error {
	return t.tokenRepo.DeleteByID(ctx, tokenID)
}

func (t tokenUS) DeleteTokenByAccessKey(ctx context.Context, accessKey uuid.UUID) error {
	return t.tokenRepo.DeleteByAccessKey(ctx, accessKey)
}

func (t tokenUS) DeleteOldTokens(ctx context.Context) error {
	return t.tokenRepo.DeleteOldTokens(ctx)
}
