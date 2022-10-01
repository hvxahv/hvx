package internal

import (
	"context"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/matrix"
	"github.com/hvxahv/hvx/microsvc"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/message"
)

// TODO - MESSAGE API WILL BE REDESIGNED.

func (s *server) AccessRegister(ctx context.Context, in *pb.AccessRegisterRequest) (*pb.AccessRegisterResponse, error) {
	// Go to the account server to verify the username and password.
	v, err := clientv1.New(ctx, microsvc.AccountServiceName).Verify(in.GetUsername(), in.GetPassword())
	if err != nil {
		return nil, errors.New(errors.ErrAccountVerification)
	}

	instance, err := clientv1.New(context.Background(), microsvc.PublicServiceName).GetInstance()
	if err != nil {
		return nil, err
	}

	dummy, err := matrix.New(instance.MatrixAPI).RegisterDummy(in.GetUsername(), in.GetPassword())
	if err != nil {
		return nil, err
	}

	if err := NewMatrices(
		uint(v.ActorId),
		dummy.DeviceID.String(),
		instance.MatrixAPI,
		dummy.UserID.String(),
	).Create(); err != nil {
		return nil, err
	}

	return &pb.AccessRegisterResponse{
		Code:        "200",
		AccessToken: dummy.AccessToken,
		UserId:      dummy.UserID.String(),
	}, nil
}

func (s *server) AccessLogin(ctx context.Context, in *pb.AccessLoginRequest) (*pb.AccessLoginResponse, error) {
	return &pb.AccessLoginResponse{}, nil
}

func (s *server) AccessDelete(ctx context.Context, in *pb.AccessDeleteRequest) (*pb.AccessDeleteResponse, error) {
	return &pb.AccessDeleteResponse{}, nil
}

// TODO - EDIT PASSWORD...
