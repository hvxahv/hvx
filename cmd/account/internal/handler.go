/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	a := NewUsername(in.GetUsername())
	return &pb.IsExistResponse{IsExist: a.IsExist()}, nil
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	if err := NewAccountsCreate(in.GetUsername(), in.GetMail(), in.GetPassword()).Create(in.GetPublicKey()); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) GetByUsername(ctx context.Context, in *pb.GetByUsernameRequest) (*pb.GetByUsernameResponse, error) {
	account, err := NewUsername(in.GetUsername()).GetAccountByUsername()
	if err != nil {
		return nil, err
	}
	return &pb.GetByUsernameResponse{
		AccountId: int64(account.ID),
		Username:  account.Username,
		Mail:      account.Mail,
		Password:  account.Password,
		ActorId:   int64(account.ActorID),
		IsPrivate: account.IsPrivate,
	}, nil
}

// DeleteAccount need to delete the actor of the account, and all data related to this account.
func (s *server) DeleteAccount(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	// TODO - Delete Account related data.
	// TODO - Delete Actor related data.
	// TODO - Delete Device related data.
	// TODO - Delete Saved related data.
	// TODO - Delete Article related data.
	// TODO - Exit all channels.
	// TODO - Delete Message accounts.
	return nil, nil
}

func (s *server) EditUsername(ctx context.Context, in *pb.EditUsernameRequest) (*pb.EditUsernameResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewAccountsID(parse.AccountId).EditUsername(in.GetUsername()); err != nil {
		return nil, err
	}
	return &pb.EditUsernameResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) EditEmail(ctx context.Context, in *pb.EditEmailRequest) (*pb.EditEmailResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewAccountsID(parse.AccountId).EditEmail(in.GetMail()); err != nil {
		return nil, err
	}
	return &pb.EditEmailResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) EditPassword(ctx context.Context, in *pb.EditPasswordRequest) (*pb.EditPasswordResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewEditPassword(in.GetUsername(), in.GetPassword()).EditPassword(in.GetNewPassword()); err != nil {
		return nil, err
	}

	devices, err := clientv1.New(ctx, microsvc.DeviceServiceName).DeleteDevices(int64(parse.AccountId))
	if err != nil {
		errors.Throw("error occurred while connecting to the device service while edit the password.", err)
		return nil, err
	}

	// TODO - EDIT MATRIX ACCESS PASSWORD.
	return &pb.EditPasswordResponse{Code: "200", Status: devices.Status}, nil
}

func (s *server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	verify, err := NewVerify(in.GetUsername()).Verify(in.GetPassword())
	if err != nil {
		return &pb.VerifyResponse{
			Code:   "401",
			Status: "UNAUTHORIZED",
		}, err
	}
	return &pb.VerifyResponse{
		Code:      "200",
		Status:    "ok",
		AccountId: int64(verify.ID),
		Username:  verify.Username,
		Mail:      verify.Mail,
		ActorId:   int64(verify.ActorID),
	}, nil
}

func (s *server) GetPrivateKey(ctx context.Context, in *pb.GetPrivateKeyRequest) (*pb.GetPrivateKeyResponse, error) {
	x, err := NewAccountsID(uint(in.GetAccountId())).GetPrivateKey()
	if err != nil {
		return nil, err
	}
	return &pb.GetPrivateKeyResponse{
		Code:       "200",
		PrivateKey: x,
	}, nil
}

func (s *server) IAm(ctx context.Context, in *emptypb.Empty) (*pb.IAmResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	account, err := NewUsername(parse.Username).GetAccountByUsername()
	if err != nil {
		return nil, err
	}

	a, err := clientv1.New(ctx, microsvc.ActorServiceName).GetActor(int64(account.ActorID))
	if err != nil {
		return nil, err
	}
	return &pb.IAmResponse{
		Account: &pb.GetByUsernameResponse{
			AccountId: int64(account.ID),
			Username:  account.Username,
			Mail:      account.Mail,
			Password:  account.Password,
			ActorId:   int64(account.ActorID),
			IsPrivate: account.IsPrivate,
		},
		Actor: a.Actor,
	}, nil
}
