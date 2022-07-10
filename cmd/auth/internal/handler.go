/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/grpc/v1alpha1/auth"
	"golang.org/x/net/context"
)

func (s *server) Authorization(ctx context.Context, in *pb.AuthorizationRequest) (*pb.AuthorizationResponse, error) {

	// TODO: implement the authorization logic here.
	// You can use the `in` parameter to get the username and password.

	// TODO: implement the create devices.
	//cli, err := clientv1.New(ctx,
	//	cfg.SetEndpoints(microsvc.GetGRPCServiceAddress("device")),
	//	cfg.SetDialTimeout(10*time.Second),
	//)
	//if err != nil {
	//	return nil, err
	//}
	//defer cli.Close()
	//device, err := cli.CreateDevice(ctx, &v1alpha.CreateDeviceRequest{
	//	AccountId: conv.UintToString(a.ID),
	//	Ua:        in.Ua,
	//	Hash:      uuid.New().String(),
	//})
	//if err != nil {
	//	return nil, err
	//}

	// TODO: implement Generate TOKEN...
	//expired := time.Duration(viper.GetInt("authentication.token.expired"))
	//g := jwt.NewClaims(a.Mail, conv.UintToString(a.ID), a.Username, device.DeviceId, expired)
	//k, err := g.JWTTokenGenerator(viper.GetString("authentication.token.signed"))
	//if err != nil {
	//	return nil, err
	//}

	return &pb.AuthorizationResponse{
		Code:               "",
		Status:             "",
		Id:                 "",
		AuthorizationToken: "",
		ActorId:            "",
		Mail:               "",
		DeviceId:           "",
	}, nil
}
