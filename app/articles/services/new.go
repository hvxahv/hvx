package services

import (
	"errors"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/db"
	"hvxahv/pkg/models"
)

// CreateArticleHandler Articles 微服务服务端创建文章的 Handler
// 将数据库执行的结果发送给客户端，返回 string 类型的 error 或者 ok
func CreateArticleHandler(in *pb.ArticleData) string {

	a := &models.Articles{
		Author:    in.Author,
		Content:   in.Content,
		ContentType: in.Type,
	}
	d := a.NewArticle()

	db := db.GetMaria()
	db.AutoMigrate(*d)
	if err := db.Debug().Table("articles").Create(&d).Error; err != nil {
		errors.New("Failed to write new article to db... ")
		return "error"
	} else {
		// 将通过 http 发送到我的关注人
		go SendActivity(d.ID, in.Author, in.Content)
		return "ok"
	}
}

