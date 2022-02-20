package message

import (
	pb "github.com/hvxahv/hvxahv/api/message/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/microservices"
	"google.golang.org/grpc"
)

func NewMessageClient() (pb.MessagesClient, error) {
	conn, err := grpc.Dial(microservices.GetMessageAddress(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	return pb.NewMessagesClient(conn), nil
}
