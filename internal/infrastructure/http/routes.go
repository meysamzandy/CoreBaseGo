package http

import (
	middlewares "CoreBaseGo/internal/infrastructure/http/middlewares/common"
	sampleFeatureRoute "CoreBaseGo/internal/interfaces/rest/sampleFeature"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	// Group API endpoints for version 1
	router.Use(middlewares.CORS())   // Apply CORS middleware
	router.Use(middlewares.Logger()) // Existing logger middleware
	v1 := router.Group("api/admin/v1")
	{
		sampleFeatureRoute.Routes(v1)
	}
}
