package channel

import (
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/hvxahv/hvxahv/pkg/rsa"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"log"
	"math/big"
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

func (c *Channels) GetActorDataByLink() (*account.Actors, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("channel").Where("link = ?", c.Link).First(&c).Error; err != nil {
		return nil, err
	}
	actor, err := account.NewActorID(c.ActorID).GetByActorID()
	if err != nil {
		return nil, err
	}
	return actor, nil
}

func (c *Channels) GetByLink() (*Channels, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("channel").Where("link = ?", c.Link).First(&c).Error; err != nil {
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

	if err := db.Debug().Table("channel").Where("owner_username = ?", c.OwnerUsername).Where("id = ?", c.ID).First(&c).Unscoped().Delete(&Channels{}).Error; err != nil {
		return err
	}

	if err := db.Debug().Table("administrators").Where("channel_id = ?", c.ID).Unscoped().Delete(&Administrators{}).Error; err != nil {
		return err
	}
	if err := account.NewActorID(c.ActorID).Delete(); err != nil {
		return err
	}
	return nil
}

func (c *Channels) GetByOwnerID() (*[]Channels, error) {
	var ch []Channels
	db := cockroach.GetDB()
	if err := db.Debug().Table("channel").Where("owner_id = ?", c.OwnerID).Find(&ch).Error; err != nil {
		return nil, errors.Errorf("querying channel by link error: %v", err)
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

	privateKey, publicKey, err := rsa.GenRSA()
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
	inbox := fmt.Sprintf("https://%s/c/%s/inbox", domain, c.Link)
	acct, err := account.NewAddActor(c.Link, domain, c.Avatar, c.Name, c.Bio, inbox, url, publicKey, c.Link, "Service").Create()
	if err != nil {
		return err
	}
	fmt.Println(acct)

	c.PrivateKey = privateKey
	c.ActorID = acct.ID

	if err := db.Debug().Table("channel").Create(&c).Error; err != nil {
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

	// Create a channel.
	Create() error

	GetActorDataByLink() (*account.Actors, error)
	// GetByOwnerID Find the channel created by the actor by Actor ID.
	GetByOwnerID() (*[]Channels, error)

	GetByLink() (*Channels, error)
	// Update channel information.
	Update()

	Delete() error

	DeleteHistory()
	DeleteUserHistory()
	EditCreator()

	IsExist() error
}

// randomString Generate a random string and receive a parameter with a set length (INT).
func randomString(n int) (string, error) {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		if err != nil {
			return "", err
		}
		ret[i] = letters[num.Int64()]
	}

	return string(ret), nil
}

func NewChannels(name, link, avatar, bio, username string, isPrivate bool) *Channels {
	// Generated if the set link is empty.
	if isPrivate || link == "" {
		random, err := randomString(15)
		if err != nil {
			link = uuid.New().String()
		}
		link = random
	}

	acct, err := account.NewAccountsUsername(username).GetAccountByUsername()
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
