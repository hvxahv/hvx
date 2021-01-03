package social

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/pkg/auth"
	"log"
)

// CreateStatusHandler 创建状态 handler
func CreateArticleHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	conn, err := grpc.Dial(viper.GetString("port.article"), grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status app: %v", err)
	}
	defer conn.Close()
	cli := pb.NewArticlesClient(conn)
	r, err := cli.NewArticle(context.Background(), &pb.ArticleData{
		Author: author,
		Article: "这是我的第一条 文章",
	})
	if err != nil {
		log.Printf("新建状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	log.Println("创建状态返回的数据", r.Reply)
}

// UpdateStatusListHandler 更新状态 Handler
func UpdateArticleHandler(c *gin.Context) {
	author := auth.GetUserName(c)
	conn, err := grpc.Dial(viper.GetString("port.article"), grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status app: %v", err)
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
func DeleteArticleHandler(c *gin.Context) {
	conn, err := grpc.Dial(viper.GetString("port.article"), grpc.WithInsecure())
	if err != nil {
		log.Printf("Faild to connect to Status app: %v", err)
	}
	defer conn.Close()
	cli := pb.NewArticlesClient(conn)
	r, err := cli.DeleteArticle(context.Background(), &pb.DeleteArticleByID{
		Id: "123123444455622",
	})
	if err != nil {
		log.Printf("删除状态错误，发送消息给 Status 服务端失败: %v", err)
	}
	log.Println("删除状态返回的数据", r.Reply)
}

