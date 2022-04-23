package activity

import (
	pb "github.com/hvxahv/hvxahv/api/activity/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func GetActivityClient() (pb.ActivityClient, error) {
	conn, err := grpc.Dial(x.NewService(serviceName).GetAddress(),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewActivityClient(conn), nil
}
