/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

// Microservice Authorization Certification Package.

package v

import (
	"github.com/hvxahv/hvx/pkg/identity"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"strings"
)

// GetUsernameByTokenWithContext To get the username from TOKEN,
// use the metadata in the GRPC service to get the TOKEN in the
// GRPC GATEWAY HTTP request header and parse it.
func GetUsernameByTokenWithContext(ctx context.Context) (string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	bearer := md.Get("authorization")
	if len(bearer) != 1 {
		return "", errors.New("UNAUTHORIZED")
	}
	token := strings.TrimPrefix(bearer[0], "Bearer ")
	pares, err := identity.ParseToken(token)
	if err != nil {
		return "", errors.New("UNAUTHORIZED")
	}
	return pares.Username, nil
}
