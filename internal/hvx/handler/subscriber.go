/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/channel"
	"github.com/hvxahv/hvxahv/pkg/identity/middleware"
)

func AddSubscriberToChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelSubscriberClient()
	if err != nil {
		return
	}
	subscriber, err := client.AddSubscriber(c, &v1alpha1.AddSubscriberRequest{
		ChannelId: c.PostForm("channel_id"),
		AccountId: middleware.GetAccountID(c),
	})
	if err != nil {
		return
	}
	c.JSON(200, subscriber)
}

func UnsubscribeFromChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelSubscriberClient()
	if err != nil {
		return
	}
	subscriber, err := client.Unsubscribe(c, &v1alpha1.UnsubscribeRequest{
		ChannelId: c.Param("id"),
		AccountId: middleware.GetAccountID(c),
	})
	if err != nil {
		return
	}
	c.JSON(200, subscriber)
}

func RemoveSubscriberFromChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelSubscriberClient()
	if err != nil {
		return
	}
	subscriber, err := client.RemoveSubscriber(c, &v1alpha1.RemoveSubscriberRequest{
		AdminId:      middleware.GetAccountID(c),
		ChannelId:    c.Param("id"),
		SubscriberId: c.PostForm("subscriber_id"),
	})
	if err != nil {
		return
	}
	c.JSON(200, subscriber)
}

func GetSubscribersOfChannelHandler(c *gin.Context) {
	client, err := channel.GetChannelSubscriberClient()
	if err != nil {
		return
	}
	subscribers, err := client.GetAllSubscribers(c, &v1alpha1.GetAllSubscribersRequest{
		ChannelId: c.Param("id"),
		AdminId:   middleware.GetAccountID(c),
	})
	if err != nil {
		return
	}
	c.JSON(200, subscribers)
}
