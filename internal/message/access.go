package message

import (
	"github.com/hvxahv/hvx/pkg/cockroach"
	"github.com/hvxahv/hvx/pkg/matrix"
	"github.com/matrix-org/gomatrix"
	"golang.org/x/net/context"
	"strconv"
)

type Register struct {
	ExampleCredential string `json:"example_credential"`
	Session           string `json:"session"`
	Type              string `json:"type"`
}

func (m *message) MessageAccessRegister(ctx context.Context, in *pb.MessageAccessRegisterRequest) (*pb.MessageAccessRegisterResponse, error) {
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
	return &pb.MessageAccessRegisterResponse{Code: "200", Reply: "ok"}, nil
}

func (m *message) MessageAccessLogin(ctx context.Context, in *pb.MessageAccessLoginRequest) (*pb.MessageAccessLoginResponse, error) {
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

	return &pb.MessageAccessLoginResponse{Code: "200", Reply: "ok"}, nil
}

func (m *message) MessageAccessDelete(ctx context.Context, in *pb.MessageAccessDeleteRequest) (*pb.MessageAccessDeleteResponse, error) {
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	db := cockroach.GetDB()
	if err := db.Debug().Table("matrices").Where("account_id = ?", uint(id)).Unscoped().Delete(&Matrices{}).Error; err != nil {
		return nil, err
	}
	return &pb.MessageAccessDeleteResponse{Code: "200", Reply: "ok"}, nil
}
