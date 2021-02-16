package social

import (
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/api-client"
	"log"
)

// CreateArticleClient 创建文章的客户端，连接到 Accounts 服务端
// 提交请求数据并获取服务端返回的结果数据 r.Reply
func CreateArticleClient(data *pb.ArticleData) (string, error) {
	p := viper.GetString("port.articles")
	conn, err := api_client.Conn(p, "Article")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()


	cli := pb.NewArticlesClient(conn)

	r, err := cli.NewArticle(context.Background(), data)
	if err != nil {
		log.Printf("创建内容错误，发送消息给 Status 服务端失败: %v", err)
	}
	return r.Reply, err
}

// GetArticlesClient 通过用户名获取数据
func GetArticlesClient(data *pb.GetArticleData) ([]*pb.ArticleData, error) {
	p := viper.GetString("port.articles")
	conn, err := api_client.Conn(p, "Article")
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()


	cli := pb.NewArticlesClient(conn)


	r, err := cli.GetArticles(context.Background(), data)
	if err != nil {
		log.Printf("GetArticles 客户端方法失败，无法获取文章: %v", err)
	}
	return r.Articles, err
}
//
//// UpdateStatusListHandler 更新状态 Handler
//func UpdateArticleClient(author string) {
//	p := viper.GetString("port.articles")
//	conn, err := client.Conn(p, "Article")
//	if err != nil {
//		log.Println(err)
//	}
//	defer conn.Close()
//
//	cli := pb.NewArticlesClient(conn)
//	r, err := cli.UpdateArticle(context.Background(), &pb.ArticleData{
//		Id: "123123",
//		Author: author,
//		Article: "这是我的第一条 Status 的修改内容（开始更新状态）",
//	})
//	if err != nil {
//		log.Printf("更新状态错误，发送消息给 Status 服务端失败: %v", err)
//	}
//	log.Println("更新状态返回的数据", r.Reply)
//}
//
//// DeleteStatusHandler 删除状态 Handler 通过 ID
//func DeleteArticleClient(id string) {
//	p := viper.GetString("port.articles")
//	conn, err := client.Conn(p, "Article")
//	if err != nil {
//		log.Println(err)
//	}
//	defer conn.Close()
//
//	cli := pb.NewArticlesClient(conn)
//	r, err := cli.DeleteArticle(context.Background(), &pb.DeleteArticleByID{
//		Id: id,
//	})
//	if err != nil {
//		log.Printf("删除状态错误，发送消息给 Status 服务端失败: %v", err)
//	}
//	log.Println("删除状态返回的数据", r.Reply)
//}
//
