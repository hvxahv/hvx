package accounts

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"golang.org/x/net/context"
)

// New Implementation of the method of creating an account.
func (s *server) New(ctx context.Context, in *pb.NewAccountData) (*pb.AccountsReply, error) {
	a, _ := NewAccounts(in.Username, in.Password, in.Mail)
	code, messages := a.New()
	return &pb.AccountsReply{Code: code, Message: messages}, nil
}

// Login The implementation of the login account method returns the user name and token,
// and returns a specific error if there is an error.
func (s *server) Login(ctx context.Context, in *pb.AuthData) (*pb.AuthReply, error) {
	r := NewAccountAuth(in.Mail, in.Password)
	u, err := r.Login()
	if err != nil {
		return nil, err
	}
	return &pb.AuthReply{
		Username: u,
	}, nil
}

// Update Implementation of the method to update the account.
func (s *server) Update(ctx context.Context, in *pb.AccountData) (*pb.AccountsReply, error) {
	a := Accounts{
		Username:   in.Username,
		Password:   in.Password,
		Avatar:     in.Avatar,
		Bio:        in.Bio,
		Name:       in.Name,
		Mail:       in.Mail,
		Phone:      in.Phone,
		IsPrivate:  in.IsPrivate,
		PrivateKey: in.PrivateKey,
		PublicKey:  in.PublicKey,
	}

	if err := a.Update(); err != nil {
		return &pb.AccountsReply{Code: 500, Message: "update error!"}, nil
	}

	return &pb.AccountsReply{Code: 200, Message: "ok!"}, nil
}

// QueryByName Implementation of the method of querying the account.
func (s *server) QueryByName(ctx context.Context, in *pb.NewAccountByName) (*pb.AccountData, error) {
	r := NewAccountByName(in.Username)
	a, err := r.QueryByName()
	if err != nil {
		return nil, err
	}

	return &pb.AccountData{
		Id:        uint64(a.ID),
		Username:  a.Username,
		Mail:      a.Mail,
		Avatar:    a.Avatar,
		Bio:       a.Bio,
		Name:      a.Name,
		Phone:     a.Phone,
		IsPrivate: a.IsPrivate,
		Follower:  int32(a.Follower),
		Following: int32(a.Following),
		Friend:    int32(a.Friend),
		PublicKey: a.PublicKey,
	}, nil
}

// QueryByID Implementation of the method of querying the account.
func (s *server) QueryByID(ctx context.Context, in *pb.NewAccountByID) (*pb.AccountData, error) {
	r := NewAccountByID(uint(in.Id))
	a, err := r.QueryByID()
	if err != nil {
		return nil, err
	}

	return &pb.AccountData{
		Id:        uint64(a.ID),
		Username:  a.Username,
		Mail:      a.Mail,
		Avatar:    a.Avatar,
		Bio:       a.Bio,
		Name:      a.Name,
		Phone:     a.Phone,
		IsPrivate: a.IsPrivate,
		Follower:  int32(a.Follower),
		Following: int32(a.Following),
		Friend:    int32(a.Friend),
		PublicKey: a.PublicKey,
	}, nil
}

// Delete Implementation of delete account method.
func (s *server) Delete(ctx context.Context, in *pb.AuthData) (*pb.AccountsReply, error) {
	fmt.Println(in.Mail, in.Password)
	//r := NewAccountAuth(in.Mail, in.Password)
	//err := r.Login()
	//if err != nil {
	//	return &pb.AccountsReply{Code: 500, Message: "delete error!"}, nil
	//}
	return &pb.AccountsReply{Code: 200, Message: "ok!"}, nil
}

// NewFollow Implementation of the method of querying the account.
//func (s *server) NewFollow(ctx context.Context, in *pb.FollowersData) (*pb.AccountsReply, error) {
//	nf := NewFos(in.Follower, in.Following)
//	if err := nf.New(); err != nil {
//		log.Println(err)
//		return &pb.AccountsReply{
//			Code:    202,
//			Message: "Follow failed!",
//		}, nil
//	}
//	return &pb.AccountsReply{
//		Code:    200,
//		Message: "Followed!",
//	}, nil
//}
