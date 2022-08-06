/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"gorm.io/gorm"
)

const (
	DeviceTable = "devices"
)

type Devices struct {
	gorm.Model

	// AccountID The associated account ID.
	AccountID uint `gorm:"primaryKey;type:bigint;account_id"`

	// Device The device name.
	Device string `gorm:"type:text;device"`
}

type Device interface {
	// IsExist Get the result of whether the device exists by hash,
	// return true if it exists, otherwise it will return false.
	IsExist() (bool, error)

	Get() (*Devices, error)

	// Create Device entity with account ID and user device identifier (ua).
	Create() (*Devices, error)

	// Delete the device by id and account id.
	Delete() error

	DeleteDevices() error
}

func NewDevicesId(id uint) *Devices {
	return &Devices{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func (d *Devices) IsExist() (bool, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table(DeviceTable).Where("id = ?", d.ID).First(&Devices{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return ok, nil
	}
	return false, nil
}

func (d *Devices) Get() (*Devices, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table(DeviceTable).Where("id = ?", d.ID).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}

func NewDevices(accountID uint, ua string) *Devices {
	return &Devices{AccountID: accountID, Device: ua}
}

func (d *Devices) Create() (*Devices, error) {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Devices{}); err != nil {
		return nil, err
	}

	v := NewDevices(d.AccountID, d.Device)
	if err := db.Debug().Where(DeviceTable).Create(&v).Error; err != nil {
		return nil, err
	}
	return v, nil

}

func NewDevicesAccountId(accountId uint) *Devices {
	return &Devices{
		AccountID: accountId,
	}
}

func (d *Devices) GetDevices() ([]*Devices, error) {
	db := cockroach.GetDB()
	var ds []*Devices
	if err := db.Debug().Table(DeviceTable).
		Where("account_id = ?", d.AccountID).Find(&ds).Error; err != nil {
		return nil, err
	}
	return ds, nil
}

func NewDevicesDelete(id, accountID uint) *Devices {
	return &Devices{
		Model: gorm.Model{
			ID: id,
		},
		AccountID: accountID,
	}
}

func (d *Devices) Delete() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table(DeviceTable).
		Where("id = ? AND account_id = ?", d.ID, d.AccountID).Unscoped().Delete(&Devices{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (d *Devices) DeleteDevices() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table(DeviceTable).
		Where("account_id = ?", d.AccountID).Unscoped().Delete(&Devices{}).Error; err != nil {
		return err
	}
	return nil
}
