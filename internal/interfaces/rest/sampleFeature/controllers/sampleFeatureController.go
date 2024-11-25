package sampleFeatureController

import (
	sampleFeatureApplication "CoreBaseGo/internal/application/sampleFeature"
	"CoreBaseGo/internal/interfaces/rest"
	"CoreBaseGo/internal/interfaces/rest/messages"
	"github.com/gin-gonic/gin"
	"net/http"
)

// List retrieves all plans
func List(c *gin.Context) {
	list, err := sampleFeatureApplication.ListSampleFeature()
	if err != nil {
		rest.JSONOutput(c, http.StatusInternalServerError, nil, messages.InternalServerError, err.Error())
		return
	}
	rest.JSONOutput(c, http.StatusOK, list, messages.Success, "Plans retrieved successfully")
}

type SampleFeatureInput struct {
	Name string `json:"name" binding:"required"`
}

// Store retrieves all plans
func Store(c *gin.Context) {
	var input SampleFeatureInput
	if err := c.ShouldBindJSON(&input); err != nil {
		rest.JSONOutput(c, http.StatusBadRequest, nil, messages.BadRequest, "Invalid input")
		return
	}

	sampleFeature, err := sampleFeatureApplication.CreateSampleFeature(input.Name)
	if err != nil {
		rest.JSONOutput(c, http.StatusInternalServerError, nil, messages.InternalServerError, err.Error())
		return
	}

	rest.JSONOutput(c, http.StatusOK, sampleFeature, messages.Success, "Feature created successfully")
}
