package middleware

import (
	"github.com/RajabovIlyas/golang-crud/internal/app/models"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func (m *middlewareManager) AuthSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userToken string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			userToken = fields[1]
		}

		if userToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{"You are not logged in"})
			return
		}

		accessKey, err := utils.ValidateToken[string](userToken, m.cfg.Server.JwtSecretKey)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{"You access token is invalid"})
			return
		}

		foundToken, err := m.tokenUC.FindTokenByAccessKey(ctx.Request.Context(), accessKey)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.Message{"You access token is invalid"})
			return
		}

		foundUser, err := m.userUC.FindById(ctx.Request.Context(), foundToken.UserID.String())
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusForbidden, models.Message{"You access token is invalid"})
			return
		}

		ctx.Set("accessTokenKey", foundToken.AccessTokenKey)
		ctx.Set("userID", foundUser.ID)
		ctx.Next()
	}
}
