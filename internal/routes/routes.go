package routes

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/mohar9h/go-graphql-ws-api/internal/routes/v1"
)

func RegisterRoutes(router *gin.Engine) {
	apiGroup := router.Group("/api")
	v1.RegisterRoutes(apiGroup)
}
