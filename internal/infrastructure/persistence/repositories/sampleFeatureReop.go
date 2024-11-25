package sampleFeatureRepo

import (
	sampleFeatureEntity "CoreBaseGo/internal/domain/sampleFeature/entity"
	"fmt"
	"gorm.io/gorm"
)

// StoreSampleFeature StorePlan creates a new plan in the database
func StoreSampleFeature(db *gorm.DB, sampleFeature *sampleFeatureEntity.SampleFeature) (*sampleFeatureEntity.SampleFeature, error) {
	if err := db.Create(sampleFeature).Error; err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}
	return sampleFeature, nil
}

// GetSampleFeature retrieves all plans
func GetSampleFeature(db *gorm.DB) ([]sampleFeatureEntity.SampleFeature, error) {
	var sampleFeature []sampleFeatureEntity.SampleFeature
	if err := db.Find(&sampleFeature).Error; err != nil {
		return nil, fmt.Errorf("failed to fetch plans: %w", err)
	}
	return sampleFeature, nil
}
