package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mohar9h/go-graphql-ws-api/internal/handler"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	auth := routerGroup.Group("/auth")

	auth.POST("/login", handler.Login)

	auth.POST("/register", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Register successful"})
	})

	auth.POST("/logout", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Logout successful"})
	})

	auth.GET("/refresh_token", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Refresh token successful"})
	})
}
