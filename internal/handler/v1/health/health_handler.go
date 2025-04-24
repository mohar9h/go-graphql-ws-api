package health

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/go-graphql-ws-api/internal/utils"
)

func HealthCheck(c *gin.Context) {
	utils.Success(c.Writer, gin.H{"status": "ok"}, "Service is healthy")
}
