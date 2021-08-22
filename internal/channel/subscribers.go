package channel

import (
	"github.com/disism/hvxahv/internal"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/db"
	"github.com/pkg/errors"
	"log"
)


type Subscriber interface {
	// New subscribers
	New() (int, string, error)

	// GetSubscriberByID ...
	GetSubscriberByID() (int, []accounts.Accounts, error)

	// GetSubscriberList get your subscribed channels by username.
	GetSubscriberList()
}

type Subscribes struct {
	Id         string `gorm:"primaryKey;type:varchar(100);id;unique"`
	Subscriber string `gorm:"primaryKey;type:varchar(999);subscriber"`
}

func (c *Subscribes) GetSubscriberByID() (int, []accounts.Accounts, error) {
	d := db.GetDB()

	var lis []Subscribes
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

func (c *Subscribes) GetSubscriberList() {
	panic("implement me")
}


func (c *Subscribes) New() (int, string, error) {
	d := db.GetDB()

	if err := d.Debug().Table("subscribes").Create(&c).Error; err != nil {
		return 500, internal.ServerError, err
	}
	return 200, internal.SuccessSubscribed, nil
}

func NewSubscriber(id string, subscriber string) (*Subscribes, error) {
	dbs := db.GetDB()

	// Find own: Determine whether the subscribed is your own channel
	fo := dbs.Debug().Table("channels").Where("owner = ?", subscriber).Where("id = ?", id).First(&Channels{})
	isFO, err := db.IsNotFound(fo.Error)
	if err != nil {
		log.Printf("channels table database retrieval error: %v", err)
		return nil, errors.Errorf("error inside the server!")
	}

	// If it’s yourself channel.
	if !isFO {
		return nil, errors.Errorf("you can't subscribe to yourself channel.")
	}

	if err := dbs.AutoMigrate(&Subscribes{}); err != nil {
		return nil, errors.Errorf("failed to create channel subscribes database automatically: %s", err)
	}

	// Find subscribes: Find out if a subscribes exists.
	fs := dbs.Debug().Table("subscribes").Where("id = ?", id).Where("subscriber = ?", subscriber).First(&Subscribes{})
	noSub, err := db.IsNotFound(fs.Error)
	if err != nil {
		log.Printf("subscribes table database retrieval error: %v", err)
		return nil, errors.Errorf("error inside the server!")
	}
	if !noSub {
		return nil, errors.Errorf("subscriber: %s already exists!", subscriber)
	}

	return &Subscribes{Id: id, Subscriber: subscriber}, nil
}

func NewSubscriberByID(id string) *Subscribes {
	return &Subscribes{Id: id}
}
