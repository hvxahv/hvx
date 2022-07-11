package channel

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/channel"
	"github.com/hvxahv/hvx/clientv1"
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

func NewChannel(c *clientv1.Client) Channel {
	return &channel{
		AdministrativeClient: pb.NewAdministrativeClient(c.conn),
		BroadcastClient:      pb.NewBroadcastClient(c.conn),
		ChannelClient:        pb.NewChannelClient(c.conn),
		SubscriberClient:     pb.NewSubscriberClient(c.conn),
	}
}
