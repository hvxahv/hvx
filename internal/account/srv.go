package account

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type account struct {
	pb.AccountsServer
	pb.ActorsServer
	pb.ECDHServer
	*Accounts
	*Actors
}

// Run creates a new account server, and starts it.
// that implements the gRPC server interface,
// and returns a pointer to the server.
func Run() error {
	name := "account"
	port := viper.GetString("microservices.account.port")
	log.Printf("App %s Started at %s\n", name, time.Now())

	s := grpc.NewServer()

	// Create a new account and actor server.
	pb.RegisterAccountsServer(s, &account{})
	pb.RegisterActorsServer(s, &account{})
	pb.RegisterECDHServer(s, &account{})
	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
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
