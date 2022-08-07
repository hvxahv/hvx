package internal

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/hvxahv/hvx/rsa"
	"strconv"
	"time"
)

// CreateChannel ...
func (s *server) CreateChannel(ctx context.Context, in *pb.CreateChannelRequest) (*pb.CreateChannelResponse, error) {
	client, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")})
	if err != nil {
		return nil, err
	}
	defer client.Close()

	exist, err := actor.NewActorClient(client.Conn).IsExist(ctx, &actor.IsExistRequest{
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
	rsa, err := rsa.NewRsa(2048).Generate()
	if err != nil {
		return nil, err
	}

	cli, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")},
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	create, err := actor.NewActorClient(cli.Conn).Create(ctx, &actor.CreateRequest{
		PreferredUsername: in.PreferredUsername,
		PublicKey:         rsa.PublicKey,
		ActorType:         "service",
	})
	if err != nil {
		return nil, err
	}
	actorId, err := strconv.Atoi(create.ActorId)
	if err != nil {
		return nil, err
	}

	accountId, err := strconv.Atoi(parse.AccountId)
	if err != nil {
		return nil, err
	}

	if err := NewChannels(uint(actorId), uint(accountId), rsa.Private).CreateChannel(); err != nil {
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
	accountId, err := strconv.Atoi(parse.AccountId)
	if err != nil {
		return nil, err
	}

	channels, err := NewChannelsAccountId(uint(accountId)).GetChannels()
	if err != nil {
		return nil, err
	}

	var res []*actor.ActorData
	for _, d := range channels {
		client, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")})
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

		res = append(res, as.Actor)
	}
	return &pb.GetChannelsResponse{
		Code:     "200",
		Channels: res,
	}, nil
}

func (s *server) DeleteChannel(ctx context.Context, in *pb.DeleteChannelRequest) (*pb.DeleteChannelResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	client, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")})
	if err != nil {
		return nil, err
	}
	defer client.Close()

	d, err := actor.NewActorClient(client.Conn).Delete(ctx, &actor.DeleteRequest{
		Id: in.ActorId,
	})
	if err != nil {
		return nil, err
	}

	actorId, err := strconv.Atoi(in.ActorId)
	if err != nil {
		return nil, err
	}

	accountId, err := strconv.Atoi(parse.AccountId)
	if err != nil {
		return nil, err
	}
	if err := NewChannelsDelete(uint(actorId), uint(accountId)).DeleteChannel(); err != nil {
		return nil, err
	}
	return &pb.DeleteChannelResponse{
		Code:   "200",
		Status: d.Status,
	}, nil
}

func (s *server) DeleteChannels(ctx context.Context, in *pb.DeleteChannelsRequest) (*pb.DeleteChannelsResponse, error) {
	return &pb.DeleteChannelsResponse{}, nil
}
