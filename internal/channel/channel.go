package channel

import (
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/disism/hvxahv/pkg/security"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
)

type Channels struct {
	gorm.Model
	Name      string `gorm:"type:varchar(100);name"`
	Id        string `gorm:"primaryKey;type:varchar(100);id;unique"`
	Avatar    string `gorm:"type:varchar(999);avatar"`
	Bio       string `gorm:"type:varchar(999);bio"`
	OwnerID   uint   `gorm:"primaryKey;owner_id"`
	IsPrivate bool   `gorm:"type:boolean;is_private"`
}

func (c *Channels) New() error {

	err := NewChannel(c)
	if err != nil {
		return err
	}
	return nil
}

type Channel interface {
	// New  Create a channel and return status code, information, id,  and errors.
	New() error

	// Find channel by ID.
	//Find() Channels

	// U　　pdate channel information.
	//Update()

	//GetMyChanByName()
}

/*func (c *Channels) Find() Channels {
	db := cockroach.GetDB()

	if err := db.Debug().Table("channels").Where("id = ?", c.Id).Find(&c).Error; err != nil {
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
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Channels{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
		return 500, internal.ServerError, "", err
	}

	if err := db.Debug().Table("channels").Create(&c).Error; err != nil {
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


func NewChannelsByID(id string) *Channels {
	return &Channels{Id: id}
}
*/

func NewChannels(name, id, avatar, bio, owner string, isPrivate bool) *Channels {
	if isPrivate || id == "" {
		random, err := security.GenerateRandomString(15)
		if err != nil {
			log.Println(err)
		}
		id = random
	}

	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	acct, err := cli.Find(context.Background(), &pb.NewAccountByName{
		Username: owner,
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}

	return &Channels{Name: name, Id: id, Avatar: avatar, Bio: bio, OwnerID: uint(acct.Id), IsPrivate: isPrivate}
}
