/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package middleware

import (
	"github.com/gin-gonic/gin"
	v1alpha "github.com/hvxahv/hvx/api/grpc/proto/device/v1alpha1"
	v "github.com/hvxahv/hvx/pkg/microsvc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hvxahv/hvx/pkg/identity"
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

	conn, err := grpc.DialContext(c, v.GetGRPCServiceAddress("device"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		c.JSON(500, gin.H{
			"code":    "500",
			"message": "HTTP_INTERNAL_SERVER_ERROR",
		})
		c.Abort()
		return
	}
	defer conn.Close()

	client := v1alpha.NewDevicesClient(conn)
	exist, err := client.DeviceIsExistByID(c, &v1alpha.DeviceIsExistByIDRequest{
		Id: pares.DeviceID,
	})
	if err != nil {
		c.JSON(400, gin.H{
			"code":    "400",
			"message": "AUTH_TOKEN_INVALID",
		})
		c.Abort()
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

	c.Next()
}
