package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"

	"gorm.io/gorm"
)

const (
	SubscribesTable = "subscribes"
)

type Subscribes struct {
	gorm.Model

	ChannelId    uint `gorm:"primaryKey;channel_id"`
	SubscriberId uint `gorm:"primaryKey;subscriber_id"`
}

type Subscribe interface {
	IsSubscriber() (bool, error)
	AddSubscriber(adminId uint) error
	GetSubscribers(adminId uint) ([]*Subscribes, error)
	RemoveSubscriber(adminId uint) error
	Subscription() error
	Unsubscribe() error
}

func NewSubscribe(channelId, subscriberId uint) *Subscribes {
	return &Subscribes{
		ChannelId:    channelId,
		SubscriberId: subscriberId,
	}
}

func (sub *Subscribes) IsSubscriber() (bool, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Subscribes{}); err != nil {
		return true, errors.NewDatabaseCreate("subscribes")
	}

	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND subscriber_id = ?", sub.ChannelId, sub.SubscriberId).
		First(sub); err != nil {
		if cockroach.IsNotFound(err.Error) {
			return false, nil
		}
	}
	return true, nil
}

func (sub *Subscribes) AddSubscriber(adminId uint) error {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}
	isSub, err := sub.IsSubscriber()
	if err != nil {
		return err
	}
	if isSub {
		return errors.New(errors.ErrAlreadySubscribed)
	}

	db := cockroach.GetDB()
	if err := db.Debug().
		Table(SubscribesTable).
		Create(&sub).
		Error; err != nil {
		return err
	}
	return nil
}

func NewSubscriberChannelId(channelId uint) *Subscribes {
	return &Subscribes{
		ChannelId: channelId,
	}
}

func (sub *Subscribes) GetSubscribers(adminId uint) ([]*Subscribes, error) {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return nil, errors.New(errors.ErrNotAchannelAdministrator)
	}

	db := cockroach.GetDB()

	var subs []*Subscribes
	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ?", sub.ChannelId).
		Find(&subs).
		Error; err != nil {
		return nil, err
	}
	return subs, nil
}

func (sub *Subscribes) RemoveSubscriber(adminId uint) error {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}
	isSub, err := sub.IsSubscriber()
	if err != nil {
		return err
	}
	if !isSub {
		return errors.New(errors.ErrNotSubscribed)
	}

	db := cockroach.GetDB()

	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND subscriber_id = ?", sub.ChannelId, sub.SubscriberId).
		Unscoped().
		Delete(&Subscribes{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (sub *Subscribes) Subscription() error {
	db := cockroach.GetDB()
	isSub, err := sub.IsSubscriber()
	if err != nil {
		return err
	}
	if isSub {
		return errors.New(errors.ErrAlreadySubscribed)
	}

	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND subscriber_id = ?", sub.ChannelId, sub.SubscriberId).
		Create(&sub).
		Error; err != nil {
		return err
	}
	return nil
}

func (sub *Subscribes) Unsubscribe() error {
	db := cockroach.GetDB()
	isSub, err := sub.IsSubscriber()
	if err != nil {
		return err
	}
	if !isSub {
		return errors.New(errors.ErrNotSubscribed)
	}

	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND subscriber_id = ?", sub.ChannelId, sub.SubscriberId).
		Unscoped().
		Delete(&Subscribes{}).
		Error; err != nil {
		return err
	}
	return nil
}
