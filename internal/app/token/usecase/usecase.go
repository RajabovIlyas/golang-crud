package usecase

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	utils2 "github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type tokenUS struct {
	cfg       *config.Config
	tokenRepo token.Repository
	logger    zerolog.Logger
}

func NewTokenUseCase(cfg *config.Config, tokenRepo token.Repository, logger zerolog.Logger) token.UseCase {
	return &tokenUS{cfg, tokenRepo, logger}
}

func (t tokenUS) GenerateToken(ctx context.Context, userID uuid.UUID) (models.ResponseToken, error) {
	newToken, err := t.tokenRepo.Create(ctx, userID)

	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.GenerateToken: create a new token")
		return models.ResponseToken{}, err
	}

	generatedTokens, err := t.GenerateTokenByToken(models.GenerateTokenModel{ID: newToken.ID, AccessTokenKey: newToken.AccessTokenKey})
	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.GenerateToken: generate new token")
		return models.ResponseToken{}, err
	}

	return generatedTokens, nil
}

func (t tokenUS) GenerateTokenByToken(generateToken models.GenerateTokenModel) (models.ResponseToken, error) {
	at, err := utils2.GenerateToken[uuid.UUID](generateToken.AccessTokenKey, t.cfg.Server.JwtSecretKey, &t.cfg.Server.TokenExpiresIn)

	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.GenerateTokenByToken: generate new AccessToken")
		return models.ResponseToken{}, err
	}

	rt, err := utils2.GenerateToken[uuid.UUID](generateToken.ID, t.cfg.Server.JwtSecretKey, &t.cfg.Server.TokenRefreshExpiresIn)
	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.GenerateTokenByToken: generate new RefreshToken")
		return models.ResponseToken{}, err
	}

	return models.ResponseToken{AccessToken: at, RefreshToken: rt}, nil
}

func (t tokenUS) FindTokenById(ctx context.Context, tokenIDStr string) (database.FindTokenByIdRow, error) {
	tokenID, err := uuid.Parse(tokenIDStr)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenById(invalid tokenID): %s", tokenIDStr)
		return database.FindTokenByIdRow{}, err
	}

	foundToken, err := t.tokenRepo.FindByID(ctx, tokenID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenById: find a token by id=%v", tokenID)
		return database.FindTokenByIdRow{}, err
	}

	return foundToken, nil
}

func (t tokenUS) FindTokenByAccessKey(ctx context.Context, accessKeyStr string) (database.FindTokenByAccessKeyRow, error) {
	accessKey, err := uuid.Parse(accessKeyStr)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenByAccessKey(invalid accessKey): %s", accessKeyStr)
		return database.FindTokenByAccessKeyRow{}, err
	}

	foundToken, err := t.tokenRepo.FindByAccessKey(ctx, accessKey)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenByAccessKey: find token by accessKey = %v", accessKey)
	}

	return foundToken, nil
}

func (t tokenUS) CreateToken(ctx context.Context, userID uuid.UUID) (database.CreateTokenRow, error) {
	createdToken, err := t.tokenRepo.Create(ctx, userID)
	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.CreateToken: create a new token")
		return database.CreateTokenRow{}, err
	}

	return createdToken, nil
}

func (t tokenUS) UpdateToken(ctx context.Context, tokenID uuid.UUID) (database.UpdateTokenByIdRow, error) {
	updatedToken, err := t.tokenRepo.UpdateByID(ctx, tokenID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.UpdateToken: update token by tokenID = %v", tokenID)
		return database.UpdateTokenByIdRow{}, err
	}

	return updatedToken, nil
}

func (t tokenUS) DeleteTokenById(ctx context.Context, tokenID uuid.UUID) error {
	err := t.tokenRepo.DeleteByID(ctx, tokenID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.DeleteTokenById: delete token by tokenID = %v", tokenID)
		return err
	}

	return nil
}

func (t tokenUS) DeleteTokenByAccessKey(ctx context.Context, accessKey uuid.UUID) error {
	err := t.tokenRepo.DeleteByAccessKey(ctx, accessKey)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.DeleteTokenByAccessKey: delete token by accessKey = %v", accessKey)
		return err
	}

	return nil
}

func (t tokenUS) DeleteOldTokens(ctx context.Context) error {
	err := t.tokenRepo.DeleteOldTokens(ctx)
	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.DeleteOldTokens: delete Old tokens")
		return err
	}
	return err
}
