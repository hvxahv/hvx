package accounts

import (
	"fmt"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Follows struct {
	gorm.Model
	Name       string `gorm:"primaryKey;type:varchar(100);name"`
	TargetName string `gorm:"primaryKey;type:varchar(100);target_name"`
}

func NewFollow(name string, targetName string) Follow {
	return &Follows{Name: name, TargetName: targetName}
}

func NewGetFollow(name string) Follow {
	return &Follows{Name: name}
}

type Follow interface {
	New() error
	Get()
}

func (f *Follows) New() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Follows{}); err != nil {
		log.Printf("failed to automatically create database: %v", err)
	}

	//  TODO - Optimization logic.

	// Check the following relationship, If you have followed,
	// then return to have been followed, if not followed,
	// insert a piece of following data and update the number of followers of the account.

	if r := db.Debug().Table("follows").
		Where("name = ? AND target_name = ?", f.Name, f.TargetName).First(&Accounts{}); r.Error != nil {

		if err := db.Debug().Table("follows").Create(&f).Error; err != nil {
			log.Printf("follow failed: %v", err)
			return err
		}

		if err := db.Debug().Table("accounts").Where("username = ?", f.Name).
			Update("following", gorm.Expr("following + ?", 1)).Error; err != nil {
		}

		return nil
	}

	return errors.Errorf("following already.")
}

func (f *Follows) Get() {
	db := cockroach.GetDB()
	res := Follows{}

	if r := db.Debug().Table("follows").Where("name = ?", f.Name).Find(&res); r.Error != nil {
		log.Printf(r.Error.Error())
	}
	fmt.Println(res.TargetName)
}
