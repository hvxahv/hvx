package saved

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/saved/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type saved struct {
	pb.SavedServer
	*Saves
}

func Run() error {
	name := "saved"
	port := viper.GetString("microservices.saved.port")
	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()

	pb.RegisterSavedServer(s, &saved{})
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
