package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchMonos search mono data
func SearchMonos(c *gin.Context) {
	//samp
	c.JSON(http.StatusOK, "search")
}
