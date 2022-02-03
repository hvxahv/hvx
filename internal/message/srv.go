package message

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/message/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type message struct {
	pb.MessageServer
	*Matrices
}

func Run() error {
	name := "device"
	port := viper.GetString("microservices.message.port")
	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()

	pb.RegisterMessageServer(s, &message{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println(err)
			return
		}
	}()

	return nil
}
