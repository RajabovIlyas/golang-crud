package authService

import (
	"context"
	"errors"
	"github.com/RajabovIlyas/golang-crud/internal/app/common"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	tokensService "github.com/RajabovIlyas/golang-crud/internal/app/services/tokens-service"
	usersService "github.com/RajabovIlyas/golang-crud/internal/app/services/users-service"
	"github.com/RajabovIlyas/golang-crud/internal/app/utils"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/google/uuid"
)

type AuthService struct {
	ts *tokensService.TokensService
	us *usersService.UsersService
	c  *common.Config
}

func NewAuthService(db *database.Queries) *AuthService {
	c, _ := common.GetConfig(".")
	return &AuthService{tokensService.NewTokensService(db), usersService.NewUsersService(db), &c}
}

func (as *AuthService) GenerateTokenByToken(token models.GenerateTokenModel) (models.ResponseToken, error) {
	at, err := utils.GenerateToken[uuid.UUID](token.AccessTokenKey, as.c.TokenSecret, &as.c.TokenExpiresIn)

	if err != nil {
		return models.ResponseToken{}, err
	}

	rt, err := utils.GenerateToken[uuid.UUID](token.ID, as.c.TokenSecret, &as.c.TokenRefreshExpiresIn)

	if err != nil {
		return models.ResponseToken{}, err
	}

	return models.ResponseToken{AccessToken: at, RefreshToken: rt}, nil
}

func (as *AuthService) GenerateToken(c context.Context, userID uuid.UUID) (models.ResponseToken, error) {
	token, err := as.ts.CreateToken(c, userID)

	if err != nil {
		return models.ResponseToken{}, err
	}

	return as.GenerateTokenByToken(models.GenerateTokenModel{ID: token.ID, AccessTokenKey: token.AccessTokenKey})
}

func (as *AuthService) Login(c context.Context, userLogin models.UserLogin) (models.ResponseToken, error) {
	user, err := as.us.FindUserByUserName(c, userLogin.Username)
	if err != nil {
		return models.ResponseToken{}, err
	}

	err = utils.VerifyPassword(user.Password, userLogin.Password)
	if err != nil {
		return models.ResponseToken{}, errors.New("invalid username or password")
	}

	return as.GenerateToken(c, user.ID)
}

func (as *AuthService) Logout(c context.Context, accessTokenKey string) error {
	return as.ts.DeleteTokenByAccessKey(c, accessTokenKey)
}

func (as *AuthService) Register(c context.Context, newUser models.CreateUser) (models.ResponseToken, error) {
	user, err := as.us.Register(c, newUser)
	if err != nil {
		return models.ResponseToken{}, err
	}

	return as.GenerateToken(c, user.ID)
}

func (as *AuthService) Refresh(c context.Context, refreshToken string) (models.ResponseToken, error) {
	tokenID, err := utils.ValidateToken[string](refreshToken, as.c.TokenSecret)

	token, err := as.ts.UpdateToken(c, tokenID)
	if err != nil {
		return models.ResponseToken{}, err
	}

	return as.GenerateTokenByToken(models.GenerateTokenModel{ID: token.ID, AccessTokenKey: token.AccessTokenKey})
}

func (as *AuthService) AuthMe(c context.Context, userIDStr string) (database.FindUserByIdRow, error) {
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return database.FindUserByIdRow{}, err
	}

	return as.us.FindUserById(c, userID)
}
