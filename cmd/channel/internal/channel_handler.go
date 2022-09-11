package internal

import (
	"context"
	"strconv"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/hvxahv/hvx/rsa"
)

func (s *server) CreateChannel(ctx context.Context, in *pb.CreateChannelRequest) (*pb.CreateChannelResponse, error) {
	// TODO - IMPROVED: When calling the Create method (Create()),
	// the method should check if the Actor PreferredUsername exists instead
	// of calling Create() after calling the IsExist() method when creating the Actor.
	// This problem is revealed in the current scenario.
	_ := clientv1.New(ctx, microsvc.NewGRPCAddress("actor").Get())
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	exist, err := actor.NewActorClient(cli.Conn).IsExist(ctx, &actor.IsExistRequest{
		PreferredUsername: in.GetPreferredUsername(),
	})
	if err != nil {
		return nil, err
	}
	if !exist.IsExist {
		return nil, errors.New(errors.ErrChannelAlready)
	}

	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	// Generate rsa Key.
	rsa, err := rsa.NewRsa(2048).Generate()
	if err != nil {
		return nil, err
	}

	// Use the Actor (actorId) of ActivityPub as the data source of the channel.
	// and set the type to Service.
	// https://www.w3.org/TR/activitystreams-vocabulary/#actor-types
	create, err := actor.NewActorClient(cli.Conn).Create(ctx, &actor.CreateRequest{
		PreferredUsername: in.PreferredUsername,
		PublicKey:         rsa.PublicKey,
		ActorType:         "Service",
	})
	if err != nil {
		return nil, err
	}

	actorId, err := strconv.Atoi(create.ActorId)
	if err != nil {
		return nil, err
	}

	if err := NewChannels(uint(actorId), parse.ActorId, rsa.PrivateKey).CreateChannel(); err != nil {
		return nil, err
	}
	return &pb.CreateChannelResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) GetChannels(ctx context.Context, in *empty.Empty) (*pb.GetChannelsResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	channels, err := NewChannelsCreatorId(parse.ActorId).GetChannels()
	if err != nil {
		return nil, err
	}

	// TODO - When the Actor data is fetched, the Actor data needs to be added to the Channel data.
	// In large scale data, such an operation can cause performance problems.
	// So you need to add concurrent design for optimization.
	var data []*pb.ChannelData
	for _, d := range channels {
		var cd pb.ChannelData
		_ := clientv1.New(ctx,
			microsvc.NewGRPCAddress("actor").Get())
		if err != nil {
			return nil, err
		}
		defer client.Close()

		as, err := actor.NewActorClient(client.Conn).Get(ctx, &actor.GetRequest{
			ActorId: strconv.Itoa(int(d.ActorId)),
		})
		if err != nil {
			return nil, err
		}
		cd.Channel = as.Actor
		cd.ChannelId = strconv.Itoa(int(d.ID))

		data = append(data, &cd)
	}
	return &pb.GetChannelsResponse{
		Code:     "200",
		Channels: data,
	}, nil
}

func (s *server) DeleteChannel(ctx context.Context, in *pb.DeleteChannelRequest) (*pb.DeleteChannelResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	cid, err := strconv.Atoi(in.ChannelId)
	if err != nil {
		return nil, err
	}

	if err := NewChannelsDelete(uint(cid), parse.ActorId).DeleteChannel(); err != nil {
		return nil, err
	}
	return &pb.DeleteChannelResponse{
		Code:   "200",
		Status: "ok",
	}, nil
}

func (s *server) DeleteChannels(ctx context.Context, in *pb.DeleteChannelsRequest) (*pb.DeleteChannelsResponse, error) {
	// TODO - Implement deleting all created channels.
	return &pb.DeleteChannelsResponse{}, nil
}
