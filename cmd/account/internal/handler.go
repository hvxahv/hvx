/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/account"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"strconv"
)

// IsExist ...
func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	a := NewUsername(in.Username)
	return &pb.IsExistResponse{IsExist: a.IsExist()}, nil
}

// Create ...
func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	if err := NewAccountsCreate(in.Username, in.Mail, in.Password).Create(in.PublicKey); err != nil {
		return nil, err
	}
	return &pb.CreateResponse{Code: "200", Reply: "ok"}, nil
}

// GetByUsername ...
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

// DeleteAccount ...
func (s *server) DeleteAccount(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	username, err := microsvc.GetUsernameByTokenWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewAccountsDelete(username, in.Password).Delete(); err != nil {
		return nil, err
	}

	// TODO - delete devices...
	// Connect to the device service to delete all login information for the account.
	//cli, err := clientv1.New(ctx,
	//	cfg.SetEndpoints(microsvc.GetGRPCServiceAddress("device")),
	//	cfg.SetDialOptionsWithToken(),
	//	cfg.SetDialTimeout(10*time.Second),
	//)
	//if err != nil {
	//	return nil, err
	//}
	//defer cli.Close()
	//
	//reply, err := cli.DeleteDeviceAllByAccountID(ctx, &emptypb.Empty{})
	//if err != nil {
	//	return nil, err
	//}

	// Delete actor ...
	//client, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")})
	//if err != nil {
	//	return nil, err
	//}
	//d, err := client.Delete(ctx, &empty.Empty{})
	//if err != nil {
	//	return nil, err
	//}
	//if d.Status != "ok" {
	//	return nil, fmt.Errorf(errors.ErrActorDelete)
	//}
	//
	//return &pb.DeleteResponse{
	//	Code:  "200",
	//	Reply: d.Status,
	//}, nil

	return nil, nil
}

// EditUsername ...
func (s *server) EditUsername(ctx context.Context, in *pb.EditUsernameRequest) (*pb.EditUsernameResponse, error) {
	id, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewAccountsID(id).EditUsername(in.Username); err != nil {
		return nil, err
	}
	return &pb.EditUsernameResponse{Code: "200", Reply: "ok"}, nil
}

// EditEmail ...
func (s *server) EditEmail(ctx context.Context, in *pb.EditEmailRequest) (*pb.EditEmailResponse, error) {
	id, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewAccountsID(id).EditEmail(in.Mail); err != nil {
		return nil, err
	}
	return &pb.EditEmailResponse{Code: "200", Reply: "ok"}, nil
}

// EditPassword ...
func (s *server) EditPassword(ctx context.Context, in *pb.EditPasswordRequest) (*pb.EditPasswordResponse, error) {
	if err := NewEditPassword(in.Username, in.Password).EditPassword(in.NewPassword); err != nil {
		return nil, err
	}
	// TODO - Edit Account related data.
	return &pb.EditPasswordResponse{Code: "200", Reply: "ok"}, nil
}

// Verify ...
func (s *server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	verify, err := NewVerify(in.Username).Verify(in.Password)
	if err != nil {
		return &pb.VerifyResponse{
			Code:   "401",
			Status: "unauthorized",
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
