package powers

import (
	"github.com/gin-gonic/gin"
	"hvxahv/internal/powers/handlers"
)

func v1(r *gin.Engine) {

	// Simple group: v1
	v1 := r.Group("/v1")
	{

		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status": "ok",
			})
		})

		v1.POST("/upload/avatar", handlers.UploadAvatar)

		//v1.GET("/u/:user", v1alpha1.GetActorHandler)
		//v1.GET("/u/:user/outbox", v1alpha1.GetActorOutbox)
		//v1.POST("/u/:user/inbox", activity.InboxHandler)
		//
		//v1.GET("/u/:user/article/:id", activity.GetPublicArticleHandler)
		////v1 Http api interface for testing
		//v1.POST("/accept", test.AcceptHandler)
		//
		//v1.GET("/u/:user/following", accounts.FollowingResponse)
		//v1.GET("/u/:user/followers", accounts.FollowersResponse)

		// Default account login and registration system


		//v1.POST("/account/login", v1alpha1.VerificationHandler)
		//
	}

}
