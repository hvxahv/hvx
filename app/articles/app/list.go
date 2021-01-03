package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

// ShowStatusListHandler ...
func ShowArticleListHandler(c *gin.Context) {
	name, ok := c.Get("loginUser")
	if !ok {
		log.Println("获取用户名失败")
	}
	author, ok := name.(string)
	if !ok {
		log.Println("用户名转换成字符串失败")
	}

	r, err := ShowArticleLis(author)
	if err != nil {
		log.Println("Query Status Errors", err)
	}
	c.JSON(200, gin.H{
		"state": "200",
		"status": r,
	})
}

// ShowStatusLis ...
func ShowArticleLis(author string) (*[]Article, error) {
	var s *[]Article
	if db2.Debug().Table("status").Where("author = ?", author).Find(&s).RecordNotFound() {
		return nil, errors.New("未找到个人中心的文章")
	}
	return s, nil


}