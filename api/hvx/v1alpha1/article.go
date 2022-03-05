/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package v1alpha1

import "github.com/gin-gonic/gin"

func V1Articles(v1 *gin.RouterGroup) {
	v1.POST("/article/create")

	v1.GET("/article/:id")
}
