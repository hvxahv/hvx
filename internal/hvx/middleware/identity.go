package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// GetUsername Get the username of the context login user through the username key.
func GetUsername(c *gin.Context) string {
	name, ok := c.Get("username")
	if !ok {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "failed to get name by token!",
		})
		return ""
	}

	author, ok := name.(string)
	if !ok {
		log.Println("failed to convert username to string.")
	}
	return author
}

func GetDevicesID(c *gin.Context) string {
	device, ok := c.Get("devices")
	if !ok {
		c.JSON(500, gin.H{
			"code":    500,
			"message": "failed to get name by token!",
		})
		return ""
	}

	d, ok := device.(string)
	if !ok {
		log.Println("failed to convert username to string.")
	}
	return d
}
