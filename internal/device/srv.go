package device

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type device struct {
	pb.DevicesServer
	*Devices
}

// Run creates a new account server, and starts it.
// that implements the gRPC server interface,
// and returns a pointer to the server.
func Run() error {
	name := "device"
	port := viper.GetString("microservices.device.port")
	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()

	// Create a new account and actor server.
	pb.RegisterDevicesServer(s, &device{})
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
