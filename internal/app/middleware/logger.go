package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func Logger(g *gin.Context) {
	t := time.Now()
	g.Next()

	if g.Writer.Status() >= http.StatusBadRequest {
		log.Error().Fields(map[string]interface{}{
			"method": g.Request.Method,
			"uri":    g.Request.URL.Path,
			"error":  g.Errors.ByType(gin.ErrorTypePrivate).String(),
		}).Msg("Response")
		return
	}
	log.Info().Fields(map[string]interface{}{
		"status":  g.Writer.Status(),
		"latency": time.Since(t).String(),
		"method":  g.Request.Method,
		"uri":     g.Request.URL.Path,
	}).Msg("Request")
}
