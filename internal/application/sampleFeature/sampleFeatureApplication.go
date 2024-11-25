package sampleFeatureApplication

import (
	sampleFeatureEntity "CoreBaseGo/internal/domain/sampleFeature/entity"
	"CoreBaseGo/internal/domain/sampleFeature/service"
	"CoreBaseGo/internal/infrastructure/persistence"
	sampleFeatureRepo "CoreBaseGo/internal/infrastructure/persistence/repositories"
	"CoreBaseGo/internal/interfaces/rest/messages"
	"CoreBaseGo/internal/utils"
	"log"
	"net/http"
)

// CreateSampleFeature ListSampleFeature request flow should be here
func CreateSampleFeature(name string) (*sampleFeatureEntity.SampleFeature, error) {
	if err := service.ValidateSampleFeatureInput(name); err != nil {
		log.Printf("validation failed: %v", err)
		utils.Out(http.StatusBadRequest, messages.BadRequest, "validation failed: "+err.Error())
	}
	sampleFeature := service.CreateSampleFeature(name)

	db, err := persistence.GetInstance()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		utils.Out(http.StatusInternalServerError, messages.InternalServerError, "Internal server error")
	}

	sampleFeature, err = sampleFeatureRepo.StoreSampleFeature(db, sampleFeature)
	if err != nil {
		log.Printf("Database error: %v", err)
		utils.Out(http.StatusInternalServerError, messages.InternalServerError, err.Error())
	}
	return sampleFeature, err

}

// ListSampleFeature request flow should be here
func ListSampleFeature() ([]sampleFeatureEntity.SampleFeature, error) {
	db, err := persistence.GetInstance()
	if err != nil {
		log.Printf("Database connection error: %v", err)
		utils.Out(http.StatusInternalServerError, messages.InternalServerError, "Internal server error")
		return nil, err
	}

	return sampleFeatureRepo.GetSampleFeature(db)
}
