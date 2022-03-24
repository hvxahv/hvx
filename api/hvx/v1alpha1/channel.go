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
	v1.POST("/channel", handler.CreateChannelHandler)
	v1.GET("/channels", handler.GetChannelsHandler)
	v1.DELETE("/channel/:id", handler.DeleteChannelHandler)

	v1.POST("/channel/admin", handler.AddAdminToChannelHandler)
	v1.DELETE("/channel/admin", handler.RemoveAdminFromChannelHandler)
	v1.GET("/channel/admins/:id", handler.GetAdminsOfChannelHandler)
}
