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
		v1.POST("/channel/update", handlers.UpdateChannelHandler)
		v1.POST("/channel/delete", handlers.DeleteChannelHandler)

		v1.POST("/channel/admin/new", handlers.NewChannelAdminHandler)
		v1.POST("/channel/admin/remove", handlers.RemoveChannelAdminHandler)

		v1.POST("/channel/subscriber/new", handlers.NewSubscriberHandler)
		v1.POST("/channel/subscriber/list", handlers.GetSubscriberListHandler)
	}
}