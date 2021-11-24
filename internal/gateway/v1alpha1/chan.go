package v1alpha1

import (
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func V1Channels(v1 *gin.RouterGroup) {

	v1.POST("/channels/new", handlers.NewChannelHandler)
	v1.POST("/channels/update", handlers.UpdateChannelHandler)
	v1.POST("/channels/delete", handlers.DeleteChannelHandler)

	v1.POST("/channels/admin/new", handlers.NewChannelAdminHandler)
	v1.POST("/channels/admin/remove", handlers.RemoveChannelAdminHandler)

	v1.POST("/channels/subscriber/new", handlers.NewSubscriberHandler)
	v1.POST("/channels/subscriber/list", handlers.GetSubscriberListHandler)

}