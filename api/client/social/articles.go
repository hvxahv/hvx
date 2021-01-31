package social

import (
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/pkg/structs"
	"log"
)

// CreateArticleClient 创建文章的客户端，连接到 Accounts 服务端
// 提交请求数据并获取服务端返回的结果数据 r.Reply
func CreateArticleClient(data *structs.Articles) (string, error) {
	addr := fmt.Sprintf("localhost:%s", viper.GetString("port.articles"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("连接到 Article 服务失败: %v", err)
	}
	defer conn.Close()
	cli := pb.NewArticlesClient(conn)

	d := &pb.ArticleData{
		Author: data.Author,
		Article: data.Article,
	}
	r, err := cli.NewArticle(context.Background(), d)
	if err != nil {
		log.Printf("新建状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	return r.Reply, err
}

// UpdateStatusListHandler 更新状态 Handler
func UpdateArticleClient(author string) {
	addr := fmt.Sprintf("localhost:%s", viper.GetString("port.articles"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status services.bac: %v", err)
	}
	defer conn.Close()
	cli := pb.NewArticlesClient(conn)
	r, err := cli.UpdateArticle(context.Background(), &pb.ArticleData{
		Id: "123123",
		Author: author,
		Article: "这是我的第一条 Status 的修改内容（开始更新状态）",
	})
	if err != nil {
		log.Printf("更新状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	log.Println("更新状态返回的数据", r.Reply)
}

// DeleteStatusHandler 删除状态 Handler 通过 ID
func DeleteArticleClient(id string) {
	addr := fmt.Sprintf("localhost:%s", viper.GetString("port.articles"))
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status services.bac: %v", err)
	}
	defer conn.Close()
	cli := pb.NewArticlesClient(conn)
	r, err := cli.DeleteArticle(context.Background(), &pb.DeleteArticleByID{
		Id: id,
	})
	if err != nil {
		log.Printf("删除状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	log.Println("删除状态返回的数据", r.Reply)
}

