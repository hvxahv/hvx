package clientv1

import (
	"context"

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

func New(ctx context.Context, endpoint string, cfg ...Option) (*Client, error) {
	c := &Config{
		Endpoint: endpoint,
	}

	for _, opt := range cfg {
		opt(c)
	}

	c.DialOptions = append(c.DialOptions, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.DialContext(ctx, c.Endpoint, c.DialOptions...)
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
