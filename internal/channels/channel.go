package channels

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/accounts"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/security"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
)

type Channels struct {
	gorm.Model
	ActorID       uint   `gorm:"primaryKey;actor_id"`
	Name          string `gorm:"type:text;name"`
	Link          string `gorm:"primaryKey;type:text;link;unique"`
	Avatar        string `gorm:"type:text;avatar"`
	Bio           string `gorm:"type:text;bio"`
	OwnerUsername string `gorm:"type:text;owner_username"`
	OwnerID       uint   `gorm:"primaryKey;owner_id"`
	IsPrivate     bool   `gorm:"type:boolean;is_private"`
	PrivateKey    string `gorm:"type:text;private_key"`
}

func (c *Channels) IsExist() error {
	//TODO implement me
	panic("implement me")
}

func NewChannelsByLink(link string) *Channels {
	return &Channels{Link: link}
}

func (c *Channels) GetActorDataByLink() (*accounts.Actors, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("channels").Where("link = ?", c.Link).First(&c).Error; err != nil {
		return nil, err
	}
	actor, err := accounts.NewActorID(c.ActorID).GetByActorID()
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (c *Channels) GetByLink() (*Channels, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("channels").Where("link = ?", c.Link).First(&c).Error; err != nil {
		return nil, err
	}
	return c, nil
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

func (c *Channels) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("channels").Where("owner_username = ?", c.OwnerUsername).Where("id = ?", c.ID).First(&c).Unscoped().Delete(&Channels{}).Error; err != nil {
		return err
	}

	if err := db.Debug().Table("administrators").Where("channel_id = ?", c.ID).Unscoped().Delete(&Administrators{}).Error; err != nil {
		return err
	}
	if err := accounts.NewActorID(c.ActorID).Delete(); err != nil {
		return err
	}
	return nil
}

func (c *Channels) GetByOwnerID() (*[]Channels, error) {
	var ch []Channels
	db := cockroach.GetDB()
	if err := db.Debug().Table("channels").Where("owner_id = ?", c.OwnerID).Find(&ch).Error; err != nil {
		return nil, errors.Errorf("querying channels by link error: %v", err)
	}
	return &ch, nil
}

func NewChannelsOwnerID(actorID uint) *Channels {
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

	// Channel is an extension of ActivityPub. When subscribing to a channels,
	// ActivityPub is called follow and hvxahv is called subscription.
	// The type of Channel is a service in Activitypub. Details:
	// https://www.w3.org/TR/activitystreams-vocabulary/#actor-types
	domain := viper.GetString("localhost")
	url := fmt.Sprintf("https://%s/channels/%s", domain, c.Link)
	inbox := fmt.Sprintf("https://%s/c/%s/inbox", domain, c.Link)
	acct, err := accounts.NewAddActor(c.Link, domain, c.Avatar, c.Name, c.Bio, inbox, url, publicKey, c.Link, "Service").Create()
	if err != nil {
		return err
	}
	fmt.Println(acct)

	c.PrivateKey = privateKey
	c.ActorID = acct.ID

	if err := db.Debug().Table("channels").Create(&c).Error; err != nil {
		return errors.Errorf("failed to create channels: %v", err)
	}

	if err := db.AutoMigrate(&Administrators{}); err != nil {
		return errors.Errorf("failed to create channels admin database automatically: %s", err)
	}
	adm := &Administrators{
		ChannelID: c.ID,
		ActorID:   c.OwnerID,
	}

	if err := db.Debug().Table("administrators").Create(&adm).Error; err != nil {
		return errors.Errorf("failed to create channels: %v", err)
	}

	return nil
}

// NewDeleteChannel Only managers are allowed to delete Channel.
func NewDeleteChannel(username string, id uint) *Channels {
	return &Channels{
		Model: gorm.Model{
			ID: id,
		},
		OwnerUsername: username,
	}
}

type Channel interface {

	// Create a channels.
	Create() error

	GetActorDataByLink() (*accounts.Actors, error)
	// GetByOwnerID Find the channels created by the actor by Actor ID.
	GetByOwnerID() (*[]Channels, error)

	GetByLink() (*Channels, error)
	// Update channels information.
	Update()

	Delete() error

	DeleteHistory()
	DeleteUserHistory()
	EditCreator()

	IsExist() error
}

func NewChannels(name, link, avatar, bio, username string, isPrivate bool) *Channels {
	// Generated if the set link is empty.
	if isPrivate || link == "" {
		random, err := security.GenerateRandomString(15)
		if err != nil {
			link = uuid.New().String()
		}
		link = random
	}

	acct, err := accounts.NewAccountsUsername(username).GetAccountByUsername()
	if err != nil {
		fmt.Println(err)
	}

	return &Channels{
		Name:          name,
		Link:          link,
		Avatar:        avatar,
		Bio:           bio,
		OwnerUsername: username,
		OwnerID:       acct.ID,
		IsPrivate:     isPrivate,
	}
}

func NewChannelID(channelID uint) *Channels {
	return &Channels{
		Model: gorm.Model{
			ID: channelID,
		},
	}
}
