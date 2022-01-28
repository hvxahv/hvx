package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Devices(v1 *gin.RouterGroup) {
	/**
	 * Devices for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/devices).
	 */
	v1.GET("/devices", handler.GetDevicesHandler)
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
