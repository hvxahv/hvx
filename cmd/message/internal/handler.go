package internal

import (
	"context"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/message"
)

// TODO - MESSAGE API WILL BE REDESIGNED.

func (s *server) AccessRegister(ctx context.Context, in *pb.AccessRegisterRequest) (*pb.AccessRegisterResponse, error) {
	d, err := NewRegister().Register(in.Username, in.Password)
	if err != nil {
		return nil, err
	}
	return &pb.AccessRegisterResponse{
		Code:        "200",
		AccessToken: d.AccessToken,
		UserId:      d.UserId,
	}, nil
}

func (s *server) AccessLogin(ctx context.Context, in *pb.AccessLoginRequest) (*pb.AccessLoginResponse, error) {
	return &pb.AccessLoginResponse{}, nil
}

func (s *server) AccessDelete(ctx context.Context, in *pb.AccessDeleteRequest) (*pb.AccessDeleteResponse, error) {
	return &pb.AccessDeleteResponse{}, nil
}

// TODO - EDIT PASSWORD...
