package service

import (
	"CoreBaseGo/internal/domain/sampleFeature/entity"
	"errors"
)

func ValidateSampleFeatureInput(name string) error {
	if len(name) == 0 {
		return errors.New("name is required")
	}
	if len(name) > 100 {
		return errors.New("name cannot exceed 100 characters")
	}
	return nil
}

func CreateSampleFeature(name string) *entity.SampleFeature {
	return &entity.SampleFeature{
		Name: name,
	}
}
