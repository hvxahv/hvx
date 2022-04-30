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

func AddAdminToChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelAdminClient()
	if err != nil {
		return
	}
	cc, err := client.AddAdministrator(c, &v1alpha1.AddAdministratorRequest{
		ChannelId:      c.PostForm("channel_id"),
		AdminAccountId: middleware.GetAccountID(c),
		AddAdminId:     c.PostForm("admin_id"),
		IsOwner:        false,
	})
	if err != nil {
		return
	}
	c.JSON(200, cc)
}

func RemoveAdminFromChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelAdminClient()
	if err != nil {
		return
	}
	cc, err := client.RemoveAdministrator(c, &v1alpha1.RemoveAdministratorRequest{
		OwnerId:       middleware.GetAccountID(c),
		ChannelId:     c.PostForm("channel_id"),
		RemoveAdminId: c.PostForm("admin_id"),
	})
	if err != nil {
		return
	}
	c.JSON(200, cc)
}

func GetAdminsOfChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelAdminClient()
	if err != nil {
		return
	}
	cc, err := client.GetAdministrators(c, &v1alpha1.GetAdministratorsRequest{
		ChannelId: c.Param("id"),
		AccountId: middleware.GetAccountID(c),
	})
	if err != nil {
		return
	}
	c.JSON(200, cc)
}
