package transport

import (
	"context"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/metadata"
)

type CustomerTokenAuth struct{}

func (c CustomerTokenAuth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	md, _ := metadata.FromIncomingContext(ctx)
	bearer := md.Get("authorization")
	if len(bearer) != 1 {
		return nil, errors.New("UNAUTHORIZED")
	}
	token := strings.TrimPrefix(bearer[0], "Bearer ")
	return map[string]string{
		"authorization": token,
	}, nil
}
func (c CustomerTokenAuth) RequireTransportSecurity() bool {
	return false
}

type PerRPCCredentials interface {
	GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error)
	RequireTransportSecurity() bool
}
