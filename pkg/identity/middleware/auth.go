/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/device"
	"github.com/hvxahv/hvxahv/pkg/identity"
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
	client, err := device.GetDeviceClient()
	if err != nil {
		return
	}

	exist, err := client.DeviceIsExistByID(c, &v1alpha1.DeviceIsExistByIDRequest{
		Id: pares.DeviceID,
	})
	if err != nil {
		return
	}

	if exist.IsExist {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "AUTH_TOKEN_INVALID",
		})
		c.Abort()
		return
	}

	c.Set("device_id", pares.DeviceID)
	c.Set("username", pares.Username)
	c.Set("id", pares.ID)
	c.Next()
}

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

func GetDeviceID(c *gin.Context) string {
	id, ok := c.Get("device_id")
	if !ok {
		c.JSON(401, gin.H{
			"code":    "500",
			"message": "FAILED_PARSING_TOKEN",
		})
		return ""
	}

	d, ok := id.(string)
	if !ok {
		log.Println("DEVICE_ID_INCORRECT_FORMAT")
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
