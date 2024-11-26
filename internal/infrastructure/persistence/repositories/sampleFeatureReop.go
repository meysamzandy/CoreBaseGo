package sampleFeatureRepo

import (
	sampleFeatureEntity "CoreBaseGo/internal/domain/sampleFeature/entity"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

// StoreSampleFeature StorePlan creates a new plan in the database
func StoreSampleFeature(db *gorm.DB, sampleFeature *sampleFeatureEntity.SampleFeature) (*sampleFeatureEntity.SampleFeature, error) {
	if err := db.Create(sampleFeature).Error; err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return sampleFeature, nil
}

func GetSampleFeature(c *gin.Context, db *gorm.DB) paginate.Page {
	// Create a base query
	query := db.Model(&sampleFeatureEntity.SampleFeature{})

	// Initialize the pagination library
	pg := paginate.New()

	// Use the HTTP request from the Gin context
	req := c.Request

	// Execute pagination
	result := pg.With(query).Request(req).Response(&[]sampleFeatureEntity.SampleFeature{})

	return result
}
