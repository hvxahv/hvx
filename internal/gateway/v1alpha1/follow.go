package v1alpha1

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func V1Follow(v1 *gin.RouterGroup) {
	v1.POST("/follow/requests", handlers.FollowReqHandler)

	// Accept Follow.
	v1.POST("/follow/:id/authorize", handlers.FollowAcceptHandler)
}

