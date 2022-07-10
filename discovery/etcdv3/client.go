package etcdv3

import (
	"context"
)

const (
	prefix = "/"
)

type Client interface {
	Register()
	Deregister()
}

type client struct {
	ctx context.Context
	// cfg *v3.Client
	// kv v3.KV
	// watcher v3.Watcher
	// leaseID v3.LeaseID
}

type ClientOptions struct {
	Username string
	Password string
}

func NewClient() {

}
