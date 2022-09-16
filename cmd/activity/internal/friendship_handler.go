package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cmd/activity/internal/friendship"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GetFollower(ctx context.Context, in *emptypb.Empty) (*pb.FriendshipResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	follows, err := friendship.NewFollows(parse.ActorId, friendship.Follower).Get()
	if err != nil {
		return nil, err
	}

	var f []*actor.ActorData
	for _, i := range follows {
		a, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(i))
		if err != nil {
			return nil, err
		}
		f = append(f, a.Actor)
	}

	return &pb.FriendshipResponse{
		Code:   "200",
		Actors: f,
	}, nil
}

func (s *server) GetFollowing(ctx context.Context, in *emptypb.Empty) (*pb.FriendshipResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	follows, err := friendship.NewFollows(parse.ActorId, friendship.Following).Get()
	if err != nil {
		return nil, err
	}

	var f []*actor.ActorData
	for _, i := range follows {
		a, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(i))
		if err != nil {
			return nil, err
		}
		f = append(f, a.Actor)
	}

	return &pb.FriendshipResponse{
		Code:   "200",
		Actors: f,
	}, nil
}

func (s *server) GetFriend(ctx context.Context, in *emptypb.Empty) (*pb.FriendshipResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	follows, err := friendship.NewFollows(parse.ActorId, friendship.Friend).Get()
	if err != nil {
		return nil, err
	}

	var f []*actor.ActorData
	for _, i := range follows {
		a, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(i))
		if err != nil {
			return nil, err
		}
		f = append(f, a.Actor)
	}

	return &pb.FriendshipResponse{
		Code:   "200",
		Actors: f,
	}, nil
}
