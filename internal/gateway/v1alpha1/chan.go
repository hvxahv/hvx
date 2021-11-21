package v1alpha1

import (
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func V1Channels(v1 *gin.RouterGroup) {

	v1.POST("/channel/new", handlers.NewChannelHandler)
	v1.POST("/channel/update", handlers.UpdateChannelHandler)
	v1.POST("/channel/delete", handlers.DeleteChannelHandler)

	v1.POST("/channel/admin/new", handlers.NewChannelAdminHandler)
	v1.POST("/channel/admin/remove", handlers.RemoveChannelAdminHandler)

	v1.POST("/channel/subscriber/new", handlers.NewSubscriberHandler)
	v1.POST("/channel/subscriber/list", handlers.GetSubscriberListHandler)

}