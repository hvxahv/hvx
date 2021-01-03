package app

import (
	"github.com/gin-gonic/gin"
	"log"
)

// DeleteStatusHandler 删除一篇状态
func DeleteStatusHandler(c *gin.Context) {
	id := c.PostForm("id")
	if err := DeleteStatusByID(id); err != nil {
		log.Println("Delete Error!", err)
	}
	c.JSON(200, gin.H{
		"status": "200",
		"message": "DELETE SUCCESS",
	})

}

// ShowStatusListHandler ... 获取所有状态列表
func ShowStatusListHandler(c *gin.Context) {
	name, ok := c.Get("loginUser")
	if !ok {
		log.Println("获取用户名失败")
	}
	author, ok := name.(string)
	if !ok {
		log.Println("用户名转换成字符串失败")
	}

	r, err := ShowStatusLis(author)
	if err != nil {
		log.Println("Query Status Errors", err)
	}
	c.JSON(200, gin.H{
		"state": "200",
		"status": r,
	})
}

// CreateStatusHandler 创建一篇状态
func CreateStatusHandler(c *gin.Context) {
	name, ok := c.Get("loginUser")
	if !ok {
		log.Println("获取用户名失败")
	}
	author, ok := name.(string)
	if !ok {
		log.Println("用户名转换成字符串失败")
	}

	con := c.PostForm("status")
	s := *NewStatus(con, author)
	if err := s.CreateStatus(); err != nil {
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