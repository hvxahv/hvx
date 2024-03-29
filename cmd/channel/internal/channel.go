package internal

import (
	"context"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"gorm.io/gorm"
)

const (
	// ErrNotFound is returned when a channel is not found.
	ErrNotFound   = "CHANNEL_NOT_FOUND"
	ChannelsTable = "channels"
)

type Channels struct {
	gorm.Model

	// ActorId Since the Actor of ActivityPub is to be used as the data source of the Channel,
	// the ActorId needs to be saved here.
	ActorId uint `gorm:"primaryKey;type:bigint;actor_id"`

	// CreatorId Store the creator's ActorID in Channel as the owner of the channel.
	CreatorId uint `gorm:"primaryKey;type:bigint;creator_id"`

	// The current Channel design differs from the account design in that there is no privacy in the Channel.
	// So the key will be generated by the server and stored in the database.
	// The public key will be saved in the Actor table when the actor is created
	// And the private key will be saved in the Channel.
	PrivateKey string `gorm:"private_key;type:text;private_key"`
}

type Channel interface {
	// CreateChannel is used to create a channel.
	CreateChannel() error

	// GetChannels is used to retrieve all channels for a creator.
	GetChannels() ([]*Channels, error)

	// DeleteChannel is used to delete a channel.
	DeleteChannel() error

	// DeleteChannels is used to delete all channels.
	DeleteChannels() error

	GetPrivateKeyByActorId() (*Channels, error)
}

// NewChannels channels constructor. The channel can be created by this constructor.
func NewChannels(ActorId, creatorId uint, privateKey string) *Channels {
	return &Channels{ActorId: ActorId, CreatorId: creatorId, PrivateKey: privateKey}
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

	if err := NewAdministratesAddOwner(c.ID, c.CreatorId).AddAdministratorOwner(); err != nil {
		return err
	}
	return nil
}

// NewChannelsCreatorId Constructing the creator ID (actorId).
func NewChannelsCreatorId(creatorId uint) *Channels {
	return &Channels{CreatorId: creatorId}
}

func (c *Channels) GetChannels() ([]*Channels, error) {
	db := cockroach.GetDB()

	var channels []*Channels

	if err := db.Debug().
		Table(ChannelsTable).
		Where("creator_id = ?", c.CreatorId).
		Find(&channels).
		Error; err != nil {
		return nil, err
	}

	return channels, nil
}

// NewChannelsDelete Delete the constructor of the method.
func NewChannelsDelete(channelId, creatorId uint) *Channels {
	return &Channels{
		Model: gorm.Model{
			ID: channelId,
		},
		CreatorId: creatorId,
	}
}

func (c *Channels) DeleteChannel() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ChannelsTable).
		Where("id = ? AND creator_id = ?", c.ID, c.CreatorId).
		First(&c).
		Unscoped().
		Delete(&Channels{}).
		Error; err != nil {
		return err
	}

	// DELETE CHANNEL (IS ACTIVITY PUB ACTOR SERVICE)
	ctx := context.Background()
	d, err := clientv1.New(ctx, microsvc.ActorServiceName).DeleteActor(int64(c.ActorId))
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}

	if d.Code != "200" {
		return errors.New(errors.ErrDeleteChannelActor)
	}

	// DELETE ALL ADMIN...
	if err := NewAdministratesPermission(c.ID, c.CreatorId).DeleteAdministrators(); err != nil {
		return err
	}

	// TODO - DELETE ALL SUB...
	return nil
}

func (c *Channels) DeleteChannels() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ChannelsTable).
		Where("creator_id = ?", c.CreatorId).
		Unscoped().
		Delete(&Channels{}).
		Error; err != nil {
		return err
	}
	return nil
}

func NewChannelActorId(id uint) *Channels {
	return &Channels{ActorId: id}
}
func (c *Channels) GetPrivateKeyByActorId() (*Channels, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ChannelsTable).
		Where("actor_id = ?", c.ActorId).
		First(&c).
		Error; err != nil {
		return nil, err
	}

	return c, nil
}
