package channel

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const serviceName = "channel"

type channel struct {
	pb.ChannelServiceServer
	pb.AdministrativeServiceServer
	pb.SubscriberServiceServer
	pb.BroadcastServiceServer
	*Channels
	*Administrates
	*Subscribes
}

// Run starts the server. It will block until the server is shutdown.
// If the server fails to start, it will return an error.
func Run() error {
	log.Printf("App %s Started at %s\n", serviceName, time.Now())

	// Create a new server instance.
	s := grpc.NewServer()

	pb.RegisterChannelServiceServer(s, &channel{})
	pb.RegisterAdministrativeServiceServer(s, &channel{})

	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", microservices.NewService(serviceName).GetPort()))
	if err != nil {
		return err
	}

	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println(err)
			return
		}
	}()

	return nil
}
