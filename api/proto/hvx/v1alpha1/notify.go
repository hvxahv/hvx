/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package v1alpha1

import "github.com/gin-gonic/gin"

func V1Notify(v1 *gin.RouterGroup) {
	v1.POST("/notify/subscription")
}
