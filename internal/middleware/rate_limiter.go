package middleware

import (
	"net/http"

	"github.com/didip/tollbooth"
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/go-graphql-ws-api/internal/config"
	"github.com/mohar9h/go-graphql-ws-api/internal/utils"
)

func RateLimiter() gin.HandlerFunc {
	cfg := config.GetConfig()
	limiter := tollbooth.NewLimiter(float64(cfg.RateLimit.RequestsPerSecond), nil)
	limiter.SetBurst(cfg.RateLimit.BurstSize)

	return func(c *gin.Context) {
		err := tollbooth.LimitByRequest(limiter, c.Writer, c.Request)
		if err != nil {
			utils.Error(c.Writer, http.StatusTooManyRequests, "Too many requests", err.Error())
			c.Abort()
			return
		}
		c.Next()
	}
}
