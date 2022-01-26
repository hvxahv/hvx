package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Accounts(v1 *gin.RouterGroup) {

	v1.GET("/account/iam", handler.GetAccountHandler)
	//v1.GET("/inbox", handler.GetInboxesHandler)
	//v1.GET("/iam/timelines")

	// Delete account
	v1.DELETE("/account", handler.DeleteAccount)
	//v1.POST("/account/update", handler.UpdateAccount)
	//
	//v1.POST("/account/avatar", handler.UploadAvatar)
	//
	////v1alpha1.GET("/u/:user", v1alpha1.GetActorHandler)
	////v1alpha1.GET("/u/:user/outbox", v1alpha1.GetActorOutbox)
	////
	////v1alpha1.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)
	//
	//// Get a list of devices that have been logged in.
	//v1.GET("/account/devices", handler.GetDevicesHandler)
	//
	//// Delete the logged-in device based on the device ID.
	//v1.POST("/account/devices/delete", handler.DeleteDevicesHandler)
	//
	//v1.GET("/account/rsa/public", handler.GetPublicKeyHandlers)
	//v1.GET("/account/rsa/public/:id", handler.GetDHPublicJWKHandlers)
	//v1.GET("/account/rsa/private/:id", handler.GetDHPrivateJWKHandlers)
	//v1.POST("/account/rsa/private/request", handler.RequestPrivateKeyHandlers)
	//v1.POST("/account/rsa/private/send", handler.SendPrivateKeyHandlers)
	//
	//// Exit current device.
	//v1.GET("/account/logout", handler.LogoutHandler)
}
