package accounts

import (
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
)

// NewAccounts Implementation of the method of creating an account.
func (s *server) NewAccounts(ctx context.Context, in *pb.NewAccountsData) (*pb.AccountsReply, error) {
	a := NewAccounts(in.Username, in.Password, in.Avatar, in.Name, in.Email, in.Private)
	code, err := a.New()
	return &pb.AccountsReply{Code: code, Message: err.Error()}, nil
}

// UpdateAccounts Implementation of the method to update the account.
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

	return &pb.AccountsReply{Code: 200, Message: "ok"}, nil
}

// QueryAccounts Implementation of the method of querying the account.
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

// DeleteAccounts Implementation of the delete account method.
func (s *server) DeleteAccounts(ctx context.Context, in *pb.AccountsName) (*pb.AccountsReply, error) {
	r := NewDelAcctByName(in.Username)
	err := r.Delete()
	if err != nil {
		return nil, err
	}
	return &pb.AccountsReply{Code: 200, Message: "ok"}, nil
}

// LoginAccounts The implementation of the login account method returns the user name and token,
// and returns a specific error if there is an error.
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
