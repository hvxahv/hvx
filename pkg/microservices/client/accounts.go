package client

import (
	"fmt"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"log"
)

// Accounts Microservice client for Accounts.
func Accounts() (*grpc.ClientConn, error) {
	name := "accounts"

	host := viper.GetString(fmt.Sprintf("microservices.%s.host", name))
	port := viper.GetString(fmt.Sprintf("microservices.%s.port", name))
	addr := fmt.Sprintf("%s:%s", host, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Failed to connect to %s service: %v", name , err)
		return nil, err
	}


	return conn, nil
}
