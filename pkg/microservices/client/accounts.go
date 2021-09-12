package client

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	errors "github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

type client struct {
	name string
	host string
	port string
	addr string
}

func NewClient(name string) *client {
	host := viper.GetString(fmt.Sprintf("microservices.%s.host", name))
	port := viper.GetString(fmt.Sprintf("microservices.%s.port", name))
	addr := fmt.Sprintf("%s:%s", host, port)


	return &client{name: name, host: host, port: port, addr: addr}
}

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

// FetchAccountIdByName Fetch actor id by name from accounts service.
func FetchAccountIdByName(name string) uint {
	cli, conn, err := Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	acct, err := cli.QueryByName(context.Background(), &pb.NewAccountByName{
		Username: name,
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}
	return uint(acct.Id)
}

// FetchAccountNameByID Fetch actor id by name from accounts service.
func FetchAccountNameByID(id uint) (string, error) {
	cli, conn, err := Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	acct, err := cli.QueryByID(context.Background(), &pb.NewAccountByID{
		Id: uint64(id),
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
		return "", errors.New("ACCOUNT_NOT_FOUND")
	}
	return acct.Username, nil
}


