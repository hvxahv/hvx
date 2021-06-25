package powers

import (
	"github.com/gin-gonic/gin"
	"hvxahv/internal/powers/handlers"
	"hvxahv/internal/powers/middleware"
)

// Router Used to provide routing for RESTful access,
// set up middleware to solve cross-domain (CORS),
// set up JWTAuth middleware,
// Implementation of routing using gin web framework.
func Router() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	r.GET("ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Account function of hvxahv
	r.POST("/accounts/new", handlers.NewAccountsHandler)
	r.POST("/accounts/login", handlers.LoginHandler)
	r.POST("/upload/avatar", handlers.Avatar)


	activityPubV1(r)
	v1(r)

	return r
}

