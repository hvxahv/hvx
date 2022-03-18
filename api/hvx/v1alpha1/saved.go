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

func V1Saved(v1 *gin.RouterGroup) {
	v1.GET("/saves", handler.GetSaves)

	v1.GET("/saved/:id")

	v1.POST("/saved")
}
