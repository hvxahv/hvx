package channel

import (
	pb "github.com/hvxahv/hvxahv/api/channel/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetChannelClient() (pb.ChannelServiceClient, error) {
	conn, err := grpc.Dial(microservices.NewService(serviceName).GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewChannelServiceClient(conn), nil
}

func GetChannelAdminClient() (pb.AdministrativeServiceClient, error) {
	conn, err := grpc.Dial(microservices.NewService(serviceName).GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewAdministrativeServiceClient(conn), nil
}

func GetChannelBroadcastClient() (pb.BroadcastServiceClient, error) {
	conn, err := grpc.Dial(microservices.NewService(serviceName).GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewBroadcastServiceClient(conn), nil
}

func GetChannelSubscriberClient() (pb.SubscriberServiceClient, error) {
	conn, err := grpc.Dial(microservices.NewService(serviceName).GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewSubscriberServiceClient(conn), nil
}
