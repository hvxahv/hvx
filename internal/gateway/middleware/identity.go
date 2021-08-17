package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"log"
)

// GetUserName Get the username of the context login user through the loginUser key.
func GetUserName(c *gin.Context) (string, error) {
	name, ok := c.Get("loginUser")
	if !ok {
		log.Println("failed to get username")
		return "", errors.Errorf("failed to get username")
	}

	author, ok := name.(string)
	if !ok {
		log.Println("failed to convert username to string.")
	}

	return author, nil
}
