package response

import "github.com/gin-gonic/gin"

// CreateArticleResponse 创建文章时返回的 http response
func CreateArticleResponse(c *gin.Context, r string) {
	if r == "error" {
		c.JSON(500, gin.H{
			"state": "500",
			"message": "创建文章失败",
		})
	} else {
		c.JSON(200, gin.H{
			"state": "200",
			"message": "创建文章成功",
		})
	}
}
