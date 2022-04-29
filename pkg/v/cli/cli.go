package cli

import (
	"google.golang.org/grpc"
)

// conn, err := grpc.DialContext(ctx, "grpc://0.0.0.0:0000", grpc.WithInsecure())
//	if err != nil {
//		return err
//	}
//	defer conn.Close()
//	client := NewPublicClient(conn)

type Client interface {
	Account
	Activity
	Article
	Channel
	Device
	Message
	Public
	Saved
}

type client struct {
	Account
	Activity
	Article
	Channel
	Device
	Message
	Public
	Saved
}

func NewHvxClient(conn *grpc.ClientConn) Client {
	return &client{
		Account:  NewAccount(conn),
		Activity: NewActivity(conn),
		Article:  NewArticle(conn),
		Channel:  NewChannel(conn),
		Device:   NewDevice(conn),
		Message:  NewMessage(conn),
		Public:   NewPublic(conn),
		Saved:    NewSaved(conn),
	}
}
