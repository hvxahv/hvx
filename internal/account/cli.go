package account

import (
	"fmt"
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func NewAccountClient() (pb.AccountsClient, error) {
	address := fmt.Sprintf("%s:%s", viper.GetString("microservices.account.localhost"), viper.GetString("microservices.account.port"))

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewAccountsClient(conn), nil
}

func NewActorClient() (pb.ActorsClient, error) {
	address := fmt.Sprintf("%s:%s", viper.GetString("microservices.account.localhost"), viper.GetString("microservices.account.port"))
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewActorsClient(conn), nil
}
