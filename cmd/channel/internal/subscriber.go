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

	ChannelId uint `gorm:"primaryKey;channel_id"`
	ActorId   uint `gorm:"primaryKey;actor_id"`
}

type Subscribe interface {
	IsSubscriber() bool
	AddSubscriber() error
	GetSubscribers(adminId uint) (*[]Subscribes, error)
	RemoveSubscriber(adminId uint) error
	Unsubscribe() error
}

func NewSubscribe(channelId, actorId uint) *Subscribes {
	return &Subscribes{
		ChannelId: channelId,
		ActorId:   actorId,
	}
}

func (sub *Subscribes) IsSubscriber() bool {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND actor_id = ?", sub.ChannelId, sub.ActorId).
		First(&Subscribes{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return true
		}
	}

	return false
}

func (sub *Subscribes) AddSubscriber(adminId uint) error {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}
	is := sub.IsSubscriber()
	if is {
		return errors.New(errors.ErrAlreadySubscribed)
	}

	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Subscribes{}); err != nil {
		return errors.NewDatabaseCreate("subscribes")
	}

	if err := db.Debug().
		Table(SubscribesTable).
		Create(&sub).
		Error; err != nil {
		return err
	}
	return nil
}

func (sub *Subscribes) GetSubscribers(adminId uint) (*[]Subscribes, error) {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return nil, errors.New(errors.ErrNotAchannelAdministrator)
	}

	db := cockroach.GetDB()

	var subs []Subscribes
	if err := db.Debug().
		Table("subscribes").
		Where("channel_id = ?", sub.ChannelId).
		Find(&subs).
		Error; err != nil {
		return nil, err
	}
	return &subs, nil
}

func (sub *Subscribes) RemoveSubscriber(adminId uint) error {
	isAdmin := NewAdministratesPermission(sub.ChannelId, adminId).IsAdministrator()
	if !isAdmin {
		return errors.New(errors.ErrNotAchannelAdministrator)
	}
	is := sub.IsSubscriber()
	if is {
		return errors.New(errors.ErrAlreadySubscribed)
	}

	db := cockroach.GetDB()

	if err := db.Debug().
		Table(SubscribesTable).
		Where("channel_id = ? AND actor_id = ?", sub.ChannelId, sub.ActorId).
		Unscoped().
		Delete(&Subscribes{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (sub *Subscribes) Unsubscribe() error {
	db := cockroach.GetDB()
	isSub := NewSubscribe(sub.ChannelId, sub.ActorId).IsSubscriber()
	if !isSub {
		return errors.New(errors.ErrNotSubscribed)
	}

	if err := db.Debug().
		Table("subscribes").
		Where("channel_id = ? AND actor_id = ?", sub.ChannelId, sub.ActorId).
		Unscoped().
		Delete(&Subscribes{}).
		Error; err != nil {
		return err
	}
	return nil
}
