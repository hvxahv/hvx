package internal

import (
	"context"
	"github.com/hvxahv/hvx/APIs/v1alpha1/actor"
	"github.com/hvxahv/hvx/clientv1"
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/hvxahv/hvx/microsvc"
	"gorm.io/gorm"
	"strconv"
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

	// CreatorId is the ActorId of the creator.
	CreatorId  uint   `gorm:"primaryKey;type:bigint;creator_id"`
	PrivateKey string `gorm:"private_key;type:text;private_key"`
}

type Channel interface {
	CreateChannel() error
	GetChannels() ([]*Channels, error)
	DeleteChannel() error
	DeleteChannels() error
}

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

	// DELETE CHANNEL DADA (IS ACTIVITY PUB ACTOR SERVICE)
	ctx := context.Background()
	client, err := clientv1.New(ctx, []string{microsvc.NewGRPCAddress("actor")})
	if err != nil {
		return err
	}
	defer client.Close()

	d, err := actor.NewActorClient(client.Conn).Delete(ctx, &actor.DeleteRequest{
		Id: strconv.Itoa(int(c.ActorId)),
	})
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
