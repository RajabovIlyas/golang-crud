package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func (mw *MiddlewareManager) Logger(g *gin.Context) {
	t := time.Now()
	g.Next()

	if g.Writer.Status() >= http.StatusBadRequest {
		mw.logger.Error().Fields(map[string]interface{}{
			"method": g.Request.Method,
			"uri":    g.Request.URL.Path,
			"error":  g.Errors.ByType(gin.ErrorTypePrivate).String(),
		}).Msg("Response")
		return
	}
	mw.logger.Info().Fields(map[string]interface{}{
		"status":  g.Writer.Status(),
		"latency": time.Since(t).String(),
		"method":  g.Request.Method,
		"uri":     g.Request.URL.Path,
	}).Msg("Request")
}
