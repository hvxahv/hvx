/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/account"
)

func CreateAccountHandler(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	mail := c.PostForm("mail")

	// https://datatracker.ietf.org/doc/html/rfc5208
	publicKey := c.PostForm("public_key")

	d := &pb.CreateRequest{
		Username:  username,
		Mail:      mail,
		Password:  password,
		PublicKey: publicKey,
	}

	cli, err := account.NewAccountClient()
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	create, err := cli.Create(c, d)
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  create.Code,
		"reply": create.Reply,
	})
}

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
