package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
)

func V1Channels(v1 *gin.RouterGroup) {

	v1.GET("/channels/managed", handlers.GetManagedChannelsHandler)

	v1.POST("/channel/create", handlers.CreateChannelHandler)
	v1.POST("/channel/delete", handlers.DeleteChannelHandler)

	v1.POST("/channels/update", handlers.UpdateChannelHandler)

	v1.POST("/channels/admin/new", handlers.NewChannelAdminHandler)
	v1.POST("/channels/admin/remove", handlers.RemoveChannelAdminHandler)

	v1.POST("/channels/subscriber/new", handlers.CreateSubscriberHandler)
	v1.GET("/channels/subscribers", handlers.GetSubscribersHandler)

	v1.GET("/channels/broadcasts", handlers.GetBroadcastsHandler)
	v1.POST("/channels/broadcasts/create", handlers.CreateBroadcastHandler)

}
