package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
)

func V1Messages(v1 *gin.RouterGroup) {
	v1.GET("/messages/access", handlers.GetMessageAccessHandler)
	v1.POST("/messages/access", handlers.NewMessagesAccessHandler)

}
