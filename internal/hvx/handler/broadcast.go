/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/channel"
	"github.com/hvxahv/hvxahv/pkg/identity/middleware"
)

func CreateBroadcastHandler(c *gin.Context) {
	client, err := channel.GetChannelBroadcastClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	broadcast, err := client.CreateBroadcast(c, &v1alpha1.CreateBroadcastRequest{
		ChannelId: c.PostForm("channel_id"),
		AdminId:   middleware.GetAccountID(c),
		ArticleId: c.PostForm("article_id"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, broadcast)
}

func GetBroadcastsOfChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelBroadcastClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	broadcasts, err := client.GetAllBroadcasts(c, &v1alpha1.GetAllBroadcastsRequest{
		ChannelId: c.Param("id"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, broadcasts)
}

func DeleteBroadcastHandler(c *gin.Context) {
	client, err := channel.GetChannelBroadcastClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	broadcast, err := client.DeleteBroadcast(c, &v1alpha1.DeleteBroadcastRequest{
		ChannelId:   c.PostForm("channel_id"),
		AdminId:     middleware.GetAccountID(c),
		BroadcastId: c.Param("id"),
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	c.JSON(200, broadcast)
}
