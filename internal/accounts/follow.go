package accounts

import (
	"github.com/disism/hvxahv/pkg/db"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Followers struct {
	gorm.Model
	Follower  string `gorm:"primaryKey;type:varchar(100);follower"`
	Following string `gorm:"primaryKey;type:varchar(100);following"`
}

func NewFollowers(follower string, following string) Follow {
	return &Followers{Follower: follower, Following: following}
}

type Follow interface {
	New() error
}

func (f *Followers) New() error {
	d := db.GetDB()

	if err := d.AutoMigrate(&Followers{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
	}

	//  TODO - Optimization logic.

	// Check the following relationship, If you have followed,
	// then return to have been followed, if not followed,
	// insert a piece of following data and update the number of followers of the account.
	if r := d.Debug().Table("followers").
		Where("follower = ? AND following = ?", f.Follower, f.Following).First(&Accounts{}); r.Error != nil {


			if err := d.Debug().Table("followers").Create(&f).Error; err != nil {
				log.Printf("follow failed: %v", err)
				return err
			}

			if err := d.Debug().Table("accounts").Where("username = ?", f.Follower).
				Update("following", gorm.Expr("following + ?", 1)).Error; err != nil {

			}
			if err := d.Debug().Table("accounts").Where("username = ?", f.Following).
				Update("follower", gorm.Expr("follower + ?", 1)).Error; err != nil {

			}

			return nil
	}

	return errors.Errorf("following already.")
}
