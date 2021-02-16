package main

import (
	"github.com/gin-gonic/gin"
	"hvxahv/app/gateway/handler"
	"hvxahv/app/gateway/handler/activity"
	"hvxahv/app/gateway/handler/follow"
	"hvxahv/app/test"
	"hvxahv/pkg/client/accounts"
	"hvxahv/pkg/middleware"
)

func IngressRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	/* 账号登录和注册 */
	r.POST("/account/new", handler.NewAccountsHandler)
	r.POST("/account/login", handler.VerificationHandler)

	// Activitypub 功能
	r.GET("/.well-known/webfinger", handler.GetWebFingerHandler)
	r.GET("/u/:user", handler.GetActorHandler)
	r.GET("/u/:user/outbox", handler.GetActorOutbox)
	r.POST("/u/:user/inbox", activity.InboxHandler)

	r.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)
	// 用于 测试的
	r.POST("/accept", test.AcceptHandler)

	r.GET("/u/:user/following", accounts.FollowingResponse)
	r.GET("/u/:user/followers", accounts.FollowersResponse)



	// 通过 Token 才能访问的功能
	v1 := r.Group("/api/v1alpha1")
	v1.Use(middleware.JWTAuth)
	{
		/* Accounts Services */
		v1.GET("/account/i", handler.GetAccountsHandler)
		v1.POST("/account/delete", handler.DeleteAccountHandler)
		v1.POST("/account/settings", handler.AccountSettingHandler)
		// 已经登录用户获取 INBOX 信息
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