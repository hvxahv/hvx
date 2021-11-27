package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
)

func V1Follow(v1 *gin.RouterGroup) {
	v1.POST("/follow/requests", handlers.FollowReqHandler)

	v1.POST("/follow/accept", handlers.FollowAcceptHandler)

	v1.GET("/follower", handlers.FollowerHandler)

	v1.GET("/following", handlers.FollowingHandler)

	// Accept Follow.
	v1.POST("/follow/:id/authorize")

}

