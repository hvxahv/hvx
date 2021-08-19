package v1alpha1

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func V1Chan(r *gin.Engine) {

	v1 := r.Group("/api/v1")
	// Load verification token middleware.
	v1.Use(middleware.Auth)
	{
		v1.POST("/channel/new", handlers.NewChannelHandler)
		v1.POST("/channel/admin/new", handlers.NewChannelAdminHandler)
	}
}