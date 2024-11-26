package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/morkid/paginate"
	"gorm.io/gorm"
)

func ListQueryWithPagination(c *gin.Context, db *gorm.DB, model interface{}) paginate.Page {
	// Create a base query using the passed model
	query := db.Model(model)

	// Initialize the pagination library
	pg := paginate.New()

	// Use the HTTP request from the Gin context
	req := c.Request

	// Execute pagination
	result := pg.With(query).Request(req).Response(model)
	return result
}
