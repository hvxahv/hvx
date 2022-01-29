package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/api/hvx/handler"
)

func V1Saved(v1 *gin.RouterGroup) {
	v1.GET("/saves", handler.GetSavesHandler)
	v1.GET("/saved/:id", handler.GetSavedByIDHandler)
	v1.POST("/saved", handler.SavedHandler)
}
