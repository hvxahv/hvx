package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/activity"
	"golang.org/x/net/context"
)

func (s *server) CreateOutbox(ctx context.Context, in *pb.CreateOutboxRequest) (*pb.CreateOutboxResponse, error) {

	return &pb.CreateOutboxResponse{}, nil
}

func (s *server) GetOutbox(ctx context.Context, in *pb.GetOutboxRequest) (*pb.GetOutboxResponse, error) {

	return &pb.GetOutboxResponse{}, nil
}

func (s *server) GetOutboxes(ctx context.Context, in *pb.GetOutboxesRequest) (*pb.GetOutboxesResponse, error) {

	return &pb.GetOutboxesResponse{}, nil
}
