package health

import (
	"github.com/gin-gonic/gin"
	healthHandler "github.com/mohar9h/go-graphql-ws-api/internal/handler/v1/health"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	healthGroup := routerGroup.Group("/health")
	healthGroup.GET("/", healthHandler.HealthCheck)

}
