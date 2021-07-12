package gateway

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/gin-gonic/gin"
)

// Router Used to provide routing for RESTful access,
// set up middleware to solve cross-domain (CORS),
// set up JWTAuth middleware,
// Implementation of routing using gin web framework.
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// The internal open API service provided by hvxahv usually does not require Token authentication,
	// as login and registration.
	r.POST("/accounts/new", handlers.NewAccountsHandler)
	r.POST("/accounts/login", handlers.LoginHandler)


	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	activityPubV1(r)

	// The v1 version of the API service used in the application
	// is usually allowed to be accessed through Token authentication.
	v1(r)

	return r
}

