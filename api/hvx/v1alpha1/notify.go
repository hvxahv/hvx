package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Notify(v1 *gin.RouterGroup) {
	v1.POST("/notify/sub", handler.NotifySubHandler)
}
