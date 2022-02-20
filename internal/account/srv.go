package account

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

type account struct {
	pb.AccountsServer
	pb.ActorServer
	pb.AuthServer
	pb.DevicesServer
	*Accounts
	*Actors
	*Devices
}

// Run starts the server. It will block until the server is shutdown. If the server fails to start, it will return an error.
func Run() error {
	name := "account"
	log.Printf("App %s Started at %s\n", name, time.Now())

	// Create a new server instance.
	s := grpc.NewServer()

	pb.RegisterAccountsServer(s, &account{})
	pb.RegisterActorServer(s, &account{})
	pb.RegisterAuthServer(s, &account{})
	pb.RegisterDevicesServer(s, &account{})

	reflection.Register(s)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", microservices.GetAccountAddress()))
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
