package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/go-graphql-ws-api/internal/routes/v1/auth"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	v1Group := routerGroup.Group("/v1")

	auth.RegisterRoutes(v1Group)
}
