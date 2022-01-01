package accounts

import (
	"log"

	pb "github.com/hvxahv/hvxahv/api/accounts/v1alpha1"
	"golang.org/x/net/context"
)

func (s *server) Create(ctx context.Context, in *pb.CreateAccountData) (*pb.AccountsReply, error) {
	// Create the Actor first, and then use the returned ActorID to create a unique account of the current instance account system.
	// The username in the account system is unique, and the Actor may have the same username in different instances.
	actor, err := NewActors(in.Username, "", "Person").Create()
	if err != nil {
		log.Println(err)
	}

	if err := NewAccounts(in.Username, in.Mail, in.Password, actor.ID).Create(); err != nil {
		log.Panicln(err)
	}

	return &pb.AccountsReply{Code: "200", Message: "ok"}, nil
}

func (s *server) SignIn(ctx context.Context, in *pb.AuthData) (*pb.SignInReply, error) {
	a := NewAuth(in.Username, in.Password)
	id, mail, err := a.SignIn()
	if err != nil {
		return &pb.SignInReply{Code: "500", AccountID: 0, Mail: ""}, err
	}
	return &pb.SignInReply{Code: "200", AccountID: uint64(id), Mail: mail}, nil
}

func (s *server) Update(ctx context.Context, in *pb.AccountData) (*pb.AccountsReply, error) {
	a := Accounts{
		Username:  in.Username,
		Mail:      in.Mail,
		IsPrivate: in.IsPrivate,
	}

	if err := a.Update(); err != nil {
		return &pb.AccountsReply{Code: "500", Message: "!ok"}, err
	}

	return &pb.AccountsReply{Code: "200", Message: "ok"}, nil
}

func (s *server) GetAccountsByUsername(ctx context.Context, in *pb.AccountUsername) (*pb.AccountData, error) {
	r := NewAccountsUsername(in.Username)
	a, err := r.GetAccountByUsername()
	if err != nil {
		return nil, err
	}

	return &pb.AccountData{
		Id:        uint64(a.ID),
		Username:  a.Username,
		Mail:      a.Mail,
		IsPrivate: a.IsPrivate,
	}, nil
}

func (s *server) FindActorByAccountsUsername(ctx context.Context, in *pb.AccountUsername) (*pb.ActorData, error) {
	a := NewActorsPreferredUsername(in.Username)
	actor, err := a.GetActorByAccountUsername()
	if err != nil {
		return nil, err
	}
	return &pb.ActorData{
		Id:                uint64(actor.ID),
		PreferredUsername: actor.PreferredUsername,
		Domain:            actor.Domain,
		Avatar:            actor.Avatar,
		Name:              actor.Name,
		Summary:           actor.Summary,
		Inbox:             actor.Inbox,
		PublicKey:         actor.PublicKey,
		ActorType:         actor.ActorType,
	}, nil
}

func (s *server) FindActorByID(ctx context.Context, in *pb.ActorID) (*pb.ActorData, error) {
	a := NewActorID(uint(in.ActorID))
	a, err := a.GetByActorID()
	if err != nil {
		return nil, err
	}

	return &pb.ActorData{
		PreferredUsername: a.PreferredUsername,
		Domain:            a.Domain,
		Avatar:            a.Avatar,
		Name:              a.Name,
		Summary:           a.Summary,
		Inbox:             a.Inbox,
		PublicKey:         a.PublicKey,
		ActorType:         a.ActorType,
	}, nil
}

// Delete Implementation of delete account method.
func (s *server) Delete(ctx context.Context, in *pb.DeleteData) (*pb.AccountsReply, error) {
	a := NewAcctNameANDActorID(in.Username, uint(in.ActorID))
	err := a.Delete()
	if err != nil {
		return &pb.AccountsReply{Code: "500", Message: "!ok"}, err
	}

	return &pb.AccountsReply{Code: "200", Message: "ok"}, nil
}
