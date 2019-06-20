package internal

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMonos get
func GetMonos(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "get monos")
}

// CreateMono create
func CreateMono(c *gin.Context) {
	// samp
	var user User
	c.Bind(&user)

	log.Println(user)

	c.JSON(http.StatusOK, "create mono")
}

// UpdateMonos update
func UpdateMonos(c *gin.Context) {
	// samp
	var user User
	c.Bind(&user)

	log.Println(user)

	c.JSON(http.StatusOK, "update mono")
}

// DeleteMono delete
func DeleteMono(c *gin.Context) {
	// samp
	var user User
	c.Bind(&user)

	log.Println(user)

	c.JSON(http.StatusOK, "delete mono")
}
