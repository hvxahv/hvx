package message

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/hvxahv/hvxahv/api/message/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "message"

type message struct {
	pb.MessagesServer
	*Matrices
}

func Run() error {
	name := "device"

	log.Printf("App %s Started at %s\n", name, time.Now())
	s := grpc.NewServer()

	pb.RegisterMessagesServer(s, &message{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", microservices.NewService(serviceName).GetPort()))
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
