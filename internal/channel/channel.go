package channel

import (
	"fmt"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/security"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

type Channels struct {
	gorm.Model

	Name          string `gorm:"type:text;name"`
	Link          string `gorm:"primaryKey;type:text;link;unique"`
	Avatar        string `gorm:"type:text;avatar"`
	Bio           string `gorm:"type:text;bio"`
	OwnerUsername string `gorm:"type:text;owner_username"`
	OwnerID       uint   `gorm:"primaryKey;owner_id"`
	IsPrivate     bool   `gorm:"type:boolean;is_private"`
	PrivateKey    string `gorm:"type:text;private_key"`
}

func (c *Channels) DeleteHistory() {
	panic("implement me")
}

func (c *Channels) DeleteUserHistory() {
	panic("implement me")
}

func (c *Channels) EditCreator() {
	panic("implement me")
}

func (c *Channels) Delete() {
	panic("implement me")
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

func (c *Channels) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Channels{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	privateKey, publicKey, err := security.GenRSA()
	if err != nil {
		log.Printf("failed to generate public and private keys: %v", err)
		log.Println(err)
	}

	// Channel is an extension of ActivityPub. When subscribing to a channel,
	// ActivityPub is called follow and hvxahv is called subscription.
	// The type of Channel is a service in Activitypub. Details:
	// https://www.w3.org/TR/activitystreams-vocabulary/#actor-types
	domain := viper.GetString("localhost")
	url := fmt.Sprintf("https://%s/channel/%s", domain, c.Link)
	inbox := fmt.Sprintf("https://%s/u/%s/inbox", domain, c.OwnerUsername)
	acct, err := accounts.NewAddActor(c.Link, domain, c.Avatar, c.Name, c.Bio, inbox, url, publicKey, c.Link, "Service").NewActor()
	if err != nil {
		return err
	}
	fmt.Println(acct)

	c.PrivateKey = privateKey

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

	// Create a channel.
	Create() error

	// FindByActorID Find the channel created by the actor by Actor ID.
	FindByActorID() (*[]Channels, error)

	// Update channel information.
	Update()

	Delete()

	DeleteHistory()
	DeleteUserHistory()
	EditCreator()
}

func NewChannels(name, link, avatar, bio, username string, owner uint, isPrivate bool) *Channels {
	// Generated if the set link is empty.
	if isPrivate || link == "" {
		random, err := security.GenerateRandomString(15)
		if err != nil {
			link = uuid.New().String()
		}
		link = random
	}

	return &Channels{
		Name:          name,
		Link:          link,
		Avatar:        avatar,
		Bio:           bio,
		OwnerUsername: username,
		OwnerID:       owner,
		IsPrivate:     isPrivate,
	}
}
