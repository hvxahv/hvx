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
	"github.com/hvxahv/hvxahv/pkg/identity/middleware"
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

func DeleteAccountHandler(c *gin.Context) {
	username := middleware.GetUsername(c)
	password := c.PostForm("password")

	cli, err := account.NewAccountClient()
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	d, err := cli.Delete(c, &pb.DeleteRequest{
		Username: username,
		Password: password,
	})
	if err != nil {
		c.JSON(500, gin.H{
			"code":  "500",
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"code":  d.Code,
		"reply": d.Reply,
	})
}
