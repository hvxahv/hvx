package clientv1

import (
	"crypto/tls"
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

func SetEndpoints(eps ...string) Option {
	if len(eps) == 0 {
		return nil
	}
	return func(c *Config) {
		c.Endpoints = eps
	}
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
		grpc.WithPerRPCCredentials(CustomerTokenAuth{}),
	}
	return func(c *Config) {
		c.DialOptions = *o
	}
}
