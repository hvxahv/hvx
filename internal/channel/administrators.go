package channel

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Administrators struct {
	gorm.Model
	ChannelID uint `gorm:"primaryKey;channel_id"`
	AccountID uint `gorm:"primaryKey;account_id"`
}

func (c *Administrators) Add() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("administrators").Where("channel_id = ?", c.ChannelID).Where("account_id", c.AccountID).First(&Administrators{}); err != nil {
		fmt.Println(err.Error)
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.Errorf("ADMINISTRATOR_ALREADY_EXISTS")
		}
	}

	if err := db.Debug().Table("administrators").Create(&c).Error; err != nil {
		return errors.Errorf("FAILED_TO_CREATE_ADMINISTRATOR")
	}

	return nil
}

func (c *Administrators) Remove() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("administrators").Where("channel_id = ?", c.ChannelID).Where("accounts_id = ?", c.AccountID).First(&Administrators{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return errors.Errorf("the administrator does not exist: %v", err.Error)
		}
	}

	if err := db.Debug().Table("administrators").Where("account_id = ?", c.AccountID).Unscoped().Delete(&Administrators{}); err != nil {
		return err.Error
	}

	return nil
}

func (c *Administrators) QueryAdmLisByCID() (*[]Administrators, error) {
	db := cockroach.GetDB()

	err := db.Debug().Table("administrators").Where("account_id = ?", c.AccountID).Where("channel_id = ?", c.ChannelID).First(&Channels{})
	if err.Error != nil {
		return nil, errors.Errorf("You are not the administrators of the channel")
	}

	var ch []Administrators
	if err := db.Debug().Table("administrators").Where("channel_id = ?", c.ChannelID).Find(&ch); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.Errorf("the administrator does not exist: %v", err.Error)
		}
	}

	return &ch, nil
}

type Admin interface {
	// Add a channel administrators, only the channel owner can operate this method.
	Add() error

	// Remove To delete a channel administrators through this method,
	// only the channel owner can use this method.
	Remove() error

	// QueryAdmLisByCID Fetch the list of administrators through channel id.
	// Only channel administrators and channel owners can use this method.
	QueryAdmLisByCID() (*[]Administrators, error)
}

// NewAddAdmins constructor for a new administrator.
func NewAddAdmins(channelId, ownerId, accountId uint) (*Administrators, error) {
	db := cockroach.GetDB()

	owner, err := client.FetchAccountNameByID(ownerId)
	if err != nil {
		return nil, err
	}

	admin, err := client.FetchAccountNameByID(accountId)
	if err != nil {
		return nil, err
	}

	if owner == admin {
		return nil, errors.Errorf("Cannot add yourself as an administrator.")
	}

	fo := db.Debug().Table("channels").Where("id = ?", channelId).Where("owner_id = ?", ownerId).First(&Channels{})
	if fo.Error != nil {
		return nil, errors.Errorf("%s not the owner of the channel.", owner)
	}

	return &Administrators{ChannelID: channelId, AccountID: accountId}, nil
}

func NewAdminsByID(channelId, accountId uint) *Administrators {
	return &Administrators{AccountID: accountId, ChannelID: channelId}
}
