/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package device

import (
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	pb "github.com/hvxahv/hvxahv/api/device/v1alpha"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strconv"
)

type Devices struct {
	gorm.Model

	// ID Separate out the Device ID for use as a field in JSON.
	ID uint `gorm:"primaryKey" json:"ID,string"`

	// AccountID The associated account ID.
	AccountID uint `gorm:"primaryKey;type:bigint;account_id" json:"account_id,string"`

	// Device The device name.
	Device string `gorm:"type:text;device"`

	// Hash Unique hash identifier of the device.
	Hash string `gorm:"primaryKey;type:text;hash"`

	PrivateKey string `gorm:"type:text;privateKey"`
	PublicKey  string `gorm:"type:text;publicKey"`
}

func (a *device) DeviceIsExistByHash(ctx context.Context, in *pb.DeviceIsExistByHashRequest) (*pb.DeviceIsExistByHashResponse, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("devices").Where("hash = ?", in.Hash).First(&Devices{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return &pb.DeviceIsExistByHashResponse{IsExist: ok}, nil
	}

	return &pb.DeviceIsExistByHashResponse{IsExist: false}, nil
}

func (a *device) DeviceIsExistByID(ctx context.Context, in *pb.DeviceIsExistByIDRequest) (*pb.DeviceIsExistByIDResponse, error) {
	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	if err := db.Debug().
		Table("devices").
		Where("id = ?", uint(id)).
		First(&Devices{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return &pb.DeviceIsExistByIDResponse{IsExist: ok}, nil
	}

	return &pb.DeviceIsExistByIDResponse{IsExist: false}, nil
}

func (a *device) CreateDevice(ctx context.Context, in *pb.CreateDeviceRequest) (*pb.CreateDeviceResponse, error) {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Devices{}); err != nil {
		fmt.Println(err)
		return nil, err
	}

	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}

	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		return nil, err
	}

	v := NewDevices(uint(id), in.Ua, in.Hash, privateKey, publicKey)
	if err := db.Debug().Where("devices").Create(&v).Error; err != nil {
		return nil, err
	}

	return &pb.CreateDeviceResponse{
		DeviceId:  strconv.Itoa(int(v.ID)),
		PublicKey: publicKey,
	}, nil
}

func (a *device) GetDevicesByAccountID(ctx context.Context, in *pb.GetDevicesByAccountIDRequest) (*pb.GetDevicesByAccountIDResponse, error) {
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	var devices []*pb.Device
	db := cockroach.GetDB()
	if err := db.Debug().
		Table("devices").
		Where("account_id = ?", id).
		Find(&devices).
		Error; err != nil {
		return nil, err
	}
	return &pb.GetDevicesByAccountIDResponse{Code: "200", Devices: devices}, nil
}

func (a *device) GetDeviceByID(ctx context.Context, in *pb.GetDeviceByIDRequest) (*pb.Device, error) {
	db := cockroach.GetDB()
	id, err := strconv.Atoi(in.DeviceId)
	if err != nil {
		return nil, err
	}
	if err := db.Debug().Table("devices").Where("id = ?", id).First(&a.Devices).Error; err != nil {
		return nil, err
	}

	return &pb.Device{
		Id:         strconv.Itoa(int(a.Devices.ID)),
		AccountId:  strconv.Itoa(int(a.Devices.AccountID)),
		Device:     a.Devices.Device,
		Hash:       a.Devices.Hash,
		PrivateKey: a.Devices.PrivateKey,
		PublicKey:  a.Devices.PublicKey,
	}, nil
}

func (a *device) GetDeviceByHash(ctx context.Context, in *pb.GetDeviceByHashRequest) (*pb.Device, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("devices").Where("hash = ?", in.Hash).First(&a.Devices).Error; err != nil {
		return nil, err
	}

	return &pb.Device{
		Id:         strconv.Itoa(int(a.Devices.ID)),
		AccountId:  strconv.Itoa(int(a.Devices.AccountID)),
		Device:     a.Devices.Device,
		Hash:       a.Devices.Hash,
		PrivateKey: a.Devices.PrivateKey,
		PublicKey:  a.Devices.PublicKey,
	}, nil
}

func (a *device) DeleteDeviceAllByAccountID(ctx context.Context, in *pb.DeleteDeviceAllByAccountIDRequest) (*pb.DeleteDeviceAllByAccountIDResponse, error) {
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ?", id).Unscoped().Delete(&Devices{}).Error; err != nil {
		return nil, err
	}
	return &pb.DeleteDeviceAllByAccountIDResponse{Code: "200", Reply: "ok"}, nil
}

func (a *device) DeleteDeviceByID(ctx context.Context, in *pb.DeleteDeviceByIDRequest) (*pb.DeleteDeviceByIDResponse, error) {
	id, err := strconv.Atoi(in.DeviceId)
	if err != nil {
		return nil, err
	}
	aid, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	db := cockroach.GetDB()

	if err := db.Debug().
		Table("devices").
		Where("id = ? AND account_id = ?", uint(id), uint(aid)).
		Unscoped().
		Delete(&Devices{}).
		Error; err != nil {
		return nil, err
	}
	return &pb.DeleteDeviceByIDResponse{Code: "200", Reply: "ok"}, nil
}

func (a *device) DeleteDeviceByHash(ctx context.Context, in *pb.DeleteDeviceByHashRequest) (*pb.DeleteDeviceByHashResponse, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("hash = ?", in.Hash).Unscoped().Delete(&Devices{}).Error; err != nil {
		return nil, err
	}
	return &pb.DeleteDeviceByHashResponse{Code: "200", Reply: "ok"}, nil
}

func NewDevices(accountID uint, ua, hash, privateKey, publicKey string) *Devices {
	return &Devices{AccountID: accountID, Device: ua, Hash: hash, PrivateKey: privateKey, PublicKey: publicKey}
}
