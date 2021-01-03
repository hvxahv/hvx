package app

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
)

func CreateArticleHandler(c *gin.Context) {
	name, ok := c.Get("loginUser")
	if !ok {
		log.Println("获取用户名失败")
	}
	author, ok := name.(string)
	if !ok {
		log.Println("用户名转换成字符串失败")
	}

	con := c.PostForm("article")
	s := *NewArticle(con, author)
	if err := s.CreateArticle(); err != nil {
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

func (s *Article) CreateArticle() error {
	db2.AutoMigrate(Article{})
	if err := db2.Debug().Table("article").Create(&s).Error; err != nil {
		return errors.New("Failed to write new status to database... ")
	}
	return nil
}
