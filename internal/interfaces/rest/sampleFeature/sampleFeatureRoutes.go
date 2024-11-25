package sampleFeatureRoute

import (
	sampleFeatureController "CoreBaseGo/internal/interfaces/rest/sampleFeature/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(v1 *gin.RouterGroup) {
	// Feature routes
	test := v1.Group("/test")
	{
		test.GET("/", sampleFeatureController.List)
		test.POST("/", sampleFeatureController.Store)

	}
}
