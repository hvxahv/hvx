package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// GetUsername Get the username of the context login user through the username key.
func GetUsername(c *gin.Context) string {
	name, ok := c.Get("username")
	if !ok {
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "FAILED_PARSING_TOKEN",
		})
		return ""
	}

	author, ok := name.(string)
	if !ok {
		log.Println("USERNAME_INCORRECT_FORMAT")
	}
	return author
}

func GetDevicesID(c *gin.Context) string {
	device, ok := c.Get("devices")
	if !ok {
		c.JSON(401, gin.H{
			"code":    "500",
			"message": "FAILED_PARSING_TOKEN",
		})
		return ""
	}

	d, ok := device.(string)
	if !ok {
		log.Println("DEVICE_INCORRECT_FORMAT")
	}
	return d
}

func GetAccountID(c *gin.Context) string {
	id, ok := c.Get("id")
	if !ok {
		c.JSON(401, gin.H{
			"code":    "401",
			"message": "FAILED_PARSING_TOKEN",
		})
		return ""
	}

	i, ok := id.(string)
	if !ok {
		log.Println("ACCOUNT_ID_INCORRECT_FORMAT")
	}
	return i
}
