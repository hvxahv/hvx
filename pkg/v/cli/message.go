package cli

import (
	msg "github.com/hvxahv/hvx/api/grpc/proto/message/v1alpha1"
	"google.golang.org/grpc"
)

type Message interface {
	msg.MessagesClient
}

type message struct {
	msg.MessagesClient
}

func NewMessage(conn *grpc.ClientConn) Message {
	return &message{msg.NewMessagesClient(conn)}
}
