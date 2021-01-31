package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/client/social"
	"hvxahv/pkg/response"
	"hvxahv/pkg/structs"
	"hvxahv/pkg/utils"
	"log"
)

// CreateArticleHandler 创建文章的 Handler，接收 http 数据请求
// 将数据处理后发送给 accounts 微服务的客户端并获得客户端返回的接收：string 类型的 r
func CreateArticleHandler(c *gin.Context) {
	author := utils.GetUserName(c)
	article := c.PostForm("article")

	data := &structs.Articles{
		Article: article,
		Author: author,
	}

	r, err := social.CreateArticleClient(data)
	if err != nil {
		log.Println(err)
	} else {
		response.CreateArticleResponse(c, r)
	}
}

// UpdateArticleHandler ...
func UpdateArticleHandler(c *gin.Context) {
	author := utils.GetUserName(c)
	social.UpdateArticleClient(author)
}
// DeleteArticleHandler ...
func DeleteArticleHandler(c *gin.Context) {
	social.DeleteArticleClient("4124141241")
}

// GetArticlesHandler 获取用户文章的 handler
// 它接收一个参数用于区分获取类型， article or status （文章 或者 状态）
func GetArticlesHandler(c *gin.Context) {
	author := utils.GetUserName(c)
	status := c.PostForm("status")
	isA := false
	if status == "1" {
		isA = true
	}
	data := &structs.Articles{
		Author:author,
		Status: isA,
	}

	log.Println(data.Status)

}