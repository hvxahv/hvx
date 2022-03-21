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
}
