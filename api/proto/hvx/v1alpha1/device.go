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

func Devices(v1 *gin.RouterGroup) {
	v1.GET("/devices", handler.GetDevices)
	v1.DELETE("/device/:id", handler.DeleteDevice)
}
