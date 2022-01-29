package v1alpha1

import "github.com/gin-gonic/gin"

func V1Storage(v1 *gin.RouterGroup) {
	v1.POST("/storage/avatar")
	v1.POST("/storage/media")
}
