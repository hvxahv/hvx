package accounts

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Devices struct {
	gorm.Model

	AccountID uint   `gorm:"primaryKey;type:bigint;account_id"`
	Device    string `gorm:"type:text;device"`
	DeviceID  string `gorm:"primaryKey;type:text;device_id"`
	Token     string `gorm:"type:text;token"`
	URL       string `gorm:"type:text;url"`
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

func (d *Devices) Get() (*[]Devices, error) {
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

func NewDevicesByID(id, accountID uint) *Devices {
	return &Devices{
		Model: gorm.Model{
			ID: id,
		},
		AccountID: accountID,
	}
}

func NewDevicesByAccountID(accountID uint) *Devices {
	return &Devices{AccountID: accountID}
}

func NewDevices(accountID uint, device string, deviceID string, token string, URL string) *Devices {
	return &Devices{AccountID: accountID, Device: device, DeviceID: deviceID, Token: token, URL: URL}
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

func (d *Devices) DeleteByID() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").
		Where("id = ?", d.ID).
		Where("account_id = ?", d.AccountID).
		Unscoped().Delete(&Devices{}).Error; err != nil {
		return err
	}
	// Token If it fails, notify the client to go offline.
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
	Get() (*[]Devices, error)
	IsNotExist() bool
	DeleteByID() error
	// DeleteByDeviceID This method is used when exiting the current device.
	DeleteByDeviceID() error

	DeleteALLByAccountID() error
}
