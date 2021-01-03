package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/app/status/app"
	"hvxahv/pkg/bot"
	"log"
	"net"
)
const (
	port = ":8005"
)

func main()  {
	app.InitDB()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("Status gRPC Services Failed to Listen: %v", err)
		return
	} else {
		log.Println("Status gRPC Services is running", port)
		go bot.ServicesRunningNotice("status", port)
	}
	s := grpc.NewServer()
	pb.RegisterStatusServer(s, &server{})
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		fmt.Printf("Status gRPC Services failed to start: %v", err)
		return
	}
}
