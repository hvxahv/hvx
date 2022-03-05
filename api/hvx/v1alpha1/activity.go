/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package v1alpha1

import "github.com/gin-gonic/gin"

func V1Activity(v1 *gin.RouterGroup) {
	v1.POST("/follow/requests")

	v1.POST("/follow/accept")

	v1.GET("/follower")

	v1.GET("/following")

	v1.POST("/follow/:id/authorize")

	// OUTBOX
	v1.POST("/activity/outbox")
}
