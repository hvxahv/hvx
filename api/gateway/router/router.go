package router

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/gateway/middleware"
	"hvxahv/api/gateway/router/v1alpha1"
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

	// Default account login and registration system
	//r.POST("/account/new", handler.NewAccountsHandler)
	//r.POST("/account/login", handler.VerificationHandler)
	//
	//// API routing for Activitypub function
	//r.GET("/.well-known/webfinger", handler.GetWebFingerHandler)
	//r.GET("/u/:user", handler.GetActorHandler)
	//r.GET("/u/:user/outbox", handler.GetActorOutbox)
	//r.POST("/u/:user/inbox", activity.InboxHandler)
	//
	//r.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)
	//// Http api interface for testing
	//r.POST("/accept", test.AcceptHandler)
	//
	//r.GET("/u/:user/following", accounts.FollowingResponse)
	//r.GET("/u/:user/followers", accounts.FollowersResponse)

	v1alpha1.V1Group(r)

	return r
}
