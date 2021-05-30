package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/pkg/bot"
	"hvxahv/pkg/maria"
	"hvxahv/pkg/mongo"
	"hvxahv/pkg/redis"
	"log"
	"net"
)

type server struct {
	pb.OutboxServer
}

func main() {
	if err := mongo.InitMongoDB(); err != nil {
		log.Println(err)
	}
	if err := maria.InitMariaDB(); err != nil {
		log.Println("数据库初始化失败：", err)
	}
	redis.InitRedis()

	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	port := viper.GetString("port.outbox")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Outbox gRPC Services Failed to Listen: %v", err)
		return
	} else {
		log.Println("Outbox gRPC Services is running", port)
		go bot.ServicesRunningNotice("Outbox", port)
	}
	s := grpc.NewServer()
	pb.RegisterOutboxServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Outbox gRPC Services failed to start: %v", err)
		return
	}
}
