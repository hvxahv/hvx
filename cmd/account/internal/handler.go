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
	"strconv"

	pb "github.com/hvxahv/hvx/APIs/v1alpha1/account"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
)

func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	a := NewUsername(in.Username)
	return &pb.IsExistResponse{IsExist: a.IsExist()}, nil
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	if err := NewAccountsCreate(in.Username, in.Mail, in.Password).Create(in.PublicKey); err != nil {
		return nil, err
	}

	return &pb.CreateResponse{Code: "200", Reply: "ok"}, nil
}

func (s *server) GetByUsername(ctx context.Context, in *pb.GetByUsernameRequest) (*pb.GetByUsernameResponse, error) {
	a, err := NewUsername(in.Username).GetAccountByUsername()
	if err != nil {
		return nil, err
	}
	return &pb.GetByUsernameResponse{
		AccountId: strconv.Itoa(int(a.ID)),
		Username:  a.Username,
		Mail:      a.Mail,
		Password:  a.Password,
		ActorId:   strconv.Itoa(int(a.ActorID)),
		IsPrivate: strconv.FormatBool(a.IsPrivate),
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

	if err := NewAccountsID(parse.AccountId).EditUsername(in.Username); err != nil {
		return nil, err
	}
	return &pb.EditUsernameResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) EditEmail(ctx context.Context, in *pb.EditEmailRequest) (*pb.EditEmailResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewAccountsID(parse.AccountId).EditEmail(in.Mail); err != nil {
		return nil, err
	}
	return &pb.EditEmailResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) EditPassword(ctx context.Context, in *pb.EditPasswordRequest) (*pb.EditPasswordResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err := NewEditPassword(in.Username, in.Password).EditPassword(in.NewPassword); err != nil {
		return nil, err
	}

	devices, err := clientv1.New(ctx, microsvc.DeviceServiceName).DeleteDevices(strconv.Itoa(int(parse.AccountId)))
	if err != nil {
		errors.Throw("error occurred while connecting to the device service while edit the password.", err)
		return nil, err
	}
	// TODO - EDIT MATRIX ACCESS PASSWORD.
	return &pb.EditPasswordResponse{Code: "200", Reply: devices.Reply}, nil
}

func (s *server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	verify, err := NewVerify(in.Username).Verify(in.Password)
	if err != nil {
		return &pb.VerifyResponse{
			Code:   "401",
			Status: "UNAUTHORIZED",
		}, err
	}
	return &pb.VerifyResponse{
		Code:     "200",
		Status:   "ok",
		Id:       strconv.Itoa(int(verify.ID)),
		Username: verify.Username,
		Mail:     verify.Mail,
		ActorId:  strconv.Itoa(int(verify.ActorID)),
	}, nil
}

func (s *server) GetPrivateKey(ctx context.Context, in *pb.GetPrivateKeyRequest) (*pb.GetPrivateKeyResponse, error) {
	id, err := strconv.Atoi(in.GetAccountId())
	if err != nil {
		return nil, err
	}
	x, err := NewAccountsID(uint(id)).GetPrivateKey()
	if err != nil {
		return nil, err
	}
	return &pb.GetPrivateKeyResponse{
		Code:       "200",
		PrivateKey: x,
	}, nil
}
