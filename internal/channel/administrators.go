package channel

import (
	"github.com/disism/hvxahv/internal"
	"github.com/disism/hvxahv/internal/accounts"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"log"
)

type Adm interface {
	// AddAdmin Add one or more managers to your channel.
	AddAdmin() (int, string, error)

	// RemoveAdmin remove administrators to your channel.
	RemoveAdmin() (int, string, error)

	// GetListByName Get the list of channels created and managed by the user.
	// Return an array of channel data.
	GetListByName() (int, []Channels, error)

	// GetAdmLisByID Query channel manager by ID.
	// Returns an array of account information.
	GetAdmLisByID() (int, []accounts.Accounts, error)
}

// NewAddAdmins constructor for a new administrator.
func NewAddAdmins(owner, id, admin string) (*Admins, error) {
	if owner == admin {
		return nil, errors.Errorf("cannot add yourself as an administrator.")
	}

	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Admins{}); err != nil {
		return nil, errors.Errorf("failed to create channel admin database automatically: %s", err)
	}

	// Find owner: check whether the owner is correct by ID.
	fo := db.Debug().Table("channels").Where("owner = ?", owner).Where("id = ?", id).First(&Channels{})
	if fo.Error != nil {
		return nil, errors.Errorf("%s not the owner of the channel.", owner)
	}

	// Find admin: check whether the administrator already exists by ID.
	fa := db.Debug().Table("admins").Where("id = ?", id).Where("admin = ?", admin).First(&Admins{})
	no, err := cockroach.IsNotFound(fa.Error)
	if err != nil {
		log.Printf("admins table database retrieval error: %v", err)
		return nil, errors.Errorf("error inside the server!")
	}
	if !no {
		return nil, errors.Errorf("administrator: %s already exists!", admin)
	}
	return &Admins{Id: id, Admin: admin}, nil
}

func NewAdminsByName(admin string) *Admins {
	return &Admins{Admin: admin}
}

func NewAdminsByID(id string) *Admins {
	return &Admins{Id: id}
}

type Admins struct {
	Id    string `gorm:"primaryKey;type:varchar(100);id"`
	Admin string `gorm:"primaryKey;type:varchar(999);admin"`
}

func (c *Admins) RemoveAdmin() (int, string, error) {
	db := cockroach.GetDB()
	// TODO - Delete channel manager.......
	qa := db.Debug().Table("admins").Where("admin = ?", c.Admin).Find(&Admins{})
	if qa.Error != nil {
		log.Println(qa.Error)
	}

	if err := db.Debug().Table("admins").Where("admin = ?", c.Admin).Unscoped().Delete(&Admins{}).Error; err != nil {
		return 500, "failed to remove administrator", err
	}
	if qa.Error != nil {
		log.Println(qa.Error)
	}
	return 200, "administrator removed successfully.", nil
}

func (c *Admins) GetListByName() (int, []Channels, error) {
	db := cockroach.GetDB()

	var lis []Admins
	if err := db.Debug().Table("admins").Where("admin = ?", c.Admin).Find(&lis).Error; err != nil {
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

func (c *Admins) GetAdmLisByID() (int, []accounts.Accounts, error) {
	db := cockroach.GetDB()

	var lis []Admins
	if err := db.Debug().Table("admins").Where("id = ?", c.Id).Find(&lis).Error; err != nil {
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

func (c *Admins) AddAdmin() (int, string, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("admins").Create(&c).Error; err != nil {
		return 500, internal.ServerError, err
	}
	return 200, internal.SuccessAddChanAdm, nil
}
