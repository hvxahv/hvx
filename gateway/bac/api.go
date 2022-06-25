/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package gateway_bac

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/internal/gateway/middleware"
)

func APIServer() *gin.Engine {
	api := gin.Default()
	api.Use(middleware.CORS())

	// The account open API service provided by hvxahv usually does not require Token authentication.
	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	//Get the total number of users of the current instance.
	//api.GET("/public/account/count", public.GetPublicAccountCountHandler)

	// Get the instance details of the current instance.
	//api.GET("/instance", public.GetPublicInstanceDetailsHandler)

	//// Open API routing for the ActivityPub protocol.
	//// ActivityPub https://www.w3.org/TR/activitypub/
	//// HTTP API for public query of ActivityPub.
	//// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	//api.GET("/.well-known/webfinger", handler.GetWebFingerHandler)
	//
	////api.GET("/.well-known/nodeinfo", public.GetNodeInfoHandler)
	//
	//// Get the actors in the activityPub protocol.
	//// https://www.w3.org/TR/activitypub/#actor-objects
	//api.GET("/u/:actor", handler.GetActorHandler)
	//
	////// The type of Channel is a service in Activitypub. Details:
	////// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-service
	////api.GET("/c/:actor", handler.GetChannelHandler)
	////api.POST("/c/:actor/inbox", handler.ChannelInboxHandler)
	//
	//// Inbox
	//// https://www.w3.org/TR/activitypub/#inbox
	//api.POST("/u/:actor/inbox", handler.InboxHandler)
	//
	//// Account creation and authorization API.
	//api.POST("/account/create", handler.CreateAccountHandler)
	//api.POST("/account/auth", handler.AuthAccountHandler)
	//
	//// The v1alpha version of the API service used in the application
	//// is usually allowed to be accessed through Token authentication.
	//v1 := api.Group("/api/v1")
	//
	//// USE AUTH MIDDLEWARE.
	//v1.Use(middleware.Auth)
	//
	//// INTERNAL API GROUP.
	//// The account API requires TOKEN authentication to access.
	//v1alpha.V1Accounts(v1)
	//
	//v1alpha.Devices(v1)
	//
	//v1alpha.V1Articles(v1)
	//
	//v1alpha.V1Activity(v1)
	//
	//v1alpha.V1Channel(v1)
	//
	//v1alpha.V1Saved(v1)
	//
	//v1alpha.V1Notify(v1)
	//
	//v1alpha.V1Messages(v1)

	return api
}
