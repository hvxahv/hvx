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

	ActorID  uint `gorm:"primaryKey;bigint;actor_id"`
	TargetID uint `gorm:"primaryKey;bigint;target_id"`
}

func (f *Follows) FetchFollowers() *[]uint {
	db := cockroach.GetDB()

	var fs []Follows
	if err := db.Debug().Table("follows").Where("target_id = ?", f.ActorID).First(&fs); err != nil {
		log.Println(err)
	}
	var c []uint
	for _, i := range fs {
		c = append(c, i.ActorID)
	}

	return &c
}

func (f *Follows) FetchFollowing() *[]uint {
	db := cockroach.GetDB()

	var fs []Follows
	if err := db.Debug().Table("follows").Where("actor_id = ?", f.ActorID).First(&fs); err != nil {
		log.Println(err)
	}
	var c []uint
	for _, i := range fs {
		c = append(c, i.TargetID)
	}

	return &c
}

func NewFetchByID(id uint) *Follows {
	return &Follows{ActorID: id}
}

func (f *Follows) New() error {
	db := cockroach.GetDB()

	// Check the following relationship, If you have followed,
	// then return to have been followed, if not followed,
	// insert a piece of following data and update the number of followers of the account.
	if err := db.AutoMigrate(&Follows{}); err != nil {
		fmt.Printf("failed to automatically create database: %v\n", err)
	}

	if err := db.Debug().Table("follows").Where("actor_id = ? AND target_id = ?", f.ActorID, f.TargetID).First(&Follows{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.Errorf("ACTOR ALREADY FOLLOWED.")
		}
	}

	if err := db.Debug().Table("follows").Create(&f).Error; err != nil {
		return errors.Errorf("failed to follow actor: %v", err)
	}

	return nil
}

func NewFollows(actorID uint, targetID uint) *Follows {
	return &Follows{ActorID: actorID, TargetID: targetID}
}

type Follow interface {
	New() error

	FetchFollowers() *[]uint

	FetchFollowing() *[]uint
}
