/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
)

func AuthAccountHandler(c *gin.Context) {
	cli, err := account.NewAuthClient()
	if err != nil {
		return
	}
	d := &pb.VerifyRequest{
		Username: c.PostForm("username"),
		Password: c.PostForm("password"),
		Ua:       c.GetHeader("User-Agent"),
	}

	verify, err := cli.Verify(c, d)
	if err != nil {
		c.JSON(401, gin.H{
			"code":  "401",
			"reply": "UNAUTHORIZED",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":       "200",
		"token":      verify.Token,
		"mail":       verify.Mail,
		"device_id":  verify.DeviceId,
		"public_key": verify.PublicKey,
	})
}

func LogoutHandler(c *gin.Context) {
	client, err := device.GetDeviceClient()
	if err != nil {
		return
	}
	reply, err := client.DeleteDeviceByID(c, &v1alpha1.DeleteDeviceByIDRequest{
		AccountId: middleware.GetAccountID(c),
		DeviceId:  middleware.GetDeviceID(c),
	})
	if err != nil {
		return
	}
	c.JSON(200, gin.H{
		"code":  "200",
		"reply": reply.Reply,
	})
}
