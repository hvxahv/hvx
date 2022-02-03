package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Saved(v1 *gin.RouterGroup) {
	v1.POST("/saved", handler.SavedHandler)
	v1.GET("/saved/:id", handler.GetSavedHandler)
	v1.GET("/saves", handler.GetSavesHandler)
}
