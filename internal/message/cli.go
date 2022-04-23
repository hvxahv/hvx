package message

import (
	pb "github.com/hvxahv/hvxahv/api/message/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/x"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewMessageClient() (pb.MessagesClient, error) {
	conn, err := grpc.Dial(x.NewService(serviceName).GetAddress(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewMessagesClient(conn), nil
}
