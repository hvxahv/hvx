package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/internal/hvx/policy"
	"github.com/pkg/errors"
	"log"
	"strings"
)

// Auth authentication middleware for gin network framework,
// Check whether the Token is carried in the request, and verify whether the Token is correct,
// Will obtain the username by parsing the Token and add the username in the context and set the key to loginUser.
func Auth(c *gin.Context) {
	ht := c.Request.Header.Get("Authorization")
	t := strings.Split(ht, "Bearer ")[1]
	if ht == "" {
		c.JSON(500, gin.H{
			"state":   "500",
			"message": "TOKEN IS NOT CARRIED IN THE REQUEST.",
		})
		c.Abort()
		return
	}

	// Because the device ID is unique, when logging in,
	// the device id obtained by the token is used to query whether the device exists.
	// If the device does not exist, the device will be returned as unregistered.
	// This method is used to revoke the issued token when resetting the password or deleting the device.
	_, err := policy.VerifyToken(t)
	if err != nil {
		c.JSON(500, gin.H{
			"state":   "500",
			"message": "LOGIN FAILED. TOKEN IS INCORRECT!",
		})
		c.Abort()
	} else {
		u, err := policy.ParseToken(t)
		if err != nil {
			log.Println("failed to obtain user through token.")
			c.Abort()
		}
		if device.NewDevicesIsNotExist(u.DevicesID).IsNotExist() {
			fmt.Println(errors.Errorf("THE_DEVICE_HAS_BEEN_BANNED"))
			c.JSON(401, gin.H{
				"code":    "401",
				"message": "Token is not available, the device has been banned.",
			})
			c.Abort()
		}
		c.Set("devices", u.DevicesID)
		c.Set("username", u.User)
		c.Next()
	}

}
