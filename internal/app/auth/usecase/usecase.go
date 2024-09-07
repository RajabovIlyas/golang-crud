package usecase

import (
	"context"
	"errors"
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/auth"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/token"
	"github.com/RajabovIlyas/golang-crud/internal/app/user"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/google/uuid"
)

type authUC struct {
	cfg     *config.Config
	userUC  user.UseCase
	tokenUC token.UseCase
}

func NewAuthUseCase(cfg *config.Config, userUC user.UseCase, tokenUC token.UseCase) auth.UseCase {
	return &authUC{cfg: cfg, userUC: userUC, tokenUC: tokenUC}
}

func (a authUC) Login(ctx context.Context, login models.UserLogin) (models.ResponseToken, error) {
	foundUser, err := a.userUC.FindByUsername(ctx, login.Username)
	if err != nil {
		return models.ResponseToken{}, err
	}

	err = utils.VerifyPassword(foundUser.Password, login.Password)
	if err != nil {
		return models.ResponseToken{}, errors.New("invalid username or password")
	}

	return a.tokenUC.GenerateToken(ctx, foundUser.ID)
}

func (a authUC) Logout(ctx context.Context, accessKeyStr string) error {
	accessKey, err := uuid.Parse(accessKeyStr)
	if err != nil {
		return err
	}
	return a.tokenUC.DeleteTokenByAccessKey(ctx, accessKey)
}

func (a authUC) Register(ctx context.Context, createUser models.CreateUser) (models.ResponseToken, error) {
	newUser, err := a.userUC.Create(ctx, createUser)
	if err != nil {
		return models.ResponseToken{}, err
	}

	return a.tokenUC.GenerateToken(ctx, newUser.ID)
}

func (a authUC) Refresh(ctx context.Context, refreshToken string) (models.ResponseToken, error) {
	tokenIDStr, err := utils.ValidateToken[string](refreshToken, a.cfg.Server.JwtSecretKey)

	tokenID, err := uuid.Parse(tokenIDStr)
	if err != nil {
		return models.ResponseToken{}, err
	}

	updatedToken, err := a.tokenUC.UpdateToken(ctx, tokenID)
	if err != nil {
		return models.ResponseToken{}, err
	}

	return a.tokenUC.GenerateTokenByToken(models.GenerateTokenModel{ID: updatedToken.ID, AccessTokenKey: updatedToken.AccessTokenKey})
}

func (a authUC) AuthMe(ctx context.Context, userIDStr string) (database.FindUserByIdRow, error) {
	return a.userUC.FindById(ctx, userIDStr)
}
