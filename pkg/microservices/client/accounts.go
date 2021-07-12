package client

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

// Accounts Microservice client for Accounts.
func Accounts() (pb.AccountsClient, *grpc.ClientConn, error) {
	name := "accounts"

	host := viper.GetString(fmt.Sprintf("microservices.%s.host", name))
	port := viper.GetString(fmt.Sprintf("microservices.%s.port", name))
	addr := fmt.Sprintf("%s:%s", host, port)

	// Connect to gPRC service.
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to %s service: %v", name , err)
		return nil, nil, nil
	}

	cli := pb.NewAccountsClient(conn)

	return cli, conn, nil
}
