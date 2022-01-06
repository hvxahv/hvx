package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Articles(v1 *gin.RouterGroup) {
	// New status & article
	v1.POST("/article/new", handler.NewArticleHandler)

	// Get article by article_id
	v1.GET("/article/:id", handler.GetArticleHandler)
}
