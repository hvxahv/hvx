package clientv1

import (
	"context"
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
}

func New(ctx context.Context, endpoints []string, cfg ...Option) (*Client, error) {
	c := &Config{
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

	return client, nil
}

func (c *Client) Close() error {
	return c.Conn.Close()
}
