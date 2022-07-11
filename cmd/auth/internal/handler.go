/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"time"

	"github.com/hvxahv/hvx/APIs/grpc/v1alpha1/account"
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/auth"
	"github.com/hvxahv/hvx/APIs/grpc/v1alpha1/device"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/identity/jwt"
	"github.com/hvxahv/hvx/microsvc"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

func (s *server) Authorization(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {

	// TODO: implement the authorization logic here.
	// You can use the `in` parameter to get the username and password.
	v, err := clientv1.New(ctx,
		cfg.SetEndpoints(microsvc.GetGRPCServiceAddress("account")),
		cfg.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer v.Close()
	a, err := v.Verify(ctx, &account.VerifyRequest{
		Username: in.Username,
		Password: in.Password,
	})
	if err != nil {
		return nil, err
	}

	// TODO: implement the create devices.
	cli, err := clientv1.New(ctx,
		cfg.SetEndpoints(microsvc.GetGRPCServiceAddress("device")),
		cfg.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	device, err := cli.Create(ctx, &device.CreateRequest{
		AccountId: conv.UintToString(a.ID),
		Ua:        in.Ua,
	})
	if err != nil {
		return nil, err
	}

	// TODO: implement Generate TOKEN...
	expired := time.Duration(viper.GetInt("authentication.token.expired"))
	g := jwt.NewClaims(a.Mail, a.Id, a.ActorId, a.Username, device.DeviceId, expired)
	k, err := g.JWTTokenGenerator(viper.GetString("authentication.token.signed"))
	if err != nil {
		return nil, err
	}

	return &pb.AuthorizationResponse{
		Code:               "200",
		Status:             "ok",
		Id:                 a.Id,
		AuthorizationToken: k,
		ActorId:            a.ActorId,
		Mail:               a.Mail,
		DeviceId:           device.DeviceId,
	}, nil
}
