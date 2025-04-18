package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	auth := routerGroup.Group("/auth")

	auth.POST("/login", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	})

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
