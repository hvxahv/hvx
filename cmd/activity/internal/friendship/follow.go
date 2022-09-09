package friendship

import (
	"fmt"

	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"gorm.io/gorm"
)

type Follows struct {
	gorm.Model

	ActorID     uint `gorm:"primaryKey;type:bigint;actor_id;unique"`
	TargetID    uint `gorm:"primaryKey;type:bigint;target_id;unique"`
	IsFollower  bool `gorm:"type:boolean;is_follower"`
	IsFollowing bool `gorm:"type:boolean;is_following"`
	IsFriend    bool `gorm:"type:boolean;is_friend"`
}

type Followee interface {
	Create() error
	UNFollow() error
	GetFollows() ([]uint, error)
}

const (
	// FollowsTableName is the table name for the Follows table.
	FollowsTableName = "follows"
)

func NewGetFollows(actorID uint, followType string) *Follows {
	switch followType {
	case "follower":
		return &Follows{ActorID: actorID, IsFollower: true}
	case "following":
		return &Follows{ActorID: actorID, IsFollowing: true}
	case "friend":
		return &Follows{ActorID: actorID, IsFriend: true}
	}
	return nil
}

func (f *Follows) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Follows{}); err != nil {
		return err
	}

	var field string
	switch {
	case f.IsFollower:
		field = "is_follower"
	case f.IsFollowing:
		field = "is_following"
	}
	if err := db.Debug().
		Table(FollowsTableName).
		Where("actor_id = ? AND target_id = ?", f.ActorID, f.TargetID).
		First(&f); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			if err := db.Debug().
				Table("follows").
				Where("id = ?", f.ID).
				Update(field, true).
				Update("is_friend", true).
				Error; err != nil {
				return err
			}
			return nil
		}
	}

	if err := db.Debug().Table("follows").Create(&f).Error; err != nil {
		return err
	}
	return nil
}

func (f *Follows) UNFollow() error {
	db := cockroach.GetDB()

	var field string
	switch {
	case f.IsFollower:
		field = "is_follower"
	case f.IsFollowing:
		field = "is_following"
	}

	if err := db.Debug().Table(FollowsTableName).
		Where("actor_id = ? AND target_id = ?", f.ActorID, f.TargetID).
		First(&f); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return errors.New("NOT_FOUND")
		}
	}
	if f.IsFollower && f.IsFollowing {
		if err := db.Debug().Table(FollowsTableName).
			Where("actor_id = ? AND target_id = ?", f.ActorID, f.TargetID).
			Update(field, false).
			Update("is_friend", false).
			Error; err != nil {
			return err
		}
	} else {
		if err := db.Debug().Table(FollowsTableName).
			Where("actor_id = ? AND target_id = ?", f.ActorID, f.TargetID).
			Unscoped().
			Delete(&Follows{}).
			Error; err != nil {
			return err
		}
	}

	return nil
}

func NewFollower(actorID uint, targetID uint) *Follows {
	return &Follows{
		ActorID:    actorID,
		TargetID:   targetID,
		IsFollower: true,
	}
}

func NewFollowing(actorID uint, targetID uint) *Follows {
	return &Follows{
		ActorID:     actorID,
		TargetID:    targetID,
		IsFollowing: true,
	}
}

func (f *Follows) GetFollows() ([]uint, error) {
	var field string
	switch {
	case f.IsFollower:
		field = "is_follower"
	case f.IsFollowing:
		field = "is_following"
	case f.IsFriend:
		field = "is_friend"
	}

	db := cockroach.GetDB()
	var followers []uint
	if err := db.Debug().Table(FollowsTableName).
		Where(fmt.Sprintf("actor_id = ? AND %s = ?", field), f.ActorID, true).
		Pluck("target_id", &followers).
		Error; err != nil {
		return nil, err
	}
	return followers, nil
}
