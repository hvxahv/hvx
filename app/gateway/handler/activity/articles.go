package activity

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	pb "hvxahv/api/hvxahv/v1alpha1"
	social2 "hvxahv/internal/client/social"
	"hvxahv/pkg/activitypub/activity"
	"hvxahv/pkg/mw"
	"log"
)

/**
NewArticleHandler 创建一个状态的 handler
它通过 GetUserName 这个方法获取当前的用户
content 接收发布的内容
type 是接收内容的类型, 接收一个参数用于区分获取类型
需要传递 activity 和 status 这两个类型, 其他类型不接收!
*/
func NewArticleHandler(c *gin.Context) {
	author := mw.GetUserName(c)
	content := c.PostForm("content")
	t := c.PostForm("type")

	if t == "activity" || t == "status" {
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

		r, err := social2.CreateArticleClient(data)
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
// GetPublicArticle ...
func GetPublicArticleHandler(c *gin.Context) {
	name := c.Param("user")
	id := c.Param("id")
	log.Printf("通过 ID %s 查找 %s 发布的文章", id, name)

	aid := fmt.Sprintf("http://%s/u/%s/%s", viper.GetString("activitypub"), name, id)

	activity.GetPublicArticleById(aid, c)


}


// GetArticles ...
func GetArticles(c *gin.Context) {
	name := mw.GetUserName(c)
	address := fmt.Sprintf("http://%s/u/%s", viper.GetString("activitypub"), name)

	r := activity.GetArticleByName(address)

	c.JSON(200, gin.H{
		"article_lens": len(r),
		"articles": r,
	})
}

