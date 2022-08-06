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
	"strings"
)

func GetUserdataByAuthorizationToken(ctx context.Context) (*auth.Claims, error) {
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

	return parse, nil
}
