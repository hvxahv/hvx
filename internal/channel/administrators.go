package channel

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Administrators struct {
	gorm.Model
	CID uint `gorm:"primaryKey;c_id"`
	AID uint `gorm:"primaryKey;a_id"`
}

func (c *Administrators) Add() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("administrators").Where("c_id = ?", c.CID).Where("a_id", c.AID).First(&Administrators{}); err != nil {
		fmt.Println(err.Error)
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.Errorf("ADMINISTRATOR_ALREADY_EXISTS")
		}
	}

	if err := db.Debug().Table("administrators").Create(&c).Error; err != nil {
		return errors.Errorf("FAILED_TO_CREATE_ADMINISTRATOR")
	}

	return nil
}

func (c *Administrators) Remove() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("administrators").Where("c_id = ?", c.CID).Where("a_id = ?", c.AID).First(&Administrators{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return errors.Errorf("the administrator does not exist: %v", err.Error)
		}
	}

	if err := db.Debug().Table("administrators").Where("a_id = ?", c.AID).Unscoped().Delete(&Administrators{}); err != nil {
		return err.Error
	}

	return nil
}

func (c *Administrators) QueryAdmLisByCID() (*[]Administrators, error) {
	db := cockroach.GetDB()

	err := db.Debug().Table("administrators").Where("a_id = ?", c.AID).Where("c_id = ?", c.CID).First(&Channels{})
	if err.Error != nil {
		return nil, errors.Errorf("You are not the administrators of the channel")
	}

	var ch []Administrators
	if err := db.Debug().Table("administrators").Where("c_id = ?", c.CID).Find(&ch); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return nil, errors.Errorf("the administrator does not exist: %v", err.Error)
		}
	}

	return &ch, nil
}

type Admin interface {
	// Add a channel administrators, only the channel owner can operate this method.
	Add() error

	// Remove To delete a channel administrators through this method,
	// only the channel owner can use this method.
	Remove() error

	// QueryAdmLisByCID Fetch the list of administrators through channel id.
	// Only channel administrators and channel owners can use this method.
	QueryAdmLisByCID() (*[]Administrators, error)
}

// NewAddAdmins constructor for a new administrator.
func NewAddAdmins(cid, oid, aid uint) (*Administrators, error) {
	db := cockroach.GetDB()

	owner, err := client.FetchAccountNameByID(oid)
	if err != nil {
		return nil, err
	}

	admin, err := client.FetchAccountNameByID(aid)
	if err != nil {
		return nil, err
	}

	if owner == admin {
		return nil, errors.Errorf("Cannot add yourself as an administrator.")
	}

	fo := db.Debug().Table("channels").Where("id = ?", cid).Where("owner_id = ?", oid).First(&Channels{})
	if fo.Error != nil {
		return nil, errors.Errorf("%s not the owner of the channel.", owner)
	}

	return &Administrators{CID: cid, AID: aid}, nil
}

func NewAdminsByID(cid, aid uint) *Administrators{
	return &Administrators{AID: aid,CID: cid}
}
