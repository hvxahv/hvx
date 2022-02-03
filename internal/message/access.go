package message

import (
	"context"
	pb "github.com/hvxahv/hvxahv/api/message/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/matrix"
	"github.com/matrix-org/gomatrix"
	"strconv"
)

type Register struct {
	ExampleCredential string `json:"example_credential"`
	Session           string `json:"session"`
	Type              string `json:"type"`
}

func (m *message) MessageAccessRegister(ctx context.Context, in *pb.NewMessageAccess) (*pb.MessageReply, error) {
	cli, err := matrix.NewClient("", "")
	if err != nil {
		return nil, err
	}

	res, _, err := cli.Register(&gomatrix.ReqRegister{
		Username:  in.Username,
		BindEmail: false,
		Password:  in.Password,
		Auth: &Register{
			ExampleCredential: "",
			Session:           "",
			Type:              "m.login.dummy",
		},
	})
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}

	c := NewMatrixAccesses(uint(id), res.AccessToken, res.HomeServer, res.UserID, res.DeviceID)
	if err := c.Create(); err != nil {
		return nil, err
	}
	return &pb.MessageReply{Code: "200", Reply: "ok"}, nil
}

func (m *message) MessageAccessLogin(ctx context.Context, in *pb.NewMessageAccess) (*pb.MessageReply, error) {
	cli, err := matrix.NewClient(in.Username, in.Password)
	if err != nil {
		return nil, err
	}

	r := &gomatrix.ReqLogin{
		Password: in.Password,
		User:     in.Username,
	}

	login, err := cli.Login(r)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}

	if err := NewAccessUpdateToken(uint(id), login.AccessToken).UpdateToken(); err != nil {
		return nil, err
	}

	return nil, nil
}
