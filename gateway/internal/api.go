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

	// Public API.
	api.GET("/public/*x", PublicHandler)
	api.POST("/public/*x", PublicHandler)
	// AUTH
	api.POST("auth", AuthHandler)
	// ACTIVITYPUB
	api.GET("/.well-known/webfinger", WellKnownHandler)
	api.GET("/u/:actor", GetActorHandler)
	api.POST("/u/:actor/inbox", InboxHandler)

	v1 := api.Group("/api/v1")
	v1.Use(Auth)
	// ACCOUNT SERVICES
	v1.PATCH("/account/*x", v1alpha1.AccountHandler)

	// ACTOR SERVICES
	v1.GET("/search/:actor", v1alpha1.SearchActorsHandler)
	v1.PUT("/actor", v1alpha1.ActorHandler)

	// DEVICES SERVICES
	v1.DELETE("/device", v1alpha1.DeviceHandler)
	v1.GET("/device/*x", v1alpha1.DeviceHandler)

	// CHANNEL SERVICES
	v1.POST("/channel", v1alpha1.ChannelHandler)
	v1.DELETE("/channel", v1alpha1.ChannelHandler)

	v1.GET("/channel/*x", v1alpha1.ChannelHandler)
	v1.POST("/channel/*x", v1alpha1.ChannelHandler)
	v1.DELETE("/channel/*x", v1alpha1.ChannelHandler)

	// ARTICLE SERVICES
	v1.POST("/article", v1alpha1.ArticleHandler)
	v1.PUT("/article", v1alpha1.ArticleHandler)
	v1.DELETE("/article", v1alpha1.ArticleHandler)
	v1.GET("/article/*x", v1alpha1.ArticleHandler)
	v1.DELETE("/article/*x", v1alpha1.ArticleHandler)

	// SAVED SERVICES
	v1.POST("/saved", v1alpha1.SavedHandler)
	v1.PUT("/saved", v1alpha1.SavedHandler)
	v1.DELETE("/saved", v1alpha1.SavedHandler)
	v1.GET("/saved/*x", v1alpha1.SavedHandler)
	v1.DELETE("/saved/*x", v1alpha1.SavedHandler)

	// ACTIVITY INBOX SERVICES
	v1.GET("/activity/*x", v1alpha1.ActivityHandler)
	v1.DELETE("/activity/*x", v1alpha1.ActivityHandler)

	// FS
	// The avatar is uploaded through the http interface provided by Gin,
	// and the uploaded link is then sent to the actor server via gRPC to change or add the actor avatar.
	v1.POST("/fs/avatar", AvatarHandler)
	v1.POST("/fs/attach", AttachHandler)
	return api
}
