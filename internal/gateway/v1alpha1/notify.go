package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
)

func V1Notify(v1 *gin.RouterGroup) {
	v1.POST("/notify/sub", handlers.NotifySubHandler)
}
