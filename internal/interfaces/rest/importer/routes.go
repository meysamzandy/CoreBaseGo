package routes

import (
	"CoreBaseGo/internal/infrastructure/http/Middlewares/common"
	"github.com/gin-gonic/gin"
)

func V1(router *gin.Engine) {
	// Group API endpoints for version 1
	router.Use(middlewares.CORS())   // Apply CORS middleware
	router.Use(middlewares.Logger()) // Existing logger middleware
	v1 := router.Group("api/admin/v1")
	{
		importer(v1)
	}
}

func importer(v1 *gin.RouterGroup) {
	// Feature routes
	planRoutes := v1.Group("/importer")
	{
		planRoutes.POST("/", planAdminControllers.CreatePlan) // Create

	}
}
