/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/v1alpha1/device"
	"github.com/hvxahv/hvx/microsvc"
	"golang.org/x/net/context"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {

	exist, err := NewDevicesId(uint(in.GetId())).IsExist()
	if err != nil {
		return nil, err
	}
	return &pb.IsExistResponse{IsExist: exist}, nil
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	create, err := NewDevices(uint(in.AccountId), in.Ua).Create()
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		DeviceId: int64(create.ID),
	}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.Device, error) {

	d, err := NewDevicesId(uint(in.GetDeviceId())).Get()
	if err != nil {
		return nil, err
	}
	return &pb.Device{
		Id:        int64(d.ID),
		AccountId: int64(d.AccountID),
		Device:    d.Device,
	}, nil
}

func (s *server) GetDevices(ctx context.Context, g *emptypb.Empty) (*pb.GetDevicesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	devices, err := NewDevicesAccountId(parse.AccountId).GetDevices()
	if err != nil {
		return nil, err
	}
	var res []*pb.Device
	for _, d := range devices {
		var pd pb.Device
		pd.Id = int64(d.ID)
		pd.AccountId = int64(d.AccountID)
		pd.Device = d.Device
		pd.CreatedAt = d.CreatedAt.Format("2006-01-02 15:04:05")
		res = append(res, &pd)
	}

	return &pb.GetDevicesResponse{Code: "200", Devices: res}, nil
}

func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}

	if err = NewDevicesDelete(uint(in.DeviceId), parse.AccountId).Delete(); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Code: "200", Status: "ok"}, nil
}

func (s *server) DeleteDevices(ctx context.Context, in *pb.DeleteDevicesRequest) (*pb.DeleteDevicesResponse, error) {
	if err := NewDevicesAccountId(uint(in.GetAccountId())).DeleteDevices(); err != nil {
		return nil, err
	}
	return &pb.DeleteDevicesResponse{Code: "200", Status: "ok"}, nil
}
