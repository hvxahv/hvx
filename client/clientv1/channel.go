package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/channel/v1alpha1"
)

type Channel interface {
	pb.AdministrativeClient
	pb.BroadcastClient
	pb.ChannelClient
	pb.SubscriberClient
}

type channel struct {
	pb.AdministrativeClient
	pb.BroadcastClient
	pb.ChannelClient
	pb.SubscriberClient
}

func NewChannel(c *Client) Channel {
	return &channel{
		AdministrativeClient: pb.NewAdministrativeClient(c.conn),
		BroadcastClient:      pb.NewBroadcastClient(c.conn),
		ChannelClient:        pb.NewChannelClient(c.conn),
		SubscriberClient:     pb.NewSubscriberClient(c.conn),
	}
}
