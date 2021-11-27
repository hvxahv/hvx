package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
)

func V1Accounts(v1 *gin.RouterGroup) {

		//v1.GET("/accounts/iam", handlers.FetchAccountsHandler)

		v1.GET("/inbox", handlers.GetInboxesHandler)
		v1.GET("/iam/timelines")

		//v1.POST("/accounts/follow", handlers.FollowReqHandler)

		// Delete accounts
		//v1.POST("/accounts/delete", handlers.DeleteAccount)
		//v1.POST("/accounts/update", handlers.UpdateAccount)

		v1.POST("/upload/avatar", handlers.UploadAvatar)

		//v1alpha1.GET("/u/:user", v1alpha1.GetActorHandler)
		//v1alpha1.GET("/u/:user/outbox", v1alpha1.GetActorOutbox)
		//v1alpha1.POST("/u/:user/inbox", activity.InboxHandler)
		//
		//v1alpha1.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)
		////v1alpha1 Http api interface for testing
		//v1alpha1.POST("/accept", test.AcceptHandler)
		//
		//v1alpha1.GET("/u/:user/following", accounts.FollowingResponse)
		//v1alpha1.GET("/u/:user/followers", accounts.FollowersResponse)

		// Default account login and registration system

		//v1alpha1.POST("/account/login", v1alpha1.VerificationHandler)

		//r.GET("/u/:actor/outbox", handlers.GetActorOutbox)
		//r.POST("/u/:user/inbox", handlers.InboxHandler)


}
