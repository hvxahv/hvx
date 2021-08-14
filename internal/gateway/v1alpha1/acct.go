package v1alpha1

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func V1Accounts(r *gin.Engine) {

	// The internal open API service provided by hvxahv usually does not require Token authentication,
	// as login and registration.
	r.POST("/accounts/new", handlers.NewAccountsHandler)
	r.POST("/accounts/login", handlers.LoginHandler)

	// The v1alpha1 version of the API service used in the application
	// is usually allowed to be accessed through Token authentication.
	v1 := r.Group("/api/v1")
	// Load verification token middleware.
	v1.Use(middleware.Auth)
	{
		v1.GET("/accounts/i", handlers.GetAccountsHandler)
		v1.POST("/accounts/follow", handlers.FollowHandler)
		//v1alpha1.POST("/upload/avatar", handlers.UploadAvatar)

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


		r.GET("/u/:actor/outbox", handlers.GetActorOutbox)
		//r.POST("/u/:user/inbox", handlers.InboxHandler)
	}
}