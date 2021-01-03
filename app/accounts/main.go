package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/app/accounts/app"
	"hvxahv/pkg/bot"
	"log"
	"net"
)

const (
	port = ":8000"
)

func main()  {
	app.InitDB()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Accounts gRPC Services Failed to Listen: %v", err)
		return
	} else {
		log.Println("Accounts gRPC Services is running", port)
		go bot.ServicesRunningNotice("account", port)
	}
	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Accounts gRPC Services failed to start: %v", err)
		return
	}
}
