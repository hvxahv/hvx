package channel

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type server struct {
	pb.ChannelServer
}

func Run() error {
	name := "Channel"
	port := viper.GetString("microservices.channel.port")

	log.Printf("app %s started at %s\n", name, time.Now())

	s := grpc.NewServer()
	pb.RegisterChannelServer(s, &server{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	log.Printf("%s gRPC Services is running.., Port: %s.", name, port)

	if err2 := s.Serve(lis); err != nil {
		return err2
	}

	return nil
}
