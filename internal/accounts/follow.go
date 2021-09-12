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

func (f *Follows) Followers() {
	db := cockroach.GetDB()
	res := Follows{}

	if r := db.Debug().Table("follows").Where("target_name = ?", f.TargetName).Find(&res); r.Error != nil {
		log.Printf(r.Error.Error())
	}

	fmt.Println(res.Name)
}

func (f *Follows) Following() {
	db := cockroach.GetDB()
	res := Follows{}

	if r := db.Debug().Table("follows").Where("name = ?", f.Name).Find(&res); r.Error != nil {
		log.Printf(r.Error.Error())
	}

	fmt.Println(res.TargetName)
}

func (f *Follows) New() error {
	db := cockroach.GetDB()

	// Check the following relationship, If you have followed,
	// then return to have been followed, if not followed,
	// insert a piece of following data and update the number of followers of the account.
	if err := db.AutoMigrate(&Follows{}); err != nil {
		fmt.Printf("failed to automatically create database: %v\n", err)
	}
	err := db.Debug().Table("follows").Where("name = ? AND target_name = ?", f.Name, f.TargetName).First(&Accounts{})
	if err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.Errorf("ACTOR ALREADY FOLLOWED.")
		}
	}

	if err := db.Debug().Table("follows").Create(&f).Error; err != nil {
		return errors.Errorf("failed to follow actor: %v", err)
	}

	if err := db.Debug().Table("accounts").Where("username = ?", f.Name).
		Update("following", gorm.Expr("following + ?", 1)).Error; err != nil {
		return errors.Errorf("failed to increase the number of followers: %v", err)
	}

	return nil
}

type Follow interface {
	New() error
	Following()
	Followers()
}

func NewFollows(name string, targetName string) Follow {
	return &Follows{Name: name, TargetName: targetName}
}

func NewFoByName(name string) Follow {
	return &Follows{Name: name}
}

func NewFoTargetByName(targetName string) Follow {
	return &Follows{TargetName: targetName}
}
