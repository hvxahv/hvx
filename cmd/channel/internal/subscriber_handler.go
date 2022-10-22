package internal

import (
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

// Subscriber ...

func (s *server) AddSubscriber(ctx context.Context, in *pb.AddSubscriberRequest) (*pb.AddSubscriberResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	var (
		adminId = parse.ActorId
	)

	if err := NewSubscribe(uint(in.GetChannelId()), uint(in.GetSubscriberId())).AddSubscriber(adminId); err != nil {
		return nil, err
	}
	return &pb.AddSubscriberResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) RemoveSubscriber(ctx context.Context, in *pb.RemoveSubscriberRequest) (*pb.RemoveSubscriberResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	var (
		adminId = parse.ActorId
	)

	if err := NewSubscribe(uint(in.GetChannelId()), uint(in.GetRemovedId())).RemoveSubscriber(adminId); err != nil {
		return nil, err
	}
	return &pb.RemoveSubscriberResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) GetSubscribers(ctx context.Context, in *pb.GetSubscribersRequest) (*pb.GetSubscribersResponse, error) {
	if in.GetAdminId() == 0 {
		parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
		if err != nil {
			return nil, err
		}
		in.AdminId = int64(parse.GetActorId())
	}

	channelId, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	subscribers, err := NewSubscriberChannelId(uint(channelId)).GetSubscribers(uint(in.GetAdminId()))
	if err != nil {
		return nil, err
	}

	var actors []*actor.ActorData
	for _, sub := range subscribers {
		a, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(sub.SubscriberId))
		if err != nil {
			return nil, err
		}
		actors = append(actors, &actor.ActorData{
			Id:                a.Actor.Id,
			PreferredUsername: a.Actor.PreferredUsername,
			Domain:            a.Actor.Domain,
			Avatar:            a.Actor.Avatar,
			Name:              a.Actor.Name,
			Summary:           a.Actor.Summary,
			Inbox:             a.Actor.Inbox,
			Address:           a.Actor.Address,
			PublicKey:         a.Actor.PublicKey,
			ActorType:         a.Actor.ActorType,
			IsRemote:          a.Actor.IsRemote,
		})
	}

	return &pb.GetSubscribersResponse{
		Code:       "200",
		Status:     "ok",
		Subscriber: actors,
	}, nil
}

func (s *server) Subscription(ctx context.Context, in *pb.SubscriptionRequest) (*pb.SubscriptionResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewSubscribe(uint(in.GetChannelId()), parse.ActorId).Subscription(); err != nil {
		return nil, err
	}

	return &pb.SubscriptionResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) Unsubscribe(ctx context.Context, in *pb.UnsubscribeRequest) (*pb.UnsubscribeResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewSubscribe(uint(in.GetChannelId()), parse.ActorId).Unsubscribe(); err != nil {
		return nil, err
	}

	return &pb.UnsubscribeResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}
