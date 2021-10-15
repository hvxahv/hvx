package channel

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Channels struct {
	gorm.Model

	Name      string `gorm:"type:varchar(100);name"`
	Link      string `gorm:"primaryKey;type:varchar(100);link;unique"`
	Avatar    string `gorm:"type:varchar(999);avatar"`
	Bio       string `gorm:"type:varchar(999);bio"`
	OwnerID   uint   `gorm:"primaryKey;owner_id"`
	IsPrivate bool   `gorm:"type:boolean;is_private"`
}

func (c *Channels) FindByActorID() (*[]Channels, error) {
	var ch []Channels
	db := cockroach.GetDB()
	if err := db.Debug().Table("channels").Where("owner_id = ?", c.OwnerID).First(&ch).Error; err != nil {
		return nil, errors.Errorf("querying channel by link error: %v", err)
	}
	return &ch, nil
}

func NewChannelOwnerID(actorID uint) *Channels {
	return &Channels{OwnerID: actorID}
}

func (c *Channels) Update() {
	panic("implement me")
}

func (c *Channels) New() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Channels{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("channels").Create(&c).Error; err != nil {
		return errors.Errorf("failed to create channel: %v", err)
	}

	if err := db.AutoMigrate(&Administrators{}); err != nil {
		return errors.Errorf("failed to create channel admin database automatically: %s", err)
	}
	adm := &Administrators{
		ChannelID: c.ID,
		ActorID:   c.OwnerID,
	}

	if err := db.Debug().Table("administrators").Create(&adm).Error; err != nil {
		return errors.Errorf("failed to create channel: %v", err)
	}
	return nil
}

type Channel interface {

	// New Create a channel.
	New() error

	// FindByActorID Find the channel created by the actor by Actor ID.
	FindByActorID() (*[]Channels, error)

	// Update channel information.
	Update()
}

func NewChannels(name, link, avatar, bio string, owner uint, isPrivate bool) *Channels {
	// Generated if the set link is empty.
	if isPrivate || link == "" {
		random, err := security.GenerateRandomString(15)
		if err != nil {
			link = uuid.New().String()
		}
		link = random
	}

	return &Channels{
		Model:     gorm.Model{},
		Name:      name,
		Link:      link,
		Avatar:    avatar,
		Bio:       bio,
		OwnerID:   owner,
		IsPrivate: isPrivate,
	}
}
