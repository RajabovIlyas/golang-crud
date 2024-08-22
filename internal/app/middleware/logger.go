package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

func Logger(c *gin.Context) {
	t := time.Now()
	c.Next()

	if c.Writer.Status() >= http.StatusBadRequest {
		log.Error().Fields(map[string]interface{}{
			"error": c.Errors.ByType(gin.ErrorTypePrivate).String(),
		}).Msg("Response")
		return
	}
	log.Info().Fields(map[string]interface{}{
		"status":  c.Writer.Status(),
		"latency": time.Since(t).String(),
		"method":  c.Request.Method,
		"uri":     c.Request.URL.Path,
	}).Msg("Request")
}
