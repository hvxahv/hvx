package clientv1

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
)

// Example...
//cli, err := New(context.Background(),
//	SetEndpoints("localhost:50051"),
//	SetDialOptionsWithToken(),
//	SetDialTimeout(10*time.Second),
//)
//if err != nil {
//	return
//}
//defer cli.Close()

type Client struct {
	Account
	Activity
	Article
	Channel
	Device
	Message
	Public
	Saved

	ctx  context.Context
	conn *grpc.ClientConn
}

func New(ctx context.Context, cfg ...Option) (*Client, error) {
	c := &Config{}

	for _, opt := range cfg {
		opt(c)
	}

	addr := fmt.Sprintf("grpc://%s", c.Endpoints[0])
	conn, err := grpc.DialContext(ctx, addr, c.DialOptions...)
	if err != nil {
		return nil, err
	}

	client := &Client{
		conn: conn,
		ctx:  ctx,
	}

	client.Account = NewAccount(client)
	client.Activity = NewActivity(client)
	client.Article = NewArticle(client)
	client.Channel = NewChannel(client)
	client.Device = NewDevice(client)
	client.Message = NewMessage(client)
	client.Public = NewPublic(client)
	client.Saved = NewSaved(client)

	return client, nil
}

func (c *Client) Close() error {
	return c.conn.Close()
}
