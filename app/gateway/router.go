package main

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/client/account"
	"hvxahv/app/gateway/handler"
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
	r.POST("/u/:user/inbox", handler.GetActorInbox)

	r.GET("/u/:user/following", account.FollowersResponse)
	r.GET("/u/:user/followers", account.FollowersResponse)

	// 通过 Token 才能访问的功能
	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWTAuth)
	{
		/* Accounts Services */
		v1.GET("/account/i", handler.GetAccountsHandler)
		v1.POST("/account/delete", handler.DeleteAccountHandler)
		v1.POST("/account/settings", handler.AccountSettingHandler)

		/*  Article Services */
		v1.POST("/article", handler.GetArticlesHandler)
		v1.POST("/article/new", handler.CreateArticleHandler)
		v1.POST("/article/update", handler.UpdateArticleHandler)
		v1.POST("/article/delete", handler.DeleteArticleHandler)

	}
	return r
}