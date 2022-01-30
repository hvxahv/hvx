package device

import (
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	pb "github.com/hvxahv/hvxahv/api/device/v1alpha1"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"strconv"
)

type Devices struct {
	gorm.Model

	ID         uint   `gorm:"primaryKey" json:"ID,string"`
	AccountID  uint   `gorm:"primaryKey;type:bigint;account_id" json:"account_id,string"`
	Device     string `gorm:"type:text;device"`
	Hash       string `gorm:"primaryKey;type:text;hash"`
	PrivateKey string `gorm:"type:text;privateKey"`
	PublicKey  string `gorm:"type:text;publicKey"`
}

func (d *device) IsExist(ctx context.Context, in *pb.NewDeviceHash) (*pb.IsDeviceExistReply, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("hash = ?", in.Hash).First(&Devices{}); err != nil {
		if cockroach.IsNotFound(err.Error) {
			return &pb.IsDeviceExistReply{IsExist: false}, nil
		}
	}

	return &pb.IsDeviceExistReply{IsExist: true}, nil
}

func (d *device) Create(ctx context.Context, in *pb.NewDeviceCreate) (*pb.DeviceCreateReply, error) {
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

	return &pb.DeviceCreateReply{
		PublicKey: publicKey,
	}, nil
}

func (d *device) GetDevicesByAccountID(ctx context.Context, in *pb.NewDeviceAccountID) (*pb.DevicesData, error) {
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	var devices []*pb.DeviceData
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ?", id).Find(&devices).Error; err != nil {
		return nil, err
	}
	return &pb.DevicesData{Code: "200", Devices: devices}, nil
}

func (d *device) GetDeviceByHash(ctx context.Context, in *pb.NewDeviceHash) (*pb.DeviceData, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("devices").Where("hash = ?", in.Hash).First(&d.Devices).Error; err != nil {
		return nil, err
	}

	return &pb.DeviceData{
		Id:         strconv.Itoa(int(d.ID)),
		AccountId:  strconv.Itoa(int(d.Devices.AccountID)),
		Device:     d.Devices.Device,
		Hash:       d.Devices.Hash,
		PrivateKey: d.Devices.PrivateKey,
		PublicKey:  d.Devices.PublicKey,
	}, nil
}

func (d *device) DeleteAllByAccountID(ctx context.Context, in *pb.NewDeviceAccountID) (*pb.DeviceReply, error) {
	id, err := strconv.Atoi(in.AccountId)
	if err != nil {
		return nil, err
	}
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ?", id).Unscoped().Delete(&Devices{}).Error; err != nil {
		return nil, err
	}
	return &pb.DeviceReply{Code: "200", Reply: "ok"}, nil
}

func (d *device) Delete(ctx context.Context, in *pb.NewDeviceID) (*pb.DeviceReply, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("id = ?", id).Unscoped().Delete(&Devices{}).Error; err != nil {
		return nil, err
	}
	return &pb.DeviceReply{Code: "200"}, nil
}

func (d *device) DeleteByDeviceHash(ctx context.Context, in *pb.NewDeviceHash) (*pb.DeviceReply, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("hash = ?", in.Hash).Unscoped().Delete(&Devices{}).Error; err != nil {
		return nil, err
	}
	return &pb.DeviceReply{Code: "200"}, nil
}

func NewDevices(accountID uint, ua, hash, privateKey, publicKey string) *Devices {
	return &Devices{AccountID: accountID, Device: ua, Hash: hash, PrivateKey: privateKey, PublicKey: publicKey}
}
