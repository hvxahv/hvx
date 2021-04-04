package mw

import (
	"github.com/gin-gonic/gin"
	"log"
)

// GetUserName 通过 loginUser 这个 key 获取到 context 登录用户的用户名
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
