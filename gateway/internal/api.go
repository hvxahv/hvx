/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/gateway/v1alpha1"
)

func APIServer() *gin.Engine {
	api := gin.Default()
	api.Use(CORS())

	api.GET("health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	// Public API.
	api.GET("/public/*x", PublicHandler)
	api.POST("/public/*x", PublicHandler)
	api.GET("/.well-known/webfinger", WellKnownHandler)
	api.GET("/u/:actor", GetActorHandler)

	v1 := api.Group("/api/v1")
	v1.Use(Auth)
	v1.GET("/search/:actor", v1alpha1.SearchActorsHandler)
	v1.DELETE("/account", v1alpha1.DeleteAccountHandler)
	return api
}
