package cli

import (
	ch "github.com/hvxahv/hvx/api/grpc/proto/channel/v1alpha1"
	"google.golang.org/grpc"
)

type Channel interface {
	ch.AdministrativeClient
	ch.BroadcastClient
	ch.ChannelClient
	ch.SubscriberClient
}

type channel struct {
	ch.AdministrativeClient
	ch.BroadcastClient
	ch.ChannelClient
	ch.SubscriberClient
}

func NewChannel(conn *grpc.ClientConn) Channel {
	return &channel{
		AdministrativeClient: ch.NewAdministrativeClient(conn),
		BroadcastClient:      ch.NewBroadcastClient(conn),
		ChannelClient:        ch.NewChannelClient(conn),
		SubscriberClient:     ch.NewSubscriberClient(conn),
	}
}
