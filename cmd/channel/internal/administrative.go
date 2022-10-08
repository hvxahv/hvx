package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"gorm.io/gorm"
)

type Administrates struct {
	gorm.Model

	ChannelId uint `gorm:"primaryKey;channel_id"`

	// AdminId admin actor id.
	AdminId uint `gorm:"primaryKey;admin_id"`
	IsOwner bool `gorm:"type:boolean;is_owner"`
}

// Note:
// In the process of implementation, the need to pay attention to the
// division of permissions between channel owners and administrators.

const (
	// AdministrateTable is the table name for the administrates table.
	AdministrateTable = "administrates"
)

type Administrator interface {
	IsAdministrator() bool
	IsChannelOwner() bool
	AddAdministratorOwner() error
	AddAdministrator(addedID uint) error
	RemoveAdministrator(removedId uint) error
	GetAdministrators() ([]*Administrates, error)
	ExitAdministrator() error
	DeleteAdministrators() error
}

func NewAdministratesPermission(channelId, adminId uint) *Administrates {
	return &Administrates{ChannelId: channelId, AdminId: adminId}
}

func (adm *Administrates) IsAdministrator() bool {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ? AND admin_id = ?", adm.ChannelId, adm.AdminId).
		First(&adm); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return true
		}
	}
	return false
}

func (adm *Administrates) IsChannelOwner() bool {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ? AND admin_id = ? AND is_owner = true", adm.ChannelId, adm.AdminId).
		First(&adm); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return adm.IsOwner
		}
	}
	return false
}

func NewAdministratesAdd(channelId, adminId uint) *Administrates {
	return &Administrates{ChannelId: channelId, AdminId: adminId, IsOwner: false}
}

func NewAdministratesAddOwner(channelId, adminId uint) *Administrates {
	return &Administrates{ChannelId: channelId, AdminId: adminId, IsOwner: true}
}

// AddAdministratorOwner Add the channel owner to the Admin table.
func (adm *Administrates) AddAdministratorOwner() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Administrates{}); err != nil {
		return errors.NewDatabaseCreate("administrates")
	}
	if err := db.Debug().
		Table(AdministrateTable).
		Create(&adm).Error; err != nil {
		return err
	}
	// ADD SUBSCRIBER ...
	if err := NewSubscribe(adm.ChannelId, adm.AdminId).AddSubscriber(adm.AdminId); err != nil {
		return err
	}
	return nil
}

// AddAdministrator adds an administrator to a channel.
// Only administrators can add administrators.
func (adm *Administrates) AddAdministrator(addedID uint) error {
	permission := NewAdministratesPermission(adm.ChannelId, adm.AdminId).IsAdministrator()
	if !permission {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}
	db := cockroach.GetDB()

	nad := NewAdministratesAdd(adm.ChannelId, addedID)
	if err := db.Debug().
		Table(AdministrateTable).
		Create(&nad).Error; err != nil {
		return err
	}

	// ADD SUBSCRIBER ...
	if err := NewSubscribe(adm.ChannelId, addedID).AddSubscriber(adm.AdminId); err != nil {
		return err
	}
	return nil
}

func (adm *Administrates) RemoveAdministrator(removedId uint) error {
	isOwner := NewAdministratesPermission(adm.ChannelId, adm.AdminId).IsChannelOwner()
	if !isOwner {
		return errors.New(errors.ErrNotTheOwner)
	}

	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ? AND admin_id = ?", adm.ChannelId, removedId).
		Unscoped().
		Delete(&Administrates{}).Error; err != nil {
		return err
	}
	return nil
}

// GetAdministrators ...
// Returns all administrators of a channel.
// Check permissions.
// Only administrator can view the list of administrators who manage the channel.
func (adm *Administrates) GetAdministrators() ([]*Administrates, error) {
	isAdmin := NewAdministratesAdd(adm.ChannelId, adm.AdminId).IsAdministrator()
	if !isAdmin {
		return nil, errors.New(errors.ErrNotAchannelAdministrator)
	}

	db := cockroach.GetDB()
	var admins []*Administrates
	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ?", adm.ChannelId).
		Find(&admins).Error; err != nil {
		return nil, err
	}

	return admins, nil
}

func (adm *Administrates) ExitAdministrator() error {
	isAdmin := NewAdministratesAdd(adm.ChannelId, adm.AdminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}

	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ? AND admin_id = ?", adm.ChannelId, adm.AdminId).
		Unscoped().
		Delete(&adm).Error; err != nil {
		return err
	}
	return nil
}

// DeleteAdministrators ...
func (adm *Administrates) DeleteAdministrators() error {
	isOwner := NewAdministratesPermission(adm.ChannelId, adm.AdminId).IsChannelOwner()
	if !isOwner {
		return errors.New(errors.ErrNotAchannelOwner)
	}
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(AdministrateTable).
		Where("channel_id = ?", adm.ChannelId).
		Unscoped().
		Delete(&Administrates{}).Error; err != nil {
		return err
	}

	return nil
}
