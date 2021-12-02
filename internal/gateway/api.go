package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
	"github.com/hvxahv/hvxahv/internal/gateway/middleware"
	"github.com/hvxahv/hvxahv/internal/gateway/v1alpha1"
)

// APIServer Used to provide routing for RESTful access,
// set up middleware to solve cross-domain (CORS),
// set up JWTAuth middleware,
// Implementation of routing using gin web framework.
func APIServer() *gin.Engine {
	api := gin.Default()
	api.Use(middleware.CORS())

	api.GET("ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong!",
		})
	})

	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	api.GET("/.well-known/webfinger", handlers.WebFingerHandler)

	// Get the actors in the activityPub protocol.
	// https://www.w3.org/TR/activitypub/#actor-objects
	api.GET("/u/:actor", handlers.GetActorHandler)

	// The type of Channel is a service in Activitypub. Details:
	// https://www.w3.org/TR/activitystreams-vocabulary/#dfn-service
	api.GET("/c/:actor", handlers.GetChannelHandler)
	api.POST("/c/:actor/inbox", handlers.ChannelInboxHandler)

	// Inbox
	// https://www.w3.org/TR/activitypub/#inbox
	api.POST("/u/:actor/inbox", handlers.InboxHandler)

	// The internal open API service provided by hvxahv usually does not require Token authentication,
	// as login and registration.
	api.POST("/accounts/signup", handlers.CreateAccountsHandler)
	api.POST("/accounts/sign", handlers.SignInHandler)

	// The v1alpha1 version of the API service used in the application
	// is usually allowed to be accessed through Token authentication.
	v1 := api.Group("/api/v1")

	// USE AUTH MIDDLEWARE.
	v1.Use(middleware.Auth)

	// INTERNAL API GROUP.
	v1alpha1.V1Accounts(v1)

	v1alpha1.V1Channels(v1)

	v1alpha1.V1Articles(v1)

	// Follow API
	v1alpha1.V1Follow(v1)

	v1alpha1.V1Saved(v1)

	return api
}
