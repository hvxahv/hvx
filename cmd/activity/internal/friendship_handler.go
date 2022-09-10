package internal

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"golang.org/x/net/context"
)

func (s *server) GetFollowers(ctx context.Context, in *pb.GetFollowersRequest) (*pb.GetFollowersResponse, error) {

	return &pb.GetFollowersResponse{}, nil
}

func (s *server) GetFollowings(ctx context.Context, in *pb.GetFollowingsRequest) (*pb.GetFollowingsResponse, error) {

	return &pb.GetFollowingsResponse{}, nil
}

func (s *server) GetFriends(ctx context.Context, in *empty.Empty) (*pb.GetFriendsResponse, error) {

	return &pb.GetFriendsResponse{}, nil
}
