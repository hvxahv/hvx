package accounts

import (
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
	"log"
)

type Devices struct {
	gorm.Model

	AccountID  uint   `gorm:"primaryKey;type:bigint;account_id"`
	Device     string `gorm:"type:text;device"`
	DeviceID   string `gorm:"primaryKey;type:text;device_id"`
	PrivateKey string `gorm:"type:text;privateKey"`
	PublicKey  string `gorm:"type:text;publicKey"`
}

func NewDevices(accountID uint, device string, deviceID string) *Devices {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Println(err)
	}
	return &Devices{AccountID: accountID, Device: device, DeviceID: deviceID, PrivateKey: privateKey, PublicKey: publicKey}
}

func (d *Devices) DeleteALLByAccountID() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ?", d.AccountID).Unscoped().Delete(&Devices{}); err != nil {
		return err.Error
	}
	return nil
}

func (d *Devices) IsNotExist() bool {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("device_id = ?", d.DeviceID).First(&Devices{}); err != nil {
		if cockroach.IsNotFound(err.Error) {
			return cockroach.IsNotFound(err.Error)
		}
	}
	return false
}

func (d *Devices) GetDevicesByID() (*Devices, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("id = ?", d.ID).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Devices) GetDevicesByDeviceID() (*Devices, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("device_id = ?", d.DeviceID).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Devices) GetDevices() (*[]Devices, error) {
	db := cockroach.GetDB()
	var devices []Devices
	if err := db.Debug().Table("devices").Where("account_id = ?", d.AccountID).Find(&devices).Error; err != nil {
		return nil, err
	}
	return &devices, nil
}

func NewDevicesByDeviceID(accountID uint, deviceID string) *Devices {
	return &Devices{
		AccountID: accountID,
		DeviceID:  deviceID,
	}
}

func NewDevicesID(deviceID string) *Devices {
	return &Devices{
		DeviceID: deviceID,
	}
}

func NewDevicesByID(id uint) *Devices {
	return &Devices{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func NewDevicesByAccountID(accountID uint) *Devices {
	return &Devices{AccountID: accountID}
}

func (d *Devices) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Devices{}); err != nil {
		fmt.Println(err)
		return err
	}

	if err := db.Debug().Where("devices").Create(&d).Error; err != nil {
		return err
	}
	return nil
}

func (d *Devices) DeleteByDeviceID() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ?", d.AccountID).
		Where("device_id = ?", d.DeviceID).
		Unscoped().Delete(&Devices{}).Error; err != nil {
		return err
	}
	return nil
}

type Device interface {
	Create() error
	GetDevicesByID() (*Devices, error)
	GetDevicesByDeviceID() (*Devices, error)
	GetDevices() (*[]Devices, error)
	IsNotExist() bool
	// DeleteByDeviceID This method is used when exiting the current device.
	DeleteByDeviceID() error
	DeleteALLByAccountID() error
}
