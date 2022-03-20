/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package v1alpha1

import (
	"github.com/gin-gonic/gin"
	"github.com/hvxahv/hvxahv/internal/hvx/handler"
)

func V1Articles(v1 *gin.RouterGroup) {
	v1.POST("/article", handler.CreateArticleHandler)

	v1.PUT("/article", handler.UpdateArticleHandler)

	v1.GET("/article/:id", handler.GetArticleHandler)

	v1.GET("/articles", handler.GetArticlesHandler)

	v1.DELETE("/article/:id", handler.DeleteArticleHandler)
}
