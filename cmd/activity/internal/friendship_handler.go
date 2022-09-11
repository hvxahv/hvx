package internal

import (
	"fmt"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/cmd/activity/internal/friendship"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) GetFollowers(ctx context.Context, in *emptypb.Empty) (*pb.GetFollowersResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	follows, err := friendship.NewFollows(parse.ActorId, friendship.Follower).Get()
	if err != nil {
		return nil, err
	}
	fmt.Println(follows)
	var f []*actor.ActorData

	return &pb.GetFollowersResponse{
		Code:      "200",
		Followers: f,
	}, nil
}

func (s *server) GetFollowings(ctx context.Context, in *emptypb.Empty) (*pb.GetFollowingsResponse, error) {

	return &pb.GetFollowingsResponse{}, nil
}

func (s *server) GetFriends(ctx context.Context, in *emptypb.Empty) (*pb.GetFriendsResponse, error) {

	return &pb.GetFriendsResponse{}, nil
}
