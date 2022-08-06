/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/auth"
	"github.com/hvxahv/hvx/auth"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"time"
)

func (s *server) Authorization(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	// Implement the authorization logic here.
	// You can use the `in` parameter to get the username and password.
	// account server.

	v, err := NewAuthorization(ctx).Authorization(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	// Implement the create devices.
	device, err := NewAuthorization(ctx).AddDevice(v.Id, in.UserAgent)
	if err != nil {
		return nil, err
	}

	// Implement Generate TOKEN...
	var (
		issuer = viper.GetString("domain")
		expir  = time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour
		secret = viper.GetString("authentication.token.secret")
	)
	g, err := auth.NewClaims(
		auth.NewUserdata(v.Id, v.ActorId, device.DeviceId, v.Username, v.Mail),
		auth.NewRegisteredClaims(issuer, device.DeviceId, v.Id, expir),
	).JWTTokenGenerator(secret)
	if err != nil {
		return nil, err
	}

	return &pb.AuthorizationResponse{
		Code:               "200",
		Status:             "ok",
		Id:                 v.Id,
		AuthorizationToken: g,
		ActorId:            v.ActorId,
		Mail:               v.Mail,
		DeviceId:           device.DeviceId,
	}, nil
}
