package message

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/message/v1alpha1"
	"github.com/hvxahv/hvx/clientv1"
)

type Message interface {
	pb.MessagesClient
}

type message struct {
	pb.MessagesClient
}

func NewMessage(c *clientv1.Client) Message {
	return &message{
		pb.NewMessagesClient(c.conn),
	}
}
