/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	"strconv"
	"time"

	"github.com/google/uuid"
	pb "github.com/hvxahv/hvx/APIs/grpc-go/account/v1alpha1"
	v1alpha "github.com/hvxahv/hvx/APIs/grpc-go/device/v1alpha1"
	"github.com/hvxahv/hvx/client/clientv1"
	"github.com/hvxahv/hvx/conv"
	"github.com/hvxahv/hvx/identity"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

// IsExist ...
func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	a := NewUsername(in.Username)
	return &pb.IsExistResponse{IsExist: a.IsExist()}, nil
}

// CreateAccount ...
func (s *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	if err := NewAccountsCreate(in.Username, in.Mail, in.Password).Create(in.PublicKey); err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{Code: "200", Reply: "ok"}, nil
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
func (s *server) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	username, err := microsvc.GetUsernameByTokenWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewAccountsDelete(username, in.Password).Delete(); err != nil {
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
	if err := NewEditPassword(in.Username, in.Password).EditPassword(in.New); err != nil {
		return nil, err
	}
	// TODO - Edit Account related data.
	return &pb.EditPasswordResponse{Code: "200", Reply: "ok"}, nil
}

// GetActorByUsername ...
func (s *server) GetActorByUsername(ctx context.Context, in *pb.GetActorByUsernameRequest) (*pb.ActorDataResponse, error) {
	actor, err := NewUsername(in.Username).GetActorByUsername()
	if err != nil {
		return nil, err
	}
	return &pb.ActorDataResponse{
		Id:                strconv.Itoa(int(actor.ID)),
		PreferredUsername: actor.PreferredUsername,
		Domain:            actor.Domain,
		Avatar:            actor.Avatar,
		Name:              actor.Name,
		Summary:           actor.Summary,
		Inbox:             actor.Inbox,
		Address:           actor.Address,
		PublicKey:         actor.PublicKey,
		ActorType:         actor.ActorType,
		IsRemote:          strconv.FormatBool(actor.IsRemote),
	}, nil
}

// Verify ...
func (s *server) Verify(ctx context.Context, in *pb.VerifyRequest) (*pb.VerifyResponse, error) {
	a, err := NewVerify(in.Username).Verify(in.Password)
	if err != nil {
		return nil, err
	}

	cli, err := clientv1.New(ctx,
		clientv1.SetEndpoints(microsvc.GetGRPCServiceAddress("device")),
		clientv1.SetDialTimeout(10*time.Second),
	)
	if err != nil {
		return nil, err
	}
	defer cli.Close()
	device, err := cli.CreateDevice(ctx, &v1alpha.CreateDeviceRequest{
		AccountId: conv.UintToString(a.ID),
		Ua:        in.Ua,
		Hash:      uuid.New().String(),
	})
	if err != nil {
		return nil, err
	}

	// Creating an authorization token.
	k, err := identity.GenToken(strconv.Itoa(int(a.ID)), a.Mail, in.Username, device.DeviceId)
	if err != nil {
		return nil, err
	}

	return &pb.VerifyResponse{
		Code:      "200",
		Reply:     "ok",
		Id:        strconv.Itoa(int(a.ID)),
		Token:     k,
		Mail:      a.Mail,
		DeviceId:  device.DeviceId,
		PublicKey: device.PublicKey,
	}, nil
}

// CreateActor ...
func (s *server) CreateActor(ctx context.Context, in *pb.CreateActorRequest) (*pb.CreateActorResponse, error) {
	actor, err := NewActors(in.PreferredUsername, in.PublicKey, in.ActorType).Create()
	if err != nil {
		return nil, err
	}
	return &pb.CreateActorResponse{Code: "200", ActorId: strconv.Itoa(int(actor.ID))}, nil
}

// GetActorsByPreferredUsername ...
func (s *server) GetActorsByPreferredUsername(ctx context.Context, in *pb.GetActorsByPreferredUsernameRequest) (*pb.GetActorsByPreferredUsernameResponse, error) {
	actors, err := NewPreferredUsername(in.GetPreferredUsername()).GetActorsByPreferredUsername()
	if err != nil {
		return nil, err
	}

	var a []*pb.ActorDataResponse
	for _, v := range actors {
		var ad pb.ActorDataResponse
		ad.Id = conv.UintToString(v.ID)
		ad.PreferredUsername = v.PreferredUsername
		ad.Domain = v.Domain
		ad.Avatar = v.Avatar
		ad.Name = v.Name
		ad.Summary = v.Summary
		ad.Inbox = v.Inbox
		ad.Address = v.Address
		ad.PublicKey = v.PublicKey
		ad.ActorType = v.ActorType
		ad.IsRemote = strconv.FormatBool(v.IsRemote)
		a = append(a, &ad)
	}

	return &pb.GetActorsByPreferredUsernameResponse{Code: "200", Actors: a}, nil
}

// GetActorByAddress ...
func (s *server) GetActorByAddress(ctx context.Context, in *pb.GetActorByAddressRequest) (*pb.ActorDataResponse, error) {
	actor, err := NewActorAddress(in.GetAddress()).GetActorByAddress()
	if err != nil {
		return nil, err
	}

	return &pb.ActorDataResponse{
		Id:                strconv.Itoa(int(actor.ID)),
		PreferredUsername: actor.PreferredUsername,
		Domain:            actor.Domain,
		Avatar:            actor.Avatar,
		Name:              actor.Name,
		Summary:           actor.Summary,
		Inbox:             actor.Inbox,
		Address:           actor.Address,
		PublicKey:         actor.PublicKey,
		ActorType:         actor.ActorType,
		IsRemote:          strconv.FormatBool(actor.IsRemote),
	}, nil
}

// EditActor ...
func (s *server) EditActor(ctx context.Context, in *pb.EditActorRequest) (*pb.EditActorResponse, error) {
	id, err := microsvc.GetAccountIDWithContext(ctx)
	if err != nil {
		return nil, err
	}

	actor := new(Actors)
	switch {
	case in.Avatar != "":
		actor.SetActorAvatar(in.Avatar)
	case in.Name != "":
		actor.SetActorName(in.Name)
	case actor.Summary != "":
		actor.SetActorSummary(in.Summary)
	}

	if err := actor.EditActor(id); err != nil {
		return nil, err
	}
	return &pb.EditActorResponse{Code: "200", Reply: "ok"}, nil
}
