package channel

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/microservices/client"
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

func (c *Channels) Update() {
	panic("implement me")
}

func (c *Channels) QueryByLink() (*Channels, error) {
	db := cockroach.GetDB()

	var ch Channels
	if err := db.Debug().Table("channels").Where("link = ?", c.Link).First(&ch).Error; err != nil {
		return nil, errors.Errorf("querying channel by link error: %v", err)
	}
	return &ch, nil
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
		AccountID: c.OwnerID,
	}

	if err := db.Debug().Table("administrators").Create(&adm).Error; err != nil {
		return errors.Errorf("failed to create channel: %v", err)
	}
	return nil
}

type Channel interface {

	// New Create a channel.
	New() error

	// QueryByLink Get the detailed content of the channel through the link of the channel.
	QueryByLink() (*Channels, error)

	// Update channel information.
	Update()
}

func NewChannelsByLink(link string) *Channels {
	return &Channels{Link: link}
}

func NewChannels(name, link, avatar, bio, owner string, isPrivate bool) *Channels {
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
		OwnerID:   client.FetchAccountIdByName(owner),
		IsPrivate: isPrivate,
	}
}
