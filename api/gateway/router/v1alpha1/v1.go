package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/gateway/middleware"
)

func V1Group(r *gin.Engine) {
	// Functions that can be accessed through Token, carry token when requesting
	v1 := r.Group("/api/v1alpha1")
	v1.Use(middleware.JWTAuth)
	//{
	//	/* Accounts Services */
	//	v1.GET("/account/i", handler.GetAccountsHandler)
	//	v1.POST("/account/delete", handler.DeleteAccountHandler)
	//	v1.POST("/account/settings", handler.AccountSettingHandler)
	//	// Logged in users get INBOX information
	//	v1.GET("/inbox", activity.GetInboxHandler)
	//
	//	// Follow
	//	v1.POST("/follow", activity.FollowHandler)
	//	v1.POST("/follower/accept", activity.FollowerAcceptHandler)
	//	v1.GET("/follower", follow.GetFollowerHandler)
	//	v1.GET("/following", follow.GetFollowingHandler)
	//
	//	/*  Article Services */
	//	v1.GET("/articles", activity.GetArticles)
	//	v1.POST("/article/new", activity.NewArticleHandler)
	//	//v1alpha1.POST("/activity/update", handler.UpdateArticleHandler)
	//	//v1alpha1.POST("/activity/delete", handler.DeleteArticleHandler)

	//}

}
