/*
 *
 * Copyright 2022 The hvxahv Authors.
 * * https://github.com/hvxahv/hvx **
 * * https://disism.com **
 * /
 */

package internal

import (
	pb "github.com/hvxahv/hvx/APIs/grpc-go/device/v1alpha1"
	"github.com/hvxahv/hvx/conv"
	"golang.org/x/net/context"
	"strconv"
)

func (a *server) DeviceIsExistByHash(ctx context.Context, in *pb.DeviceIsExistByHashRequest) (*pb.DeviceIsExistByHashResponse, error) {
	ok, err := NewDevicesHash(in.Hash).IsExistByHash()
	if err != nil {
		return nil, err
	}
	return &pb.DeviceIsExistByHashResponse{IsExist: ok}, nil
}

func (a *server) DeviceIsExistByID(ctx context.Context, in *pb.DeviceIsExistByIDRequest) (*pb.DeviceIsExistByIDResponse, error) {
	id, err := conv.StringToUint(in.Id)
	if err != nil {
		return nil, err
	}
	ok, err := NewDevicesId(id).IsExistById()
	if err != nil {
		return nil, err
	}

	return &pb.DeviceIsExistByIDResponse{IsExist: ok}, nil
}

func (a *server) CreateDevice(ctx context.Context, in *pb.CreateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	aid, err := conv.StringToUint(in.AccountId)
	if err != nil {
		return nil, err
	}
	create, err := NewDevices(aid, in.Ua).Create()
	if err != nil {
		return nil, err
	}
	return &pb.CreateDeviceResponse{
		DeviceId:  strconv.Itoa(int(create.ID)),
		PublicKey: "",
	}, nil
}

func (a *server) GetDevicesByAccountID(ctx context.Context, in *pb.GetDevicesByAccountIDRequest) (*pb.GetDevicesByAccountIDResponse, error) {
	aid, err := conv.StringToUint(in.AccountId)
	if err != nil {
		return nil, err
	}
	ds, err := NewDevicesAccountID(aid).GetListByAccountId()
	if err != nil {
		return nil, err
	}
	var res []*pb.Device
	for _, device := range ds {
		var pd pb.Device
		pd.Id = conv.UintToString(device.ID)
		pd.AccountId = conv.UintToString(device.AccountID)
		pd.Device = device.Device
		pd.Hash = device.Hash
		res = append(res, &pd)
	}

	return &pb.GetDevicesByAccountIDResponse{Code: "200", Devices: res}, nil
}

func (a *server) GetDeviceByID(ctx context.Context, in *pb.GetDeviceByIDRequest) (*pb.Device, error) {
	did, err := conv.StringToUint(in.DeviceId)
	if err != nil {
		return nil, err
	}
	d, err := NewDevicesId(did).GetById()
	if err != nil {
		return nil, err
	}
	return &pb.Device{
		Id:        strconv.Itoa(int(d.ID)),
		AccountId: strconv.Itoa(int(d.AccountID)),
		Device:    d.Device,
		Hash:      d.Hash,
	}, nil
}

func (a *server) GetDeviceByHash(ctx context.Context, in *pb.GetDeviceByHashRequest) (*pb.Device, error) {
	d, err := NewDevicesHash(in.Hash).GetByHash()
	if err != nil {
		return nil, err
	}
	return &pb.Device{
		Id:        strconv.Itoa(int(d.ID)),
		AccountId: strconv.Itoa(int(d.AccountID)),
		Device:    d.Device,
		Hash:      d.Hash,
	}, nil
}
