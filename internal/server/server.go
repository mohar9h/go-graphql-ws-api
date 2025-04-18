package server

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/pecalets_backend/internal/routes"
)

func SetupServer() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Recovery(), gin.Logger())

	routes.RegisterRoutes(router)

	return router
}
