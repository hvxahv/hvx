package router

import (
	"github.com/gin-gonic/gin"
	middleware2 "hvxahv/pkg/middleware"
)

// Router Used to provide routing for RESTful access,
// set up middleware to solve cross-domain (CORS),
// set up JWTAuth middleware,
// Implementation of routing using gin web framework.
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware2.CORS())

	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Default account login and registration system
	//r.POST("/account/new", v1alpha1.NewAccountsHandler)
	//r.POST("/account/login", v1alpha1.VerificationHandler)
	//
	//// API routing for Activitypub function
	//r.GET("/.well-known/webfinger", v1alpha1.GetWebFingerHandler)
	//r.GET("/u/:user", v1alpha1.GetActorHandler)
	//r.GET("/u/:user/outbox", v1alpha1.GetActorOutbox)
	//r.POST("/u/:user/inbox", activity.InboxHandler)
	//
	//r.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)
	//// Http api interface for testing
	//r.POST("/accept", test.AcceptHandler)
	//
	//r.GET("/u/:user/following", accounts.FollowingResponse)
	//r.GET("/u/:user/followers", accounts.FollowersResponse)

	V1Group(r)

	return r
}
