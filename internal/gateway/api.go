package gateway

import (
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

	// Open API routing for the ActivityPub protocol.
	// ActivityPub https://www.w3.org/TR/activitypub/
	v1alpha1.V1ActivityPub(api)

	v1alpha1.V1Accounts(api)

	v1alpha1.V1Chan(api)

	v1alpha1.V1Messages(api)


	return api
}
