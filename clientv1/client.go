package clientv1

import (
	"context"
	"github.com/hvxahv/hvx/clientv1/actor"
	"github.com/hvxahv/hvx/clientv1/auth"
	config "github.com/hvxahv/hvx/clientv1/cfg"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Example...
//cfg, err := New(context.Background(),
//	SetEndpoints("localhost:50051"),
//	SetDialOptionsWithToken(),
//	SetDialTimeout(10*time.Second),
//)
//if err != nil {
//	return
//}
//defer cfg.Close()

type Client struct {
	Context context.Context
	Conn    *grpc.ClientConn

	actor.Actor
	auth.Auth
}

func New(ctx context.Context, endpoints []string, cfg ...config.Option) (*Client, error) {
	c := &config.Config{
		Endpoints: endpoints,
	}

	for _, opt := range cfg {
		opt(c)
	}

	c.DialOptions = append(c.DialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if len(c.Endpoints) < 1 {
		return nil, errors.New("At least one address is required.")
	}
	addr := c.Endpoints[0]
	conn, err := grpc.DialContext(ctx, addr, c.DialOptions...)
	if err != nil {
		return nil, err
	}

	client := &Client{
		Conn:    conn,
		Context: ctx,
	}

	client.Actor = actor.NewActor(client)
	client.Auth = auth.NewAuth(client)

	return client, nil
}

func (c *Client) Close() error {
	return c.Conn.Close()
}
