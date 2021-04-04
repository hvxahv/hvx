/**
	Inbox 的 gRPC services
	作为全局的唯一收件箱，用于接收其他用户活动
 */
package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/mongo"
	"hvxahv/pkg/maria"
	"hvxahv/pkg/redis"
	"log"
	"net"
)

type server struct {
	pb.InboxServer
}

func main() {
	if err := mongo.InitMongoDB(); err != nil {
		log.Println(err)
	}
	redis.InitRedis()
	if err := maria.InitMariaDB(); err != nil {
		log.Println("数据库初始化失败：", err)
	}

	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	p := viper.GetString("port.inbox")

	// 开启服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", p))
	if err != nil {
		fmt.Printf("Accounts gRPC Services Failed to Listen: %v", err)
		return
	} else {
		// 在控制台中打印服务启动 log 通知 并启动一个协程通知 Bot Account 服务已经开启
		log.Println("Accounts gRPC Services is running....", p)
		go bot.ServicesRunningNotice("inbox", p)
	}

	s := grpc.NewServer()
	pb.RegisterInboxServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Printf("Inbox gRPC 服务启动失败: %v", err)
	}
}


