/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Channel(v1 *gin.RouterGroup) {
	// CHANNEL.
	v1.POST("/channel", handler.CreateChannelHandler)
	v1.GET("/channels", handler.GetChannelsHandler)
	v1.DELETE("/channel/:id", handler.DeleteChannelHandler)

	// CHANNEL ADMIN.
	v1.POST("/channel/admin", handler.AddAdminToChannelHandler)
	v1.DELETE("/channel/admin", handler.RemoveAdminFromChannelHandler)
	v1.GET("/channel/admins/:id", handler.GetAdminsOfChannelHandler)

	// CHANNEL SUBSCRIBER.
	v1.POST("/channel/subscriber", handler.AddSubscriberToChannelHandler)
	v1.GET("/channel/unsubscribe/:id", handler.UnsubscribeFromChannelHandler)

	// CHANNEL SUBSCRIBER FOR ADMIN.
	v1.DELETE("/channel/subscriber/:id", handler.RemoveSubscriberFromChannelHandler)
	v1.GET("/channel/subscribers/:id", handler.GetSubscribersOfChannelHandler)

	// CHANNEL BROADCAST.
	v1.POST("/channel/broadcast", handler.CreateBroadcastHandler)
	v1.GET("/channel/broadcasts/:id", handler.GetBroadcastsOfChannelHandler)
	v1.DELETE("/channel/broadcast/:id", handler.DeleteBroadcastHandler)
}
