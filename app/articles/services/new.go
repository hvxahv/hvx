package services

import (
	"errors"
	"hvxahv/pkg/db"
	"hvxahv/pkg/models"
)

// CreateArticleHandler Articles 微服务服务端创建文章的 Handler
// 将数据库执行的结果发送给客户端，返回 string 类型的 error 或者 ok
func CreateArticleHandler(author, con string) string {
	a := &models.Articles{
		Article: con,
		Author: author,
	}

	s := a.NewArticle()
	db := db.GetMaria()
	db.AutoMigrate(*s)
	if err := db.Debug().Table("articles").Create(&s).Error; err != nil {
		errors.New("Failed to write new article to db... ")
		return "error"
	} else {
		return "ok"
	}
}

