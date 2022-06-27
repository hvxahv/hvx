/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/activity"
	"golang.org/x/net/context"
)

func (s *server) Inbox(ctx context.Context, in *pb.InboxRequest) (*pb.InboxResponse, error) {

	return &pb.InboxResponse{}, nil
}

func (s *server) GetInbox(ctx context.Context, in *pb.GetInboxRequest) (*pb.GetInboxResponse, error) {

	return &pb.GetInboxResponse{}, nil
}

func (s *server) GetInboxes(ctx context.Context, in *pb.GetInboxesRequest) (*pb.GetInboxesResponse, error) {

	return &pb.GetInboxesResponse{}, nil
}

func (s *server) DeleteInbox(ctx context.Context, in *pb.DeleteInboxRequest) (*pb.DeleteInboxResponse, error) {

	return &pb.DeleteInboxResponse{}, nil
}

func (s *server) CreateOutbox(ctx context.Context, in *pb.CreateOutboxRequest) (*pb.CreateOutboxResponse, error) {

	return &pb.CreateOutboxResponse{}, nil
}

func (s *server) GetOutbox(ctx context.Context, in *pb.GetOutboxRequest) (*pb.GetOutboxResponse, error) {

	return &pb.GetOutboxResponse{}, nil
}

func (s *server) GetOutboxes(ctx context.Context, in *pb.GetOutboxesRequest) (*pb.GetOutboxesResponse, error) {

	return &pb.GetOutboxesResponse{}, nil
}

func (s *server) GetFollowers(ctx context.Context, in *pb.GetFollowersRequest) (*pb.GetFollowersResponse, error) {

	return &pb.GetFollowersResponse{}, nil
}

func (s *server) GetFollowings(ctx context.Context, in *pb.GetFollowingsRequest) (*pb.GetFollowingsResponse, error) {

	return &pb.GetFollowingsResponse{}, nil
}

func (s *server) Follow(ctx context.Context, in *pb.FollowRequest) (*pb.FollowResponse, error) {

	return &pb.FollowResponse{}, nil
}

func (s *server) UnFollow(ctx context.Context, in *pb.UnFollowRequest) (*pb.UnFollowResponse, error) {

	return &pb.UnFollowResponse{}, nil
}

func (s *server) GetFriends(ctx context.Context, in *empty.Empty) (*pb.GetFriendsResponse, error) {

	return &pb.GetFriendsResponse{}, nil
}
