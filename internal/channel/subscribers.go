package channel

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Subscribes struct {
	gorm.Model
	ChannelID       uint   `gorm:"primaryKey;channel_id"`
	Subscriber      string `gorm:"primaryKey;type:varchar(999);subscriber"`
	SubscriberInbox string `gorm:"primaryKey;subscriber_inbox"`
}

func (s *Subscribes) Remove() error {
	panic("implement me")
}

func (s *Subscribes) QueryLisByID() (*[]Subscribes, error) {
	db := cockroach.GetDB()

	var sub []Subscribes
	if err := db.Debug().Table("subscribes").Where("channel_id = ?", s.ChannelID).Find(&sub); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.Errorf("SUBSCRIBERS_NOT_FOUND")
		}
	}
	return &sub, nil
}

func (s *Subscribes) New() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(Subscribes{}); err != nil {
		return err
	}

	if err := db.Debug().Table("subscribes").Where("subscriber = ?", s.Subscriber).First(&Subscribes{}); err != nil {
		if !cockroach.IsNotFound(err.Error) {
			return errors.Errorf("ALREADY SUBSCRIBED.")
		}
	}

	if err := db.Debug().Table("subscribes").Create(&s).Error; err != nil {
		return errors.Errorf("SUBSCRIPTION_FAILED")
	}

	return nil
}

type Subscriber interface {
	// New Add a channel subscription.
	New() error

	// Remove subscribers from a channel.
	// This method only allows channel managers to operate.
	Remove() error

	// QueryLisByID Get the list of subscribers of the channel,
	// this method only allows the administrator of the channel to operate
	QueryLisByID() (*[]Subscribes, error)
}

func NewSubscribes(channelId uint, subscriber, subscriberInbox string) (*Subscribes, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("channels").Where("id = ?", channelId).First(&Channels{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.Errorf("CHANNEL_DOESN'T_EXIST")
		}
	}

	return &Subscribes{ChannelID: channelId, Subscriber: subscriber, SubscriberInbox: subscriberInbox}, nil
}

func NewSubLisByID(channelId, accountId uint) (*Subscribes, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("administrators").Where("channel_id = ?", channelId).Where("accounts_id = ?", accountId).First(&Channels{}); err.Error != nil {
		return nil, errors.Errorf("YOU ARE NOT THE MODERATOR OF THE CHANNEL")
	}

	return &Subscribes{ChannelID: channelId}, nil
}
