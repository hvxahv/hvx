package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Activity(v1 *gin.RouterGroup) {
	v1.POST("/activity/outbox", handler.OutboxHandler)
}
