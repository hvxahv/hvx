package activity

import "github.com/gin-gonic/gin"

// CreateArticleResponse 创建文章时返回的 server response
func SendActivityResponse(c *gin.Context, r int32) {
	if r == 202 {
		c.JSON(int(r), gin.H{
			"status": "202",
			"message": "发送到远程服务器成功",
		})
	} else {
		c.JSON(200, gin.H{
			"state": "200",
			"message": "发送到远程服务器失败",
		})
	}
}
