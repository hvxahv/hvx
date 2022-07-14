package cfg

import (
	"crypto/tls"
	"github.com/hvxahv/hvx/clientv1"
	"google.golang.org/grpc"
	"time"
)

type Option func(o *Config)

type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
	TLS         *tls.Config
	DialOptions []grpc.DialOption
}

func SetDialTimeout(t time.Duration) Option {
	if t < 0 {
		t = time.Second * 3
	}
	return func(c *Config) {
		c.DialTimeout = t
	}
}

func SetTLS(tls *tls.Config) Option {
	if tls == nil {
		return nil
	}
	return func(c *Config) {
		c.TLS = tls
	}
}

func SetDialOptionsWithToken() Option {
	o := &[]grpc.DialOption{
		grpc.WithPerRPCCredentials(clientv1.CustomerTokenAuth{}),
	}
	return func(c *Config) {
		c.DialOptions = *o
	}
}