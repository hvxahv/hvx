package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/channel"
	"golang.org/x/net/context"
)

// Broadcast ...

func (s *server) CreateBroadcast(ctx context.Context, in *pb.CreateBroadcastRequest) (*pb.CreateBroadcastResponse, error) {
	return &pb.CreateBroadcastResponse{}, nil
}

func (s *server) GetBroadcasts(ctx context.Context, in *pb.GetBroadcastsRequest) (*pb.GetBroadcastsResponse, error) {
	return &pb.GetBroadcastsResponse{}, nil

}

func (s *server) DeleteBroadcast(ctx context.Context, in *pb.DeleteBroadcastRequest) (*pb.DeleteBroadcastResponse, error) {
	return &pb.DeleteBroadcastResponse{}, nil
}
