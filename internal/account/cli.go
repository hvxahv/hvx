package account

import (
	pb "github.com/hvxahv/hvxahv/api/account/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
)

func NewAccountClient() (pb.AccountsClient, error) {
	conn, err := grpc.Dial(microservices.GetAccountAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewAccountsClient(conn), nil
}

func NewActorClient() (pb.ActorClient, error) {
	conn, err := grpc.Dial(microservices.GetAccountAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewActorClient(conn), nil
}

func NewDeviceClient() (pb.DevicesClient, error) {
	conn, err := grpc.Dial(microservices.GetAccountAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewDevicesClient(conn), nil
}
