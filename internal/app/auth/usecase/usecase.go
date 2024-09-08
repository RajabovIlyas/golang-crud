package usecase

import (
	"context"
	"errors"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/auth"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog"
)

type authUC struct {
	cfg     *config.Config
	userUC  user.UseCase
	tokenUC token.UseCase
	logger  zerolog.Logger
}

func NewAuthUseCase(cfg *config.Config, userUC user.UseCase, tokenUC token.UseCase, logger zerolog.Logger) auth.UseCase {
	return &authUC{cfg: cfg, userUC: userUC, tokenUC: tokenUC, logger: logger}
}

func (a authUC) Login(ctx context.Context, login models.UserLogin) (models.ResponseToken, error) {
	foundUser, err := a.userUC.FindByUsername(ctx, login.Username)
	if err != nil {
		a.logger.Warn().Err(err).Str("username", login.Username).Msg("authUC.Login(user not found)")
		return models.ResponseToken{}, errors.New("invalid username or password")
	}

	err = utils.VerifyPassword(foundUser.Password, login.Password)
	if err != nil {
		a.logger.Warn().Err(err).Str("username", login.Username).Msg("authUC.Login(invalid username or password)")
		return models.ResponseToken{}, errors.New("invalid username or password")
	}

	generatedToken, err := a.tokenUC.GenerateToken(ctx, foundUser.ID)
	if err != nil {
		a.logger.Error().Err(err).Str("username", foundUser.Username).Msg("authUC.Login(generate token)")
		return models.ResponseToken{}, errors.New("invalid username or password")
	}

	return generatedToken, nil
}

func (a authUC) Logout(ctx context.Context, accessKeyStr string) error {
	accessKey, err := uuid.Parse(accessKeyStr)
	if err != nil {
		a.logger.Error().Err(err).Msgf("authUC.Logout(invalid accessKey): %s", accessKeyStr)
		return err
	}
	err = a.tokenUC.DeleteTokenByAccessKey(ctx, accessKey)
	if err != nil {
		a.logger.Error().Err(err).Msg("authUC.Logout(delete token by accessKey)")
		return err
	}
	return nil
}

func (a authUC) Register(ctx context.Context, createUser models.CreateUser) (models.ResponseToken, error) {
	newUser, err := a.userUC.Create(ctx, createUser)
	if err != nil {
		a.logger.Error().Err(err).Msg("authUC.Register: when try to create user")
		return models.ResponseToken{}, err
	}

	generatedToken, err := a.tokenUC.GenerateToken(ctx, newUser.ID)
	if err != nil {
		a.logger.Error().Err(err).Str("username", newUser.Username).Msg("authUC.Login(generate token)")
		return models.ResponseToken{}, errors.New("invalid username or password")
	}

	return generatedToken, nil
}

func (a authUC) Refresh(ctx context.Context, refreshToken string) (models.ResponseToken, error) {
	tokenIDStr, err := utils.ValidateToken[string](refreshToken, a.cfg.Server.JwtSecretKey)

	tokenID, err := uuid.Parse(tokenIDStr)
	if err != nil {
		a.logger.Error().Err(err).Msgf("authUC.Refresh(invalid tokenID): %s", tokenIDStr)
		return models.ResponseToken{}, err
	}

	updatedToken, err := a.tokenUC.UpdateToken(ctx, tokenID)
	if err != nil {
		a.logger.Error().Err(err).Msg("authUC.Login(update token)")
		return models.ResponseToken{}, err
	}

	generatedToken, err := a.tokenUC.GenerateTokenByToken(models.GenerateTokenModel{ID: updatedToken.ID, AccessTokenKey: updatedToken.AccessTokenKey})
	if err != nil {
		a.logger.Error().Err(err).Msg("authUC.Login(generate token)")
		return models.ResponseToken{}, err
	}

	return generatedToken, nil
}

func (a authUC) AuthMe(ctx context.Context, userIDStr string) (models.ResponseUser, error) {
	foundUser, err := a.userUC.FindById(ctx, userIDStr)
	if err != nil {
		a.logger.Error().Err(err).Msgf("authUC.AuthMe(invalid userID): %s", userIDStr)
		return foundUser, err
	}

	return foundUser, nil
}
