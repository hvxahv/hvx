package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Accounts(v1 *gin.RouterGroup) {
	/**
	 * Account for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/account).
	 */
	v1.GET("/account/iam")
	//v1.GET("/iam/timelines")

	// Delete account
	v1.DELETE("/account", handler.DeleteAccount)

	v1.PATCH("/account/username", handler.EditAccountUsernameHandler)
	v1.PATCH("/account/password", handler.EditAccountPasswordHandler)
	v1.PATCH("/account/mail", handler.EditAccountMailHandler)

	v1.GET("/account/rsa/public", handler.GetPublicKeyHandlers)

	/**
	 * Actor for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/actors).
	 */
	v1.PATCH("/actor/edit", handler.EditActorHandler)

	/**
	 * Devices for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/account).
	 */
	//v1.GET("/account/devices", handler.GetDevicesHandler)
	//
	//v1.POST("/account/devices/delete", handler.DeleteDevicesHandler)

	//v1.GET("/account/rsa/public/:id", handler.GetDHPublicJWKHandlers)
	//v1.GET("/account/rsa/private/:id", handler.GetDHPrivateJWKHandlers)
	//v1.POST("/account/rsa/private/request", handler.RequestPrivateKeyHandlers)
	//v1.POST("/account/rsa/private/send", handler.SendPrivateKeyHandlers)
	//
	//// Exit current device.
	//v1.GET("/account/logout", handler.LogoutHandler)

}
