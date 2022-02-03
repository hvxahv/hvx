package notify

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/notify/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type notify struct {
	pb.NotifyServer
	*Notifies
}

func Run() error {
	name := "notify"
	port := viper.GetString("microservices.notify.port")
	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()

	pb.RegisterNotifyServer(s, &notify{})
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
