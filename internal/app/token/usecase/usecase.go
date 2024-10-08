package usecase

import (
	"context"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	utils2 "github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type tokenUS struct {
	cfg       *config.Config
	tokenRepo token.Repository
	redisRepo token.RedisRepository
	logger    zerolog.Logger
}

func NewTokenUseCase(cfg *config.Config, tokenRepo token.Repository, redisRepo token.RedisRepository, logger zerolog.Logger) token.UseCase {
	return &tokenUS{cfg, tokenRepo, redisRepo, logger}
}

func (t tokenUS) GenerateToken(ctx context.Context, userID uuid.UUID) (models.ResponseToken, error) {
	newToken, err := t.tokenRepo.Create(userID)

	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.GenerateToken: create a new token to db")
		return models.ResponseToken{}, err
	}

	generatedTokens, err := t.GenerateTokenByToken(models.GenerateTokenModel{ID: newToken.ID, AccessTokenKey: newToken.AccessTokenKey})
	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.GenerateToken: generate new tokens")
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

func (t tokenUS) FindTokenById(ctx context.Context, tokenIDStr string) (models.Tokens, error) {
	cachedToken, err := t.redisRepo.GetByIDCtx(ctx, tokenIDStr)
	if err != nil {
		t.logger.Warn().Err(err).Msgf("tokenUC.FindTokenById.GetByIDCtx: %v", err)
	}
	if cachedToken != nil {
		return *cachedToken, nil
	}

	tokenID, err := uuid.Parse(tokenIDStr)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenById(invalid tokenID): %s", tokenIDStr)
		return models.Tokens{}, err
	}

	foundToken, err := t.tokenRepo.FindByID(tokenID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenById: find a token by id=%v", tokenID)
		return models.Tokens{}, err
	}

	err = t.redisRepo.SetTokenCtx(ctx, foundToken.ID.String(), &foundToken)
	if err != nil {
		t.logger.Warn().Err(err).Msgf("tokenUC.FindTokenById.SetTokenCtx: %v", err)
	}

	return foundToken, nil
}

func (t tokenUS) FindTokenByAccessKey(ctx context.Context, accessKeyStr string) (models.Tokens, error) {
	cachedToken, err := t.redisRepo.GetByIDCtx(ctx, accessKeyStr)
	if err != nil {
		t.logger.Warn().Err(err).Msgf("tokenUC.FindTokenByAccessKey.GetByIDCtx: %v", err)
	}
	if cachedToken != nil {
		return *cachedToken, nil
	}

	accessKey, err := uuid.Parse(accessKeyStr)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenByAccessKey(invalid accessKey): %s", accessKeyStr)
		return models.Tokens{}, err
	}

	foundToken, err := t.tokenRepo.FindByAccessKey(accessKey)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.FindTokenByAccessKey: find token by accessKey = %v", accessKey)
	}

	err = t.redisRepo.SetTokenCtx(ctx, foundToken.AccessTokenKey.String(), &foundToken)
	if err != nil {
		t.logger.Warn().Err(err).Msgf("tokenUC.FindTokenByAccessKey.SetTokenCtx: %v", err)
	}

	return foundToken, nil
}

func (t tokenUS) CreateToken(ctx context.Context, userID uuid.UUID) (models.Tokens, error) {
	createdToken, err := t.tokenRepo.Create(userID)
	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.CreateToken: create a new token")
		return models.Tokens{}, err
	}

	return createdToken, nil
}

func (t tokenUS) UpdateToken(ctx context.Context, tokenID uuid.UUID) (models.Tokens, error) {
	updatedToken, err := t.tokenRepo.UpdateByID(tokenID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.UpdateToken: update token by tokenID = %v", tokenID)
		return models.Tokens{}, err
	}

	t.DeleteTokenFromRedis(ctx, updatedToken)

	return updatedToken, nil
}

func (t tokenUS) DeleteTokenById(ctx context.Context, tokenID uuid.UUID) error {
	deletedToken, err := t.tokenRepo.DeleteByID(tokenID)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.DeleteTokenById: delete token by tokenID = %v", tokenID)
		return err
	}

	t.DeleteTokenFromRedis(ctx, deletedToken)

	return nil
}

func (t tokenUS) DeleteTokenByAccessKey(ctx context.Context, accessKey uuid.UUID) error {
	deletedToken, err := t.tokenRepo.DeleteByAccessKey(accessKey)
	if err != nil {
		t.logger.Error().Err(err).Msgf("tokenUS.DeleteTokenByAccessKey: delete token by accessKey = %v", accessKey)
		return err
	}

	t.DeleteTokenFromRedis(ctx, deletedToken)

	return nil
}

func (t tokenUS) DeleteOldTokens(ctx context.Context) error {
	deletedTokens, err := t.tokenRepo.DeleteOldTokens()
	if err != nil {
		t.logger.Error().Err(err).Msg("tokenUS.DeleteOldTokens: delete Old tokens")
		return err
	}

	for _, deletedToken := range deletedTokens {
		t.DeleteTokenFromRedis(ctx, deletedToken)
	}

	return err
}

func (t *tokenUS) DeleteTokenFromRedis(ctx context.Context, token models.Tokens) {
	err := t.redisRepo.DeleteTokenCtx(ctx, token.ID.String())
	if err != nil {
		t.logger.Warn().Err(err).Msgf("tokenUC.DeleteTokenFromRedis.DeleteTokenCtx(ID): %s", err)
	}
	err = nil
	err = t.redisRepo.DeleteTokenCtx(ctx, token.AccessTokenKey.String())
	if err != nil {
		t.logger.Warn().Err(err).Msgf("tokenUC.DeleteTokenFromRedis.DeleteTokenCtx(AccessTokenKey): %s", err)
	}
}
