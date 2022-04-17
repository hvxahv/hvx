/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package v1alpha1

import "github.com/gin-gonic/gin"

func V1Messages(v1 *gin.RouterGroup) {
	v1.GET("/message/access")

	v1.POST("/message/access")
}
