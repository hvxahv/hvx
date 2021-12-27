package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
)

func V1Saved(v1 *gin.RouterGroup) {
	v1.GET("/saved")
	v1.POST("/saved", handlers.SavedHandler)
}
