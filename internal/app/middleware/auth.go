package middleware

import (
	"github.com/RajabovIlyas/golang-crud/internal/pkg/httpErrors"
	"github.com/RajabovIlyas/golang-crud/internal/pkg/utils"
	"github.com/gin-gonic/gin"
	"strings"
)

func (m *MiddlewareManager) AuthSessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userToken string
		authorizationHeader := ctx.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			userToken = fields[1]
		}

		if userToken == "" {
			ctx.AbortWithStatusJSON(httpErrors.ErrorResponse(httpErrors.InvalidJWTToken))
			return
		}

		accessKey, err := utils.ValidateToken[string](userToken, m.cfg.Server.JwtSecretKey)
		if err != nil {
			ctx.AbortWithStatusJSON(httpErrors.ErrorResponse(httpErrors.InvalidJWTToken))
			return
		}

		foundToken, err := m.tokenUC.FindTokenByAccessKey(ctx.Request.Context(), accessKey)

		if err != nil {
			ctx.AbortWithStatusJSON(httpErrors.ErrorResponse(httpErrors.InvalidJWTToken))
			return
		}

		foundUser, err := m.userUC.FindById(ctx.Request.Context(), foundToken.UserID.String())
		if err != nil {
			ctx.AbortWithStatusJSON(httpErrors.ErrorResponse(httpErrors.InvalidJWTToken))
			return
		}

		ctx.Set("accessTokenKey", foundToken.AccessTokenKey)
		ctx.Set("userID", foundUser.ID)
		ctx.Next()
	}
}
