package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/app/articles/app"
	"hvxahv/pkg/bot"
	"log"
	"net"
)
const (
	port = ":8010"
)
func main()  {
	app.InitDB()

	lis, err := net.Listen("tcp", port)
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
