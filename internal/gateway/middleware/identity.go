package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

// GetUserName Get the user name of the context login user through the loginUser key.
func GetUserName(c *gin.Context) string {
	name, ok := c.Get("loginUser")
	if !ok {
		log.Println("Failed to get username")
	}
	author, ok := name.(string)
	if !ok {
		log.Println("Failed to convert username to string")
	}

	return author
}
