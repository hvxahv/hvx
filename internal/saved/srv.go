package saved

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type saved struct {
	pb.SavedServer
	*Saves
}

func Run() error {
	name := "saved"
	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()

	pb.RegisterSavedServer(s, &saved{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", x.NewService("saved").GetPort()))
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
