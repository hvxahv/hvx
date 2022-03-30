package activity

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Follows struct {
	gorm.Model

	// Follower actor id.
	Follower uint `gorm:"primaryKey;bigint;follower"`

	// Following actor id.
	Following uint `gorm:"primaryKey;bigint;following"`
}

type Friends struct {
	gorm.Model

	// Friend actor id.
	Friend uint `gorm:"primaryKey;bigint;friend"`

	// Friend of actor id.
	FriendOf uint `gorm:"primaryKey;bigint;friend_of"`
}

const (
	// FollowsTableName is the table name for the Follows table.
	FollowsTableName = "follows"
)

func (f *Follows) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Follows{}); err != nil {
		return err
	}

	if err := db.Debug().
		Table(FollowsTableName).
		Where("follower = ? AND Following = ?", f.Follower, f.Following); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return errors.New("FOLLOW_ALREADY_EXISTS")
		}
	}

	if err := db.Debug().Table("follows").Create(&f).Error; err != nil {
		return err
	}
	return nil
}

func (f *Follows) Delete() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table(FollowsTableName).
		Where("follower = ? AND Following = ?", f.Follower, f.Following).
		Unscoped().
		Delete(&Follows{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (f *Follows) GetFollowers() ([]uint, error) {
	db := cockroach.GetDB()
	var followers []uint
	if err := db.Debug().Table(FollowsTableName).
		Where("follower = ?", f.Follower).
		Pluck("following", &followers).
		Error; err != nil {
		return nil, err
	}
	return followers, nil
}

func (f *Follows) GetFollowings() ([]uint, error) {
	db := cockroach.GetDB()
	var followings []uint
	if err := db.Debug().Table(FollowsTableName).
		Where("following = ?", f.Following).
		Pluck("follower", &followings).
		Error; err != nil {
		return nil, err
	}
	return followings, nil
}

func NewFollows(follower uint, following uint) *Follows {
	return &Follows{Follower: follower, Following: following}
}

func NewFollower(follower uint) *Follows {
	return &Follows{Follower: follower}
}

func NewFollowing(following uint) *Follows {
	return &Follows{Following: following}
}

type Followee interface {
	Create() error
	Delete() error
	GetFollowers() ([]uint, error)
	GetFollowings() ([]uint, error)
}
