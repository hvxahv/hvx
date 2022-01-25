package channel

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Subscribes struct {
	gorm.Model
	ChannelID              uint   `gorm:"primaryKey;channel_id"`
	Subscriber             uint   `gorm:"type:bigint;subscriber"`
	SubscriberInboxAddress string `gorm:"subscriber_inbox_address"`
}

func (s *Subscribes) Unsubscribe() {
	panic("implement me")
}

func (s *Subscribes) Remove() error {
	panic("implement me")
}

func (s *Subscribes) GetSubscribersByID() (*[]Subscribes, error) {
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

func (s *Subscribes) Create() error {
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
	// Create Add a channel subscription.
	Create() error

	// Remove subscribers from a channel.
	// This method only allows channel managers to operate.
	Remove() error

	// GetSubscribersByID Get the list of subscribers of the channel,
	// this method only allows the administrator of the channel to operate
	GetSubscribersByID() (*[]Subscribes, error)
	Unsubscribe()
}

func NewSubscribes(channelId uint, subscriber uint) (*Subscribes, error) {

	//actor, err := account.NewActorID(subscriber).GetByActorID()
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}
	//
	//return &Subscribes{ChannelID: channelId, Subscriber: subscriber, SubscriberInboxAddress: actor.Inbox}, nil
	return nil, nil
}

func NewGetSubscribersID(accountId, channelId uint) (*Subscribes, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("administrators").Where("channel_id = ?", channelId).Where("accounts_id = ?", accountId).First(&Channels{}); err.Error != nil {
		return nil, errors.Errorf("YOU_ARE_NOT_THE_MODERATOR_OF_THE_CHANNEL")
	}

	return &Subscribes{ChannelID: channelId}, nil
}
