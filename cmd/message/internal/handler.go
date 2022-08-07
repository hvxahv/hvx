package internal

import (
	"context"
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/message"
)

func (s *server) AccessRegister(ctx context.Context, in *pb.AccessRegisterRequest) (*pb.AccessRegisterResponse, error) {
	return &pb.AccessRegisterResponse{}, nil
}

func (s *server) AccessLogin(ctx context.Context, in *pb.AccessLoginRequest) (*pb.AccessLoginResponse, error) {
	return &pb.AccessLoginResponse{}, nil
}

func (s *server) AccessDelete(ctx context.Context, in *pb.AccessDeleteRequest) (*pb.AccessDeleteResponse, error) {
	return &pb.AccessDeleteResponse{}, nil
}
