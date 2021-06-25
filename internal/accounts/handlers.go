package accounts

import (
	"fmt"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
)

// NewAccounts ...
func (s *server) NewAccounts(ctx context.Context, in *pb.NewAccountsData) (*pb.AccountsReply, error) {
	fmt.Println(in)
	a := NewAccounts(in.Username, in.Password, in.Avatar, in.Name, in.Email, in.Private)
	if err := a.New(); err != nil {
		return nil, err
	}

	return &pb.AccountsReply{Reply: "ok"}, nil
}

// UpdateAccounts ...
func (s *server) UpdateAccounts(ctx context.Context, in *pb.UpdateAccountsData) (*pb.AccountsReply, error) {
	// Accounts
	a := NewUpdateAcct()
	a.Username = in.Username
	a.Password = in.Password
	a.Avatar = in.Avatar
	a.Bio = in.Bio
	a.Name = in.Name
	a.EMail = in.Email
	a.Phone = in.Phone
	a.Telegram = in.Telegram
	a.Private = in.Private

	if err := a.Update(); err != nil {
		return nil, err
	}

	return &pb.AccountsReply{Reply: "ok"}, nil
}

// QueryAccounts ...
func (s *server) QueryAccounts(ctx context.Context, in *pb.AccountsName) (*pb.AccountsData, error) {
	r := NewQueryAcctByName(in.Username)
	a, err := r.Query()
	if err != nil {
		return nil, err
	}

	return &pb.AccountsData{
		Uuid:      a.Uuid,
		Username:  a.Username,
		Avatar:    a.Avatar,
		Bio:       a.Bio,
		Name:      a.Name,
		Email:     a.EMail,
		Phone:     a.Phone,
		Telegram:  a.Telegram,
		Private:   a.Private,
		PublicKey: a.PublicKey,
	}, nil
}

// DeleteAccounts ...
func (s *server) DeleteAccounts(ctx context.Context, in *pb.AccountsName) (*pb.AccountsReply, error) {
	r := NewDelAcctByName(in.Username)
	err := r.Delete()
	if err != nil {
		return nil, err
	}
	return &pb.AccountsReply{Reply: "ok"}, nil
}

// LoginAccounts ...
func (s *server) LoginAccounts(ctx context.Context, in *pb.AccountsLogin) (*pb.AccountsLoginReply, error) {
	r := NewAccountLogin(in.Username, in.Password)
	token, err := r.Login()
	if err != nil {
		return nil, err
	}
	return &pb.AccountsLoginReply{
		Username: in.Username,
		Token:    token,
	}, nil
}
