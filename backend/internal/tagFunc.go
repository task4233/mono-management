package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTags get tag
func GetTags(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "get tags")
}

// CreateTag create tag
func CreateTag(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "crate tag")
}

// UpdateTag update tag
func UpdateTag(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "update tag")
}

// DeleteTag delete tag
func DeleteTag(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "delete tag")
}
