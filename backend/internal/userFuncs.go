package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetInfo get information about user account
func GetInfo(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "get info")
}

// Login user login
func Login(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "login")
}

// Logout user logout
func Logout(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "logout")
}

// CreateAccount create user account
func CreateAccount(c *gin.Context) {
	//samp
	c.JSON(http.StatusOK, "create account")
}

// UpdateAccount update user information
func UpdateAccount(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "update account")
}

// DeleteAccount delete user account
func DeleteAccount(c *gin.Context) {
	// samp
	c.JSON(http.StatusOK, "delete account")
}
