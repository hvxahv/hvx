/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/internal/gateway/middleware"
	"github.com/hvxahv/hvx/internal/gateway/public"
	"github.com/hvxahv/hvx/internal/gateway/v1alpha1"
)

func APIServer() *gin.Engine {
	api := gin.Default()
	api.Use(middleware.CORS())

	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	// Public API.
	api.GET("/public/*x", public.Handler)
	api.POST("/public/*x", public.Handler)
	api.GET("/.well-known/webfinger", public.WellKnownHandler)
	api.GET("/u/:actor", public.GetActorHandler)

	v1 := api.Group("/api/v1")
	v1.Use(middleware.Auth)
	v1.GET("/search/:actor", v1alpha1.SearchActorsHandler)
	v1.DELETE("/account", v1alpha1.DeleteAccountHandler)
	return api
}
