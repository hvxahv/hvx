/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package microsvc

import (
	"strconv"

	"github.com/hvxahv/hvx/gateway/identity"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
)

// GetUsernameByTokenWithContext To get the username from TOKEN,
// use the metadata in the GRPC service to get the TOKEN in the
// GRPC GATEWAY HTTP request header and parse it.
func GetUsernameByTokenWithContext(ctx context.Context) (string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	bearer := md.Get("authorization")
	if len(bearer) > 1 {
		return "", errors.New("UNAUTHORIZED")
	}

	pares, err := identity.ParseToken(bearer[0])
	if err != nil {
		return "", errors.New("UNAUTHORIZED")
	}
	return pares.Username, nil
}

func GetAccountIDWithContext(ctx context.Context) (uint, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	bearer := md.Get("authorization")
	if len(bearer) != 1 {
		return 0, errors.New("UNAUTHORIZED")
	}
	pares, err := identity.ParseToken(bearer[1])
	if err != nil {
		return 0, errors.New("UNAUTHORIZED")
	}
	aid, err := strconv.Atoi(pares.ID)
	if err != nil {
		return 0, err
	}
	return uint(aid), nil
}

func GetActorIdByTokenWithContext(ctx context.Context) (uint, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	bearer := md.Get("authorization")
	if len(bearer) > 1 {
		return 0, errors.New("UNAUTHORIZED")
	}

	pares, err := identity.ParseToken(bearer[0])
	if err != nil {
		return 0, errors.New("UNAUTHORIZED")
	}
	id, err := strconv.Atoi(pares.ActorId)
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}
