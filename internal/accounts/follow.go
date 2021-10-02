package accounts

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/pkg/errors"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
)

type Follows struct {
	gorm.Model

	ActorID  uint `gorm:"primaryKey;bigint;actor_id"`
	TargetID uint `gorm:"primaryKey;bigint;target_id"`
}

//
//func (f *Follows) Followers() {
//	db := cockroach.GetDB()
//	res := Follows{}
//
//	if r := db.Debug().Table("follows").Where("target_name = ?", f.TargetName).Find(&res); r.Error != nil {
//		log.Printf(r.Error.Error())
//	}
//
//	fmt.Println(res.Name)
//}
//
//func (f *Follows) Following() {
//	db := cockroach.GetDB()
//	res := Follows{}
//
//	if r := db.Debug().Table("follows").Where("name = ?", f.Name).Find(&res); r.Error != nil {
//		log.Printf(r.Error.Error())
//	}
//
//	fmt.Println(res.TargetName)
//}

func (f *Follows) New() error {
	db := cockroach.GetDB()

	// Check the following relationship, If you have followed,
	// then return to have been followed, if not followed,
	// insert a piece of following data and update the number of followers of the account.
	if err := db.AutoMigrate(&Follows{}); err != nil {
		fmt.Printf("failed to automatically create database: %v\n", err)
	}

	err := db.Debug().Table("follows").Where("actor_id = ? AND target_id = ?", f.ActorID, f.TargetID).First(&Accounts{})
	if err != nil {
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

type Follow interface {
	New() error
	Following()
	Followers()
}

func NewFollows(actorID uint, targetID uint) *Follows {

	return &Follows{ActorID: actorID, TargetID: targetID}
}

func NewFollower(actor string, targetID uint) *Follows {
	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	account, err := cli.FindAccountsByUsername(context.Background(), &pb.AccountUsername{Username: actor})
	if err != nil {
		log.Println(err)
	}

	return &Follows{ActorID: uint(account.ActorId), TargetID: targetID}
}