package v1alpha1

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/gin-gonic/gin"
)

func V1Articles(v1 *gin.RouterGroup) {
	v1.POST("/article/new", handlers.NewArticleHandler)
}

