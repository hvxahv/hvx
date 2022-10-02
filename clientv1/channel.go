package clientv1

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
)

type Channel interface {
	GetSubscribers(channelId string) (*pb.GetSubscribersResponse, error)
	GetPrivateKeyByActorId(actorId int64) (*pb.GetPrivateKeyByActorIdResponse, error)
}

func (svc *Svc) GetSubscribers(channelId string) (*pb.GetSubscribersResponse, error) {
	c, err := NewClient(svc.ctx, svc.address)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	subs, err := pb.NewSubscriberClient(c.Conn).GetSubscribers(svc.ctx, &pb.GetSubscribersRequest{
		ChannelId: channelId,
	})
	if err != nil {
		return nil, err
	}
	return subs, nil
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
