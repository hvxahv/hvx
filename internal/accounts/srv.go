package accounts

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type server struct {
	pb.AccountsServer
}

// Run accounts gRPC service.
func Run() error {
	name := "accounts"
	port := viper.GetString("microservices.accounts.port")

	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, &server{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	go func() {
		err2 := s.Serve(lis)
		if err2 != nil {
			log.Println(err2)
			return
		}
	}()

	return nil
}
