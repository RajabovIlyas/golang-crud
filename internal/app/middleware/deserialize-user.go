package middleware

import (
	"github.com/RajabovIlyas/golang-crud/config"
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/tokens-service"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/users-service"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type DeserializeMiddleware struct {
	us *usersService.UsersService
	ts *tokensService.TokensService
	c  *config.Config
}

func NewDeserializeMiddleware(p *models.DBConfigParam) *DeserializeMiddleware {
	return &DeserializeMiddleware{usersService.NewUsersService(p), tokensService.NewTokensService(p), p.C}
}

func (dm *DeserializeMiddleware) DeserializeUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			token = fields[1]
		}

		if token == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{"You are not logged in"})
			return
		}

		accessKey, err := utils.ValidateToken[string](token, dm.c.Server.JwtSecretKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{"You access token is invalid"})
			return
		}

		accessToken, err := dm.ts.FindTokenByAccessKey(ctx.Request.Context(), accessKey)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{"You access token is invalid"})
			return
		}

		user, err := dm.us.FindUserById(ctx.Request.Context(), accessToken.UserID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, models.Message{"You access token is invalid"})
			return
		}

		ctx.Set("accessTokenKey", accessToken.AccessTokenKey)
		ctx.Set("userID", user.ID)
		ctx.Next()
	}
}
