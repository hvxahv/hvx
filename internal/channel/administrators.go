package channel

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Administrators struct {
	gorm.Model

	ChannelID uint `gorm:"primaryKey;channel_id"`
	ActorID uint `gorm:"primaryKey;account_id"`
}

func (c *Administrators) Add() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("administrators").Where("channel_id = ?", c.ChannelID).Where("actor_id", c.ActorID).First(&Administrators{}); err != nil {
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

	if err := db.Debug().Table("administrators").Where("channel_id = ?", c.ChannelID).Where("actor_id = ?", c.ActorID).First(&Administrators{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return errors.Errorf("the administrator does not exist: %v", err.Error)
		}
	}

	if err := db.Debug().Table("administrators").Where("actor_id = ?", c.ActorID).Unscoped().Delete(&Administrators{}); err != nil {
		return err.Error
	}

	return nil
}

func (c *Administrators) FindAdmLisByChannelID() (*[]Administrators, error) {
	db := cockroach.GetDB()

	err := db.Debug().Table("administrators").Where("actor_id = ?", c.ActorID).Where("channel_id = ?", c.ChannelID).First(&Channels{})
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

	// FindAdmLisByChannelID Fetch the list of administrators through channel id.
	// Only channel administrators and channel owners can use this method.
	FindAdmLisByChannelID() (*[]Administrators, error)
}

// NewAddAdmins constructor for a new administrator.
func NewAddAdmins(channelID, ownerID, actorID uint) (*Administrators, error) {
	db := cockroach.GetDB()

	if ownerID == actorID {
		return nil, errors.Errorf("Cannot add yourself as an administrator.")
	}

	fo := db.Debug().Table("channels").Where("id = ?", channelID).Where("owner_id = ?", ownerID).First(&Channels{})
	if fo.Error != nil {
		return nil, errors.Errorf("%v not the owner of the channel.", ownerID)
	}

	return &Administrators{ChannelID: channelID, ActorID: actorID}, nil
}

func NewAdminsByID(channelID, actorID uint) *Administrators {
	return &Administrators{ActorID: actorID, ChannelID: channelID}
}
