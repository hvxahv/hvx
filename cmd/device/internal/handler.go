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
	"strconv"
)

func (s *server) IsExist(ctx context.Context, in *pb.IsExistRequest) (*pb.IsExistResponse, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {

	}
	exist, err := NewDevicesId(uint(id)).IsExist()
	if err != nil {
		return nil, err
	}
	return &pb.IsExistResponse{IsExist: exist}, nil
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	create, err := NewDevices(uint(aid), in.Ua).Create()
	if err != nil {
		return nil, err
	}
	return &pb.CreateResponse{
		DeviceId: strconv.Itoa(int(create.ID)),
	}, nil
}

func (s *server) Get(ctx context.Context, in *pb.GetRequest) (*pb.Device, error) {
	id, err := strconv.Atoi(in.DeviceId)
	if err != nil {
		return nil, err
	}
	d, err := NewDevicesId(uint(id)).Get()
	if err != nil {
		return nil, err
	}
	return &pb.Device{
		Id:        strconv.Itoa(int(d.ID)),
		AccountId: strconv.Itoa(int(d.AccountID)),
		Device:    d.Device,
	}, nil
}

func (s *server) GetDevices(ctx context.Context, g *emptypb.Empty) (*pb.GetDevicesResponse, error) {
	parse, err := microsvc.GetUserdataByAuthorizationToken(ctx)
	if err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(parse.AccountId)
	if err != nil {
		return nil, err
	}
	devices, err := NewDevicesAccountId(uint(aid)).GetDevices()
	if err != nil {
		return nil, err
	}
	var res []*pb.Device
	for _, d := range devices {
		var pd pb.Device
		pd.Id = strconv.Itoa(int(d.ID))
		pd.AccountId = strconv.Itoa(int(d.AccountID))
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
	aid, err := strconv.Atoi(parse.AccountId)
	if err != nil {
		return nil, err
	}

	id, err := strconv.Atoi(in.DeviceId)
	if err != nil {
		return nil, err
	}
	if err = NewDevicesDelete(uint(id), uint(aid)).Delete(); err != nil {
		return nil, err
	}
	return &pb.DeleteResponse{Code: "200", Reply: "ok"}, nil
}

func (s *server) DeleteDevices(ctx context.Context, in *pb.DeleteDevicesRequest) (*pb.DeleteDevicesResponse, error) {
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	err = NewDevicesAccountId(uint(aid)).DeleteDevices()
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDevicesResponse{Code: "200", Reply: "ok"}, nil
}
