package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	v1 "hvxahv/api/util/v1"
	"hvxahv/pkg/bot"
	"log"
	"net"
)

type server struct {
	v1.BotNoticeServer
}

const (
	port = ":9000"
)
// 实现 Telegram Bot 的服务接口，供外部调用。
func main()  {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Failed to Listen: %v", err)
		return
	} else {
		log.Println("Bot gRPC service is running", port)
		go bot.ServicesRunningNotice("bot", port)
	}
	s := grpc.NewServer()
	v1.RegisterBotNoticeServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Bot gRPC service failed to start: %v", err)
		return
	}
}