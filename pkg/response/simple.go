package response

import "github.com/gin-gonic/gin"

// SimpleResponse 简单的 https 相应, 接收状态码和一条信息
func SimpleResponse(c *gin.Context, code, m string) {
	c.JSON(401, gin.H{
		"state": code,
		"message": m,
	})
}