package accounts

import (
	pb "github.com/disism/hvxahv/api/hvxahv/v1alpha1"
	"golang.org/x/net/context"
)

// NewAccount Implementation of the method of creating an account.
func (s *server) NewAccount(ctx context.Context, in *pb.NewAccountData) (*pb.AccountsReply, error) {
	a, _ := NewAccounts(in.Username, in.Password, in.Mail)
	code, messages := a.New()
	return &pb.AccountsReply{Code: code, Message: messages}, nil
}

// LoginAccount The implementation of the login account method returns the user name and token,
// and returns a specific error if there is an error.
func (s *server) LoginAccount(ctx context.Context, in *pb.LoginData) (*pb.LoginReply, error) {
	r := NewAccountLogin(in.Mail, in.Password)
	u, token, err := r.Login()
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{
		Username: u,
		Token:    token,
	}, nil
}


//// UpdateAccounts Implementation of the method to update the account.
//func (s *server) UpdateAccount(ctx context.Context, in *pb.AccountData) (*pb.AccountsReply, error) {
//	a := AccountData{
//		Username:   in.Username,
//		Password:   in.Password,
//		Avatar:     in.Avatar,
//		Bio:        in.Bio,
//		Name:       in.Name,
//		Mail:       in.Mail,
//		Phone:      in.Phone,
//		Private:    in.Private,
//		PrivateKey: in.PrivateKey,
//		PublicKey:  in.PublicKey,
//	}
//
//	if err := a.Update(); err != nil {
//		return nil, err
//	}
//
//	return &pb.AccountsReply{Code: 200, Message: "ok"}, nil
//}

//// QueryAccounts Implementation of the method of querying the account.
//func (s *server) QueryAccounts(ctx context.Context, in *pb.AccountsName) (*pb.AccountsData, error) {
//	r := NewQueryAcctByName(in.Username)
//	a, err := r.Query()
//	if err != nil {
//		return nil, err
//	}
//
//	return &pb.AccountsData{
//		Uuid:      a.Uuid,
//		Username:  a.Username,
//		Avatar:    a.Avatar,
//		Bio:       a.Bio,
//		Name:      a.Name,
//		Email:     a.EMail,
//		Phone:     a.Phone,
//		Telegram:  a.Telegram,
//		Private:   a.Private,
//		PublicKey: a.PublicKey,
//	}, nil
//}
//
//// DeleteAccounts Implementation of the delete account method.
//func (s *server) DeleteAccounts(ctx context.Context, in *pb.AccountsName) (*pb.AccountsReply, error) {
//	r := NewDelAcctByName(in.Username)
//	err := r.Delete()
//	if err != nil {
//		return nil, err
//	}
//	return &pb.AccountsReply{Code: 200, Message: "ok"}, nil
//}
//
