package clientv1

import (
	"github.com/hvxahv/hvx/errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc/metadata"
	"strings"
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
