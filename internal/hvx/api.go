package hvx

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/hvx/v1alpha1"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
	"github.com/hvxahv/hvxahv/internal/hvx/middleware"
	"github.com/hvxahv/hvxahv/internal/hvx/public"
)

func APIServer() *gin.Engine {
	api := gin.Default()
	api.Use(middleware.CORS())

	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	api.GET("/public/account/count", public.GetPublicAccountCountHandler)

	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	api.GET("/.well-known/webfinger", handler.WebFingerHandler)

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

	// The internal open API service provided by hvxahv usually does not require Token authentication,
	// as login and registration.
	api.POST("/signup", handler.CreateAccountHandler)
	api.POST("/signin", handler.SignInHandler)
	//
	//// The v1alpha1 version of the API service used in the application
	//// is usually allowed to be accessed through Token authentication.
	v1 := api.Group("/api/v1")

	// USE AUTH MIDDLEWARE.
	v1.Use(middleware.Auth)

	// INTERNAL API GROUP.
	v1alpha1.V1Accounts(v1)
	//
	//v1alpha1.V1Channels(v1)
	//
	//v1alpha1.V1Articles(v1)
	//
	//v1alpha1.V1Follow(v1)
	//
	//v1alpha1.V1Saved(v1)
	//
	//v1alpha1.V1Storage(v1)
	//
	//v1alpha1.V1Notify(v1)
	//
	//v1alpha1.V1Messages(v1)

	return api
}
