/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

// Microservice Authorization Certification Package.

package microsvc

import (
	"context"
	"github.com/hvxahv/hvx/auth"
	"github.com/hvxahv/hvx/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc/metadata"
	"strconv"
	"strings"
)

type Userdata struct {
	AccountId uint
	ActorId   uint
	DeviceID  uint
	Username  string
	Mail      string
}

func GetUserdataByAuthorizationToken(ctx context.Context) (*Userdata, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	bearer := md.Get("authorization")
	if len(bearer) != 1 {
		return nil, errors.New(errors.ErrTokenUnauthorized)
	}

	var (
		token  = strings.Split(bearer[0], "Bearer ")[1]
		secret = viper.GetString("authentication.token.secret")
	)

	parse, err := auth.NewParseJWTToken(token, secret).JWTTokenParse()
	if err != nil {
		return nil, err
	}

	accountId, err := strconv.Atoi(parse.AccountId)
	if err != nil {
		return nil, err
	}
	actor, err := strconv.Atoi(parse.ActorId)
	if err != nil {
		return nil, err
	}
	deviceId, err := strconv.Atoi(parse.DeviceID)
	if err != nil {
		return nil, err
	}

	d := &Userdata{
		AccountId: uint(accountId),
		ActorId:   uint(actor),
		DeviceID:  uint(deviceId),
		Username:  parse.Username,
		Mail:      parse.Mail,
	}
	return d, nil
}
