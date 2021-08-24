package v1alpha1

import (
	"github.com/disism/hvxahv/internal/gateway/handlers"
	"github.com/disism/hvxahv/internal/gateway/middleware"
	"github.com/gin-gonic/gin"
)

func V1Messages(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	v1.Use(middleware.Auth)
	{
		v1.POST("/messages/outbox", handlers.OutboxHandler)
	}
}