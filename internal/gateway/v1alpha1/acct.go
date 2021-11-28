package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
)

func V1Accounts(v1 *gin.RouterGroup) {

		//v1.GET("/accounts/iam", handlers.GetAccountsHandler)
		v1.GET("/inbox", handlers.GetInboxesHandler)
		v1.GET("/iam/timelines")

		// Delete accounts
		//v1.POST("/accounts/delete", handlers.DeleteAccount)
		//v1.POST("/accounts/update", handlers.UpdateAccount)

		v1.POST("/upload/avatar", handlers.UploadAvatar)

		//v1alpha1.GET("/u/:user", v1alpha1.GetActorHandler)
		//v1alpha1.GET("/u/:user/outbox", v1alpha1.GetActorOutbox)
		//
		//v1alpha1.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)



}
