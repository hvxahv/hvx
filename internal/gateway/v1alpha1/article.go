package v1alpha1

import (
	"github.com/hvxahv/hvxahv/internal/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func V1Articles(v1 *gin.RouterGroup) {
	// New status & article
	v1.POST("/article/new", handlers.NewArticleHandler)

	// Get article by article_id
	v1.GET("/article/:id", handlers.GetArticleHandler)
}

