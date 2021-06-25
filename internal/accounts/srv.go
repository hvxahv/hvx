package accounts

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"log"
	"net"
	"time"
)

type server struct {
	pb.AccountsServer
}

// Run Run the accounts gRPC service.
func Run() error {
	name := "Accounts"
	port := viper.GetString("microservices.accounts.port")

	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()
	pb.RegisterAccountsServer(s, &server{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		return err
	}

	log.Printf("%s gRPC Services is running.., Port: %s.", name, port)

	if err := s.Serve(lis); err != nil {
		return err
	}

	return nil
}
