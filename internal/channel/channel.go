package channel

import (
	"github.com/disism/hvxahv/internal"
	"github.com/disism/hvxahv/pkg/db"
	"github.com/disism/hvxahv/pkg/security"
	"gorm.io/gorm"
	"log"
)

type Channels struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);name"`
	Id        string `gorm:"primaryKey;type:varchar(100);id;unique"`
	Avatar    string `gorm:"type:varchar(999);avatar"`
	Bio       string `gorm:"type:varchar(999);bio"`
	Owner     string `gorm:"primaryKey;type:varchar(100);owner"`
	IsPrivate bool   `gorm:"type:boolean;is_private"`
}

func (c *Channels) Find() Channels {
	d := db.GetDB()

	if err := d.Debug().Table("channels").Where("id = ?", c.Id).Find(&c).Error; err != nil {
		log.Println(err)
	}

	return *c
}

func (c *Channels) GetMyChanByName() {
	panic("implement me")
}

func (c *Channels) Update() {
	panic("implement me")
}

func (c *Channels) New() (int, string, string, error) {
	d := db.GetDB()

	if err := d.AutoMigrate(&Channels{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
		return 500, internal.ServerError, "", err
	}

	if err := d.Debug().Table("channels").Create(&c).Error; err != nil {
		return 500, internal.ServerError, "", err
	}

	return 200, internal.SuccessNewChannel, c.Id, nil
}

type Channel interface {
	// New  Create a channel and return status code, information, id,  and errors.
	New() (int, string, string, error)

	// Find channel by ID.
	Find() Channels
	// Update channel information.
	Update()

	GetMyChanByName()
}

func NewChannels(name, id, avatar, bio, owner string, isPrivate bool) *Channels {
	if isPrivate || id == "" {
		random, err := security.GenerateRandomString(15)
		if err != nil {
			log.Println(err)
		}
		id = random
	}

	return &Channels{Name: name, Id: id, Avatar: avatar, Bio: bio, Owner: owner, IsPrivate: isPrivate}
}

func NewChannelsByID(id string) *Channels {
	return &Channels{Id: id}
}
