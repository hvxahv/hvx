/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package account

import (
	"strconv"
	"time"

	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"github.com/hvxahv/hvx/pkg/microsvc"
	"github.com/hvxahv/hvx/pkg/microsvc/client/v1/clientv1"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	a := NewUsername(in.Username)
	return &pb.IsExistResponse{IsExist: a.IsExist()}, nil
}

func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	if err := NewCreateAccounts(in.Username, in.Mail, in.Password).Create(in.PublicKey); err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{Code: "200", Reply: "ok"}, nil
}

func (s *server) GetAccountByUsername(ctx context.Context, in *pb.GetAccountByUsernameRequest) (*pb.GetAccountByUsernameResponse, error) {
	a, err := NewUsername(in.Username).GetAccountByUsername()
	if err != nil {
		return nil, err
	}
	return &pb.GetAccountByUsernameResponse{
		AccountId: strconv.Itoa(int(a.ID)),
		Username:  a.Username,
		Mail:      a.Mail,
		Password:  a.Password,
		ActorId:   strconv.Itoa(int(a.ActorID)),
		IsPrivate: strconv.FormatBool(a.IsPrivate),
	}, nil
}

func (s *server) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	username, err := microsvc.GetUsernameByTokenWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewDeleteAccount(username, in.Password).DeleteAccount(in.Password); err != nil {
		return nil, err
	}

	// Connect to the device service to delete all login information for the account.
	cli, err := clientv1.New(ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("device")),
		clientv1.SetDialOptionsWithToken(),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()

	reply, err := cli.DeleteDeviceAllByAccountID(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAccountResponse{Code: "200", Reply: reply.Reply}, nil
}

func (s *server) EditUsername(ctx context.Context, in *pb.EditUsernameRequest) (*pb.EditUsernameResponse, error) {
	id, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewEditAccountID(id).EditUsername(in.Username); err != nil {
		return nil, err
	}
	return &pb.EditUsernameResponse{Code: "200", Reply: "ok"}, nil
}

func (s *server) EditPassword(ctx context.Context, in *pb.EditPasswordRequest) (*pb.EditPasswordResponse, error) {
	if err := NewEditPassword(in.Username, in.Password).EditPassword(in.New); err != nil {
		return nil, err
	}
	// TODO - Edit Account related data.
	return &pb.EditPasswordResponse{Code: "200", Reply: "ok"}, nil
}

func (s *server) EditEmail(ctx context.Context, in *pb.EditEmailRequest) (*pb.EditEmailResponse, error) {
	id, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewEditAccountID(id).EditEmail(in.Mail); err != nil {
		return nil, err
	}
	return &pb.EditEmailResponse{Code: "200", Reply: "ok"}, nil
}
