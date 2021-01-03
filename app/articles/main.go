package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/app/articles/app"
	"hvxahv/pkg/bot"
	"log"
	"net"
)

func main()  {
	app.InitDB()

	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	port := viper.GetString("port.articles")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Article gRPC Services Failed to Listen: %v", err)
		return
	} else {
		log.Println("Article gRPC Services is running", port)
		go bot.ServicesRunningNotice("article", port)
	}
	s := grpc.NewServer()
	pb.RegisterArticlesServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Article gRPC Services failed to start: %v", err)
		return
	}
}
