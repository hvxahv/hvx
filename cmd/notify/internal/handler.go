package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/notify"
	"golang.org/x/net/context"
)

func (s *server) Subscribe(ctx context.Context, req *pb.SubscribeRequest) (*pb.SubscribeResponse, error) {
	return &pb.SubscribeResponse{}, nil
}
