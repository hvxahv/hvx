/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package hvx

import (
	"github.com/gin-gonic/gin"
	v1alpha12 "github.com/hvxahv/hvxahv/api/proto/hvx/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
	"github.com/hvxahv/hvxahv/pkg/identity/middleware"
)

func APIServer() *gin.Engine {
	api := gin.Default()
	api.Use(middleware.CORS())

	// The internal open API service provided by hvxahv usually does not require Token authentication.
	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	// Get the total number of users of the current instance.
	//api.GET("/public/account/count", public.GetPublicAccountCountHandler)

	// Get the instance details of the current instance.
	//api.GET("/instance", public.GetPublicInstanceDetailsHandler)

	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	api.GET("/.well-known/webfinger", handler.GetWebFingerHandler)

	//api.GET("/.well-known/nodeinfo", public.GetNodeInfoHandler)

	// Get the actors in the activityPub protocol.
	// https://www.w3.org/TR/activitypub/#actor-objects
	api.GET("/u/:actor", handler.GetActorHandler)

	//// The type of Channel is a service in Activitypub. Details:
	//// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-service
	//api.GET("/c/:actor", handler.GetChannelHandler)
	//api.POST("/c/:actor/inbox", handler.ChannelInboxHandler)

	// Inbox
	// https://www.w3.org/TR/activitypub/#inbox
	api.POST("/u/:actor/inbox", handler.InboxHandler)

	// Account creation and authorization API.
	api.POST("/account/create", handler.CreateAccountHandler)
	api.POST("/account/auth", handler.AuthAccountHandler)

	// The v1alpha1 version of the API service used in the application
	// is usually allowed to be accessed through Token authentication.
	v1 := api.Group("/api/v1")

	// USE AUTH MIDDLEWARE.
	v1.Use(middleware.Auth)

	// INTERNAL API GROUP.
	// The internal API requires TOKEN authentication to access.
	v1alpha12.V1Accounts(v1)

	v1alpha12.Devices(v1)

	v1alpha12.V1Articles(v1)

	v1alpha12.V1Activity(v1)

	v1alpha12.V1Channel(v1)

	v1alpha12.V1Saved(v1)

	v1alpha12.V1Notify(v1)

	v1alpha12.V1Messages(v1)

	return api
}
