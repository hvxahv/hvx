package clientv1

import (
	"context"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
		conn: conn,
		ctx:  ctx,
	}
	// ...
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
