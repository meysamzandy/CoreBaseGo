package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func Pagination(c *gin.Context) (int, int) {
	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset := (page - 1) * limit
	return limit, offset
}
