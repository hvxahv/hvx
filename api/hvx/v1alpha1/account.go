package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/hvx/handler"
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

	v1.GET("/account/logout", handler.LogoutHandler)

	/**
	 * Actor for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/actors).
	 */
	v1.PATCH("/actor/edit", handler.EditActorHandler)

	/**
	 * ECDH for v1 version of the server api endpoints for the account resource type (hvx.hvxahv.com/v1/account).
	 */
	//	https://github.com/hvxahv/hvxahv/blob/main/SECURITY.md
	v1.GET("/account/dh/private")
	v1.POST("/account/dh")
	v1.GET("/account/dh/wait")
}
