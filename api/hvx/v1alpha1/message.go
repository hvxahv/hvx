package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/hvx/handler"
)

func V1Messages(v1 *gin.RouterGroup) {
	v1.GET("/message/access", handler.GetMessageAccessHandler)
	//v1.POST("/message/access", handler.NewMessagesAccessHandler)
}
