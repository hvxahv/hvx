/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/gateway/identity"

	"strings"
	"time"
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
			"message": "TOKEN IS NOT CARRIED IN THE REQUEST",
		})
		c.Abort()
		return
	}

	// Because the device ID is unique, when logging in,
	// the device id obtained by the token is used to query whether the device exists.
	// If the device does not exist, the device will be returned as unregistered.
	// This method is used to revoke the issued token when resetting the password or deleting the device.
	_, err := identity.VerifyToken(t)
	if err != nil {
		c.JSON(500, gin.H{
			"state":   "500",
			"message": "TOKEN_IS_INCORRECT",
		})
		c.Abort()
		return
	}

	pares, err := identity.ParseToken(t)
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "500",
			"message": "TOKEN_PARSING_FAILED",
		})
		c.Abort()
		return
	}

	client, err := clientv1.New(c,
		clientv1.SetEndpoints("hvxahv.disism.internal:50030"),
		clientv1.SetDialTimeout(10*time.Second),
	)
	exist, err := client.DeviceIsExistByID(c, &apis.DeviceIsExistByIDRequest{
		Id: pares.DeviceID,
	})
	if err != nil {
		return
	}
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "AUTH_TOKEN_INVALID",
		})
		c.Abort()
		return
	}

	if !exist.IsExist {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "AUTH_TOKEN_INVALID",
		})
		c.Abort()
		return
	}

	c.Next()
}
