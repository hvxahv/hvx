package account

import (
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
)

func NewAccountClient() (pb.AccountsClient, error) {
	conn, err := grpc.Dial(microservices.NewService("account").GetAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewAccountsClient(conn), nil
}

func NewAuthClient() (pb.AuthClient, error) {
	conn, err := grpc.Dial(microservices.NewService("account").GetAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return pb.NewAuthClient(conn), nil
}

func NewActorClient() (pb.ActorClient, error) {
	conn, err := grpc.Dial(microservices.NewService("account").GetAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewActorClient(conn), nil
}
