/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package account

import (
	pb "github.com/hvxahv/hvx/api/grpc/proto/account/v1alpha1"
	"github.com/hvxahv/hvx/pkg/microsvc"
	"github.com/hvxahv/hvx/pkg/microsvc/client"
	"github.com/hvxahv/hvx/pkg/transport"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (a *server) CreateAccount(ctx context.Context, in *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	if err := NewCreateAccounts(in.Username, in.Mail, in.Password).Create(in.PublicKey); err != nil {
		return nil, err
	}
	return &pb.CreateAccountResponse{Code: "200", Reply: "ok"}, nil
}

func (a *server) DeleteAccount(ctx context.Context, in *pb.DeleteAccountRequest) (*pb.DeleteAccountResponse, error) {
	username, err := microsvc.GetUsernameByTokenWithContext(ctx)
	if err != nil {
		return nil, err
	}
	if err := NewDeleteAccount(username, in.Password).DeleteAccount(in.Password); err != nil {
		return nil, err
	}
	// Connect to the device service to delete all login information for the account.
	conn, err := grpc.DialContext(ctx, microsvc.GetGRPCServiceAddress("device"),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(transport.CustomerTokenAuth{}),
	)
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	cli := client.NewHvxClient(conn)
	reply, err := cli.DeleteDeviceAllByAccountID(ctx, &emptypb.Empty{})
	if err != nil {
		return nil, err
	}

	return &pb.DeleteAccountResponse{Code: "200", Reply: reply.Reply}, nil
}
