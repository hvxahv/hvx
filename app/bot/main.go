package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	bot2 "hvxahv/api/bot"
	"hvxahv/pkg/bot"
	"log"
	"net"
)

type server struct {
	bot2.BotNoticeServer
}

// 实现 Telegram Bot 的服务接口，供外部调用。
func main() {

	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	port := viper.GetString("port.bot")

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		fmt.Printf("Failed to Listen: %v", err)
		return
	} else {
		log.Println("Bot gRPC service is running", port)
		go bot.ServicesRunningNotice("bot", port)
	}
	s := grpc.NewServer()
	bot2.RegisterBotNoticeServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Bot gRPC service failed to start: %v", err)
		return
	}
}
