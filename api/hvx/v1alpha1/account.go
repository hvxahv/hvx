package v1alpha1

import "github.com/gin-gonic/gin"

func V1Accounts(v1 *gin.RouterGroup) {
	/**
	 * Account for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/account).
	 */
	v1.GET("/account/iam")
	//v1.GET("/iam/timelines")

	// Delete account
	v1.DELETE("/account")

	v1.PATCH("/account/username")
	v1.PATCH("/account/password")
	v1.PATCH("/account/mail")

	v1.GET("/account/rsa/public")

	v1.GET("/account/logout")

	/**
	 * Actor for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/actor).
	 */
	v1.PATCH("/actor/edit")

	/**
	 * Devices for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/device).
	 */
	v1.GET("/devices")
	v1.POST("/devices/delete")

	/**
	 * ECDH for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/dh).
	 */
	//	https://github.com/hvxahv/hvxahv/blob/main/SECURITY.md
	v1.GET("/dh/private")
	v1.POST("/dh")
	v1.GET("/dh/wait")
}
