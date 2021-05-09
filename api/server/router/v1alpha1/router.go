package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/server/middleware"
	"hvxahv/app/gateway/handler"
	"hvxahv/app/gateway/handler/activity"
	"hvxahv/app/gateway/handler/follow"
	"hvxahv/app/test"
	"hvxahv/internal/client/accounts"
)

// Router Used to provide routing for http access,
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
	r.POST("/account/new", handler.NewAccountsHandler)
	r.POST("/account/login", handler.VerificationHandler)

	// API routing for Activitypub function
	r.GET("/.well-known/webfinger", handler.GetWebFingerHandler)
	r.GET("/u/:user", handler.GetActorHandler)
	r.GET("/u/:user/outbox", handler.GetActorOutbox)
	r.POST("/u/:user/inbox", activity.InboxHandler)

	r.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)
	// Http api interface for testing
	r.POST("/accept", test.AcceptHandler)

	r.GET("/u/:user/following", accounts.FollowingResponse)
	r.GET("/u/:user/followers", accounts.FollowersResponse)



	// Functions that can be accessed through Token, carry token when requesting
	v1 := r.Group("/api/v1alpha1")
	v1.Use(middleware.JWTAuth)
	{
		/* Accounts Services */
		v1.GET("/account/i", handler.GetAccountsHandler)
		v1.POST("/account/delete", handler.DeleteAccountHandler)
		v1.POST("/account/settings", handler.AccountSettingHandler)
		// Logged in users get INBOX information
		v1.GET("/inbox", activity.GetInboxHandler)

		// Follow
		v1.POST("/follow", activity.FollowHandler)
		v1.POST("/follower/accept", activity.FollowerAcceptHandler)
		v1.GET("/follower", follow.GetFollowerHandler)
		v1.GET("/following", follow.GetFollowingHandler)

		/*  Article Services */
		v1.GET("/articles", activity.GetArticles)
		v1.POST("/article/new", activity.NewArticleHandler)
		//v1alpha1.POST("/activity/update", handler.UpdateArticleHandler)
		//v1alpha1.POST("/activity/delete", handler.DeleteArticleHandler)

	}
	return r
}