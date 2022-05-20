/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/google/uuid"
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

	// Hash Unique hash identifier of the device.
	Hash string `gorm:"primaryKey;type:text;hash"`
}

type device interface {
	// IsExistByHash Get the result of whether the device exists by hash,
	// return true if it exists, otherwise it will return false.
	IsExistByHash() (bool, error)

	// GetByHash Get the device by hash.
	GetByHash() (*Devices, error)

	// IsExistById Get the result of whether the device exists by id,
	// return true if it exists, otherwise it will return false.
	IsExistById() (bool, error)

	// GetById Get the device by id.
	GetById() (*Devices, error)

	// Create Device entity with account ID and user device identifier (ua).
	Create() (*Devices, error)

	// GetListByAccountId Get the list of devices by account ID.
	GetListByAccountId() ([]*Devices, error)

	// Delete the device by id and account id.
	Delete() error

	// DeleteAccountDevices Delete all devices by account id.
	DeleteAccountDevices() error
}

func NewDevicesHash(hash string) *Devices {
	return &Devices{Hash: hash}
}

func NewDevicesId(id uint) *Devices {
	return &Devices{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func NewDevicesAccountID(accountId uint) *Devices {
	return &Devices{
		AccountID: accountId,
	}
}

func NewDevices(accountID uint, ua string) *Devices {
	hash := uuid.New().String()
	return &Devices{AccountID: accountID, Device: ua, Hash: hash}
}

func NewDevicesDelete(id, accountID uint) *Devices {
	return &Devices{
		Model: gorm.Model{
			ID: id,
		},
		AccountID: accountID,
	}
}

func (d *Devices) IsExistByHash() (bool, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table(DeviceTable).Where("hash = ?", d.Hash).First(&Devices{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return !ok, nil
	}
	return true, nil
}

func (d *Devices) GetByHash() (*Devices, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table(DeviceTable).
		Where("hash = ?", d.Hash).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Devices) IsExistById() (bool, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table(DeviceTable).Where("id = ?", d.ID).First(&Devices{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		return !ok, nil
	}
	return true, nil
}

func (d *Devices) GetById() (*Devices, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table(DeviceTable).Where("id = ?", d.ID).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
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

func (d *Devices) GetListByAccountId() ([]*Devices, error) {
	db := cockroach.GetDB()
	var ds []*Devices
	if err := db.Debug().Table(DeviceTable).
		Where("account_id = ?", d.AccountID).Find(&ds).Error; err != nil {
		return nil, err
	}
	return ds, nil
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

func (d *Devices) DeleteAccountDevices() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table(DeviceTable).
		Where("account_id = ?", d.AccountID).Unscoped().Delete(&Devices{}).Error; err != nil {
		return err
	}
	return nil
}
