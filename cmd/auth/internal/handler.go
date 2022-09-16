/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"github.com/hvxahv/hvx/errors"
	"strconv"
	"time"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/auth"
	"github.com/hvxahv/hvx/auth"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
)

// Authorization authenticate via the account service.
// add the login device information to the device service after successful authentication
// Generate an AccessToken from the authentication result and return it to the client.
func (s *server) Authorization(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {
	v, err := NewAuthorization(ctx).Authorization(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	device, err := NewAuthorization(ctx).AddDevice(v.AccountId, in.UserAgent)
	if err != nil {
		return nil, err
	}

	var (
		issuer = viper.GetString("domain")
		expir  = time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour
		secret = viper.GetString("authentication.token.secret")
	)
	g, err := auth.NewClaims(
		auth.NewUserdata(
			strconv.Itoa(int(v.AccountId)),
			strconv.Itoa(int(v.ActorId)),
			strconv.Itoa(int(device.DeviceId)),
			v.Username,
			v.Mail,
		),
		auth.NewRegisteredClaims(issuer, strconv.Itoa(int(device.DeviceId)), strconv.Itoa(int(v.AccountId)), expir),
	).JWTTokenGenerator(secret)
	if err != nil {
		errors.Throw("cannot generate token error during authentication.", err)
		return nil, err
	}

	return &pb.AuthorizationResponse{
		Code:               "200",
		Status:             "ok",
		AccountId:          v.AccountId,
		AuthorizationToken: g,
		ActorId:            v.ActorId,
		Mail:               v.Mail,
		DeviceId:           device.DeviceId,
	}, nil
}

func (s *server) SetPublicKey(ctx context.Context, in *pb.SetPublicKeyRequest) (*pb.SetPublicKeyResponse, error) {
	if err := NewAuthorization(ctx).SetPublicKey(uint(in.GetAccountId()), in.GetPublicKey()); err != nil {
		return nil, err
	}
	return &pb.SetPublicKeyResponse{Code: "200", Status: "ok"}, nil
}
