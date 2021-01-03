package auth

import (
	"github.com/gin-gonic/gin"
	"log"
)

func GetUserName(c *gin.Context) string {
	name, ok := c.Get("loginUser")
	if !ok {
		log.Println("获取用户名失败")
	}
	author, ok := name.(string)
	if !ok {
		log.Println("用户名转换成字符串失败")
	}

	return author
}
