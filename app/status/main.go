package main

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/kernel/v1"
	"hvxahv/app/status/app"
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

	port := viper.GetString("port.status")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
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
