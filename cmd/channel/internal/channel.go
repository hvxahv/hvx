package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"gorm.io/gorm"
)

const (
	// ErrNotFound is returned when a channel is not found.
	ErrNotFound   = "CHANNEL_NOT_FOUND"
	ChannelsTable = "channels"
)

type Channels struct {
	gorm.Model

	// ActorId This ID is associated with the Actor table, as an
	// ActivityPub Actor service. This ID can be used
	// to retrieve data from the Actor table. as the channel information.
	ActorId uint `gorm:"primaryKey;type:bigint;actor_id"`

	// AccountId This ID is associated with the Account table, as an owner of the channel.
	AccountId  uint   `gorm:"primaryKey;type:bigint;account_id"`
	PrivateKey string `gorm:"private_key;type:text;private_key"`
}

type Channel interface {
	CreateChannel() error
	GetChannels() ([]*Channels, error)
	DeleteChannel() error
	DeleteChannels() error
}

func NewChannels(ActorId, AccountId uint, privateKey string) *Channels {
	return &Channels{ActorId: ActorId, AccountId: AccountId, PrivateKey: privateKey}
}

func (c *Channels) CreateChannel() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Channels{}); err != nil {
		return errors.NewDatabaseCreate(serviceName)
	}

	if err := db.Debug().
		Table(ChannelsTable).
		Create(&c).
		Error; err != nil {
		return err

	}

	return nil
}

// NewChannelsAccountId returns a new Channels with the given AccountId.
func NewChannelsAccountId(accountId uint) *Channels {
	return &Channels{AccountId: accountId}
}

func (c *Channels) GetChannels() ([]*Channels, error) {
	db := cockroach.GetDB()

	var channels []*Channels

	if err := db.Debug().
		Table(ChannelsTable).
		Where("account_id = ?", c.AccountId).
		Find(&channels).
		Error; err != nil {
		return nil, err
	}

	return channels, nil
}

func NewChannelsDelete(actorId, accountId uint) *Channels {
	return &Channels{
		ActorId:   actorId,
		AccountId: accountId,
	}
}

func (c *Channels) DeleteChannel() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ChannelsTable).
		Where("account_id = ? AND actor_id = ?", c.AccountId, c.ActorId).
		Unscoped().
		Delete(&Channels{}).
		Error; err != nil {
		return err
	}

	return nil
}

func (c *Channels) DeleteChannels() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ChannelsTable).
		Where("account_id = ?", c.AccountId).
		Unscoped().
		Delete(&Channels{}).
		Error; err != nil {
		return err
	}
	return nil
}
