package microservice

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"log"
	"net"
	"time"
)

func (ms *microservice) ArticlesServer() error {
	log.Printf("App %s Started at %s\n", ms.Name, time.Now())
	s := grpc.NewServer()
	pb.RegisterArticlesServer(s, &server{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", ms.Port))
	if err != nil {
		return err
	}

	log.Printf("%s gRPC Services is running, Port: %s.", ms.Name, ms.Port)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
