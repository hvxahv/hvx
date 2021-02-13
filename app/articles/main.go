package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/db"
	"log"
	"net"
)
/**
	Articles 服务的服务端实现
	获取配置文件并初始化数据库
 */
func main()  {
	if err := db.InitMongoDB(); err != nil {
		log.Println(err)
	}
	// 初始化数据库
	if err := db.InitMariaDB(); err != nil {
		log.Println("数据库初始化失败：", err)
	}
	// 获取配置文件
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	port := viper.GetString("port.articles")

	// 开始启动微服务
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Article gRPC Services Failed to Listen: %v", err)
		return
	} else {
		log.Println("Article gRPC Services is running", port)
		// 通知 Bot 服务已经开启
		go bot.ServicesRunningNotice("activity", port)
	}
	s := grpc.NewServer()
	pb.RegisterArticlesServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Article gRPC Services failed to start: %v", err)
		return
	}
}
