package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"strconv"
)

type Channel interface {
	GetPrivateKeyByActorId(actorId int64) (*pb.GetPrivateKeyByActorIdResponse, error)
}

type Subscribe interface {
	GetSubscribers(channelId, adminId int64) (*pb.GetSubscribersResponse, error)
}

func (svc *Svc) GetSubscribers(channelId, adminId int64) (*pb.GetSubscribersResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()

	subscribers, err := pb.NewSubscriberClient(c.Conn).GetSubscribers(svc.ctx, &pb.GetSubscribersRequest{
		ChannelId: strconv.Itoa(int(channelId)),
		AdminId:   adminId,
	})
	if err != nil {
		return nil, err
	}
	return subscribers, nil
}

func (svc *Svc) GetPrivateKeyByActorId(actorId int64) (*pb.GetPrivateKeyByActorIdResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	subs, err := pb.NewChannelClient(c.Conn).GetPrivateKeyByActorId(svc.ctx, &pb.GetPrivateKeyByActorIdRequest{
		ActorId: actorId,
	})
	if err != nil {
		return nil, err
	}
	return subs, nil
}
