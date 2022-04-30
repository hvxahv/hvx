package v1alpha

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvx/internal/gateway/handler"
)

func V1Accounts(v1 *gin.RouterGroup) {
	v1.GET("/search/:actor", handler.SearchActorsHandler)

	/**
	 * Account for v1 version of the server api endpoints for the account resource type (gateway.hvxahv.com/v1/account).
	 */
	v1.GET("/account/iam")
	//v1.GET("/iam/timelines")

	// Delete account
	v1.DELETE("/account", handler.DeleteAccountHandler)
	v1.GET("/account/logout", handler.LogoutHandler)

	v1.PATCH("/account/username")

	v1.PATCH("/account/password")

	v1.PATCH("/account/mail")

	v1.GET("/account/rsa/public")

	/**
	 * Actor for v1 version of the server api endpoints for the account resource type (gateway.hvxahv.com/v1/actor).
	 */
	v1.PATCH("/actor/edit")

	/**
	 * ECDH for v1 version of the server api endpoints for the account resource type (gateway.hvxahv.com/v1/dh).
	 */
	//	https://github.com/hvxahv/hvx/blob/main/SECURITY.md
	v1.GET("/dh/private")
	v1.POST("/dh")
	v1.GET("/dh/wait")
}

func V1Activity(v1 *gin.RouterGroup) {
	v1.POST("/follow/requests")

	v1.POST("/follow/accept")

	v1.GET("/follower")

	v1.GET("/following")

	v1.POST("/follow/:id/authorize")

	// OUTBOX
	v1.POST("/activity/outbox")
}

func V1Articles(v1 *gin.RouterGroup) {
	v1.POST("/article", handler.CreateArticleHandler)

	v1.PUT("/article", handler.UpdateArticleHandler)

	v1.GET("/article/:id", handler.GetArticleHandler)

	v1.GET("/articles", handler.GetArticlesHandler)

	v1.DELETE("/article/:id", handler.DeleteArticleHandler)
}

func V1Channel(v1 *gin.RouterGroup) {
	// CHANNEL.
	v1.POST("/channel", handler.CreateChannelHandler)
	v1.GET("/channels", handler.GetChannelsHandler)
	v1.DELETE("/channel/:id", handler.DeleteChannelHandler)

	// CHANNEL ADMIN.
	v1.POST("/channel/admin", handler.AddAdminToChannelHandler)
	v1.DELETE("/channel/admin", handler.RemoveAdminFromChannelHandler)
	v1.GET("/channel/admins/:id", handler.GetAdminsOfChannelHandler)

	// CHANNEL SUBSCRIBER.
	v1.POST("/channel/subscriber", handler.AddSubscriberToChannelHandler)
	v1.GET("/channel/unsubscribe/:id", handler.UnsubscribeFromChannelHandler)

	// CHANNEL SUBSCRIBER FOR ADMIN.
	v1.DELETE("/channel/subscriber/:id", handler.RemoveSubscriberFromChannelHandler)
	v1.GET("/channel/subscribers/:id", handler.GetSubscribersOfChannelHandler)

	// CHANNEL BROADCAST.
	v1.POST("/channel/broadcast", handler.CreateBroadcastHandler)
	v1.GET("/channel/broadcasts/:id", handler.GetBroadcastsOfChannelHandler)
	v1.DELETE("/channel/broadcast/:id", handler.DeleteBroadcastHandler)
}

func Devices(v1 *gin.RouterGroup) {
	v1.GET("/devices", handler.GetDevices)
	v1.DELETE("/device/:id", handler.DeleteDevice)
}

func V1Messages(v1 *gin.RouterGroup) {
	v1.GET("/message/access")

	v1.POST("/message/access")
}

func V1Notify(v1 *gin.RouterGroup) {
	v1.POST("/notify/subscription")
}

func V1Saved(v1 *gin.RouterGroup) {
	v1.GET("/saves", handler.GetSaves)

	v1.GET("/saved/:id", handler.GetSaved)

	v1.POST("/saved", handler.CreateSaved)

	v1.PUT("/saved/:id", handler.EditSaved)

	v1.DELETE("/saved/:id", handler.DeleteSaved)
}
