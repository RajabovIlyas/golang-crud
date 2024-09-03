package middleware

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/common"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/tokens-service"
	"github.com/RajabovIlyas/golang-crud/internal/app/services/users-service"
	"github.com/RajabovIlyas/golang-crud/internal/app/utils"
	"github.com/RajabovIlyas/golang-crud/internal/database"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type DeserializeMiddleware struct {
	us *usersService.UsersService
	ts *tokensService.TokensService
}

func NewDeserializeMiddleware(db *database.Queries) *DeserializeMiddleware {
	return &DeserializeMiddleware{usersService.NewUsersService(db), tokensService.NewTokensService(db)}
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
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config, _ := common.GetConfig(".")
		accessKey, err := utils.ValidateToken[string](token, config.TokenSecret)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		accessToken, err := dm.ts.FindTokenByAccessKey(ctx.Request.Context(), accessKey)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		user, err := dm.us.FindUserById(ctx.Request.Context(), accessToken.UserID)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
			return
		}

		ctx.Set("accessTokenKey", accessToken.AccessTokenKey)
		ctx.Set("userID", user.ID)
		ctx.Next()
	}
}
