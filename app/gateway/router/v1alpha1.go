package router

import (
	"github.com/gin-gonic/gin"
	middleware2 "hvxahv/pkg/middleware"
)

func V1Group(r *gin.Engine) {
	// Functions that can be accessed through Token, carry token when requesting
	v1 := r.Group("/api/v1alpha1")
	v1.Use(middleware2.JWTAuth)
	//{
	//	/* Accounts Services */
	//	v1.GET("/account/i", v1alpha1.GetAccountsHandler)
	//	v1.POST("/account/delete", v1alpha1.DeleteAccountHandler)
	//	v1.POST("/account/settings", v1alpha1.AccountSettingHandler)
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
	//	//v1alpha1.POST("/activity/update", v1alpha1.UpdateArticleHandler)
	//	//v1alpha1.POST("/activity/delete", v1alpha1.DeleteArticleHandler)

	//}

}
