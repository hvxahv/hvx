package channel

import (
	"github.com/disism/hvxahv/internal"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/db"
	"github.com/disism/hvxahv/pkg/security"
	"github.com/pkg/errors"
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

// ChanAdmins channels admin.
type ChanAdmins struct {
	Id    string `gorm:"primaryKey;type:varchar(100);id"`
	Admin string `gorm:"primaryKey;type:varchar(999);admin"`
}


func (c *ChanAdmins) GetChannelListByName() (int, []Channels, error) {
	d := db.GetDB()

	var lis []ChanAdmins
	if err := d.Debug().Table("chan_admins").Where("admin = ?", c.Admin).Find(&lis).Error; err != nil {
		log.Println(err)
		return 500, nil, err
	}

	var chs []Channels
	for _, i := range lis {
		nfc := NewChannelsByID(i.Id)
		ch := nfc.Find()
		chs = append(chs, ch)
	}
	return 200, chs, nil

}

func (c *ChanAdmins) GetChanAdmLisByID() (int, []accounts.Accounts, error) {
	d := db.GetDB()

	var lis []ChanAdmins
	if err := d.Debug().Table("chan_admins").Where("id = ?", c.Id).Find(&lis).Error; err != nil {
		log.Println(err)
		return 500, nil, err
	}

	// The detailed data will be traversed through the acquired channel list.
	var acts []accounts.Accounts
	for _, i := range lis {
		fa := accounts.NewAcctByName(i.Admin)
		ad, err := fa.Find()
		if err != nil {
			log.Println(err)
			return 500, nil, err
		}
		acts = append(acts, *ad)
	}

	return 200, acts, nil
}

func (c *ChanAdmins) AddAdmin() (int, string, error) {
	d := db.GetDB()

	if err := d.AutoMigrate(&ChanAdmins{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
		return 500, internal.ServerError, err
	}

	if err := d.Debug().Table("chan_admins").Create(&c).Error; err != nil {
		return 500, internal.ServerError, err
	}
	return 200, internal.SuccessAddChanAdm, nil
}

func NewChanAdmins(owner, id, admin string) (*ChanAdmins, error) {
	// Check if it is a channel manager.
	dbs := db.GetDB()
	// TODO = BUG.
	if r := dbs.Debug().Table("channels").
		Where("id = ? AND owner = ?", id, owner).First(&Channels{}); r.Error != nil {
		return nil, errors.Errorf("%s not the owner of the channel.", owner)
	}
	return &ChanAdmins{Id: id, Admin: admin}, nil
}

func NewChanAdminsByName(admin string) *ChanAdmins {
	return &ChanAdmins{Admin: admin}
}

func NewChanAdminsByID(id string) *ChanAdmins {
	return &ChanAdmins{Id: id}
}

type ChanAdmin interface {
	// AddAdmin Add one or more managers to your channel.
	AddAdmin() (int, string, error)

	// GetChannelListByName Get the list of channels created and managed by the user.
	// Return an array of channel data.
	GetChannelListByName() (int, []Channels, error)

	// GetChanAdmLisByID Query channel manager by ID.
	// Returns an array of account information.
	GetChanAdmLisByID() (int, []accounts.Accounts, error)
}

type ChanSubs struct {
	Id         string `gorm:"primaryKey;type:varchar(100);id;unique"`
	Subscriber string `gorm:"primaryKey;type:varchar(999);subscriber"`
}

func (c *ChanSubs) GetChanSubByName() {
	panic("implement me")
}

func (c *ChanSubs) GetSubscriberByID() (int, []accounts.Accounts, error) {
	d := db.GetDB()

	var lis []ChanSubs
	if err := d.Debug().Table("chan_subs").Where("id = ?", c.Id).Find(&lis).Error; err != nil {
		log.Println(err)
		return 500, nil, err
	}

	var acts []accounts.Accounts
	for _, i := range lis {
		fa := accounts.NewAcctByName(i.Subscriber)
		ad, err := fa.Find()
		if err != nil {
			log.Println(err)
			return 500, nil, err
		}
		acts = append(acts, *ad)
	}

	return 200, acts, nil
}

func (c *ChanSubs) NewSubscriber() (int, string, error) {
	d := db.GetDB()

	if err := d.AutoMigrate(&ChanSubs{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
		return 500, internal.ServerError, err
	}

	if err := d.Debug().Table("chan_subs").Create(&c).Error; err != nil {
		return 500, internal.ServerError, err
	}
	return 200, internal.SuccessSubscribed, nil
}

func NewChanSub(id string, subscriber string) *ChanSubs {
	return &ChanSubs{Id: id, Subscriber: subscriber}
}

func NewChanSubByID(id string) *ChanSubs {
	return &ChanSubs{Id: id}
}

type ChanSub interface {
	// NewSubscriber channel
	NewSubscriber() (int, string, error)

	// GetSubscriberByID ...
	GetSubscriberByID() (int, []accounts.Accounts, error)

	// GetChanSubByName Get your subscribed channels by username.
	GetChanSubByName()
}
