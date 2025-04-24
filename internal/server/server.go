package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/go-graphql-ws-api/internal/middleware"
	"github.com/mohar9h/go-graphql-ws-api/internal/routes"
)

func SetupServer() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger(), middleware.RateLimiter())

	routes.RegisterRoutes(router)

	return router
}
