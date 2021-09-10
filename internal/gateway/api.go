package gateway

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/disism/hvxahv/internal/gateway/v1alpha1"
	"github.com/gin-gonic/gin"
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

	// The v1alpha1 version of the API service used in the application
	// is usually allowed to be accessed through Token authentication.
	v1 := api.Group("/api/v1")

	// USE AUTH MIDDLEWARE.
	v1.Use(middleware.Auth)

	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	// HTTP API for public query of ActivityPub.
	// ActivityPub WebFinger https://github.com/w3c/activitypub/issues/194 .
	api.GET("/.well-known/webfinger", handlers.WebFingerHandler)

	// Get the actors in the activityPub protocol.
	// https://www.w3.org/TR/activitypub/#actor-objects
	api.GET("/u/:actor", handlers.GetActorHandler)

	// Inbox
	// https://www.w3.org/TR/activitypub/#inbox
	api.POST("/u/:actor/inbox", handlers.InboxHandler)

	// The internal open API service provided by hvxahv usually does not require Token authentication,
	// as login and registration.
	api.POST("/accounts/new", handlers.NewAccountsHandler)
	api.POST("/accounts/login", handlers.LoginHandler)

	// INTERNAL API GROUP.
	v1alpha1.V1Accounts(v1)

	v1alpha1.V1Channels(v1)

	v1alpha1.V1Accounts(v1)

	v1alpha1.V1Articles(v1)

	return api
}
