package handler

import (
	"github.com/gin-gonic/gin"
	"hvxahv/api/client/social"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/utils"
	"log"
)

/**
NewArticleHandler 创建一个状态的 handler
它通过 GetUserName 这个方法获取当前的用户
content 接收发布的内容
type 是接收内容的类型, 接收一个参数用于区分获取类型
需要传递 article 和 status 这两个类型, 其他类型不接收!
*/
func NewArticleHandler(c *gin.Context) {
	author := utils.GetUserName(c)
	content := c.PostForm("content")
	t := c.PostForm("type")

	if t == "article" || t == "status" {
		data := &pb.ArticleData{
			Author: author,
			Content: content,
			Type: t,
		}
		log.Println("-------- 接收内容的 Handler ---------")
		log.Println(data.Content)
		log.Println(author)
		log.Println(t)
		log.Println("-----------------")

		r, err := social.CreateArticleClient(data)
		if err != nil {
			log.Println(err)
		}
		log.Println("---0-- 创建文章返回的数据 -0---> ", r)

	} else {
		c.JSON(503, gin.H{
			"message": "参数类型不正确",
		})
	}


}

// CreateArticleHandler 创建文章的 Handler，接收 http 数据请求
// 将数据处理后发送给 accounts 微服务的客户端并获得客户端返回的接收：string 类型的 r
//func CreateArticleHandler(c *gin.Context) {
//	author := utils.GetUserName(c)
//	article := c.PostForm("article")
//
//	data := &models.Articles{
//		Article: article,
//		Author: author,
//	}
//
//	r, err := social.CreateArticleClient(data)
//	if err != nil {
//		log.Println(err)
//	} else {
//		response.CreateArticleResponse(c, r)
//	}
//}

//// UpdateArticleHandler ...
//func UpdateArticleHandler(c *gin.Context) {
//	author := utils.GetUserName(c)
//	social.UpdateArticleClient(author)
//}
//// DeleteArticleHandler ...
//func DeleteArticleHandler(c *gin.Context) {
//	social.DeleteArticleClient("4124141241")
//}
