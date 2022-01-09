package device

import (
	"fmt"
	"github.com/SherClockHolmes/webpush-go"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
	"log"
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

func (d *Devices) DeleteAll() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ?", d.AccountID).Unscoped().Delete(&Devices{}); err != nil {
		return err.Error
	}
	return nil
}

func (d *Devices) GetDeviceByHash() (*Devices, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ? AND hash = ?", d.AccountID, d.Hash).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Devices) Delete() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("account_id = ? AND hash = ?", d.AccountID, d.Hash).Unscoped().Delete(&Devices{}).Error; err != nil {
		return err
	}
	return nil
}

func NewDevices(accountID uint, device string, deviceID string) *Devices {
	privateKey, publicKey, err := webpush.GenerateVAPIDKeys()
	if err != nil {
		log.Println(err)
	}
	return &Devices{AccountID: accountID, Device: device, Hash: deviceID, PrivateKey: privateKey, PublicKey: publicKey}
}

func (d *Devices) IsNotExist() bool {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("hash = ?", d.Hash).First(&Devices{}); err != nil {
		if cockroach.IsNotFound(err.Error) {
			return cockroach.IsNotFound(err.Error)
		}
	}
	return false
}

func (d *Devices) GetDevice() (*Devices, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("devices").Where("id = ? AND account_id = ?", d.ID, d.AccountID).First(&d).Error; err != nil {
		return nil, err
	}
	return d, nil
}

func (d *Devices) GetDevicesByAccountID() (*[]Devices, error) {
	db := cockroach.GetDB()
	var devices []Devices
	if err := db.Debug().Table("devices").Where("account_id = ?", d.AccountID).Find(&devices).Error; err != nil {
		return nil, err
	}
	for _, i := range devices {
		log.Println(i.ID)
		devices = append(devices)
	}
	return &devices, nil
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

func NewDevicesIsNotExist(hash string) *Devices {
	return &Devices{Hash: hash}
}

func NewDeviceByHash(accountID uint, hash string) *Devices {
	return &Devices{
		AccountID: accountID,
		Hash:      hash,
	}
}

func NewDevicesByID(accountID, id uint) *Devices {
	return &Devices{
		ID:        id,
		AccountID: accountID,
	}
}

func NewDevicesByAccountID(accountID uint) *Devices {
	return &Devices{AccountID: accountID}
}

type Device interface {
	Create() error

	// GetDevice Get online device details by ID.
	GetDevice() (*Devices, error)

	// GetDeviceByHash Get the device through hash.
	GetDeviceByHash() (*Devices, error)

	// GetDevicesByAccountID Get all logged-in devices of the account.
	GetDevicesByAccountID() (*[]Devices, error)

	// IsNotExist Confirm whether the device HASH exists in the middleware. Should only be used for HTTP middleware.
	IsNotExist() bool

	Delete() error

	DeleteAll() error
}
