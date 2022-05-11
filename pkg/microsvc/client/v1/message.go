package clientv1

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/message/v1alpha1"
)

type Message interface {
	pb.MessagesClient
}

type message struct {
	pb.MessagesClient
}

func NewMessage(c *Client) Message {
	return &message{
		pb.NewMessagesClient(c.conn),
	}
}
