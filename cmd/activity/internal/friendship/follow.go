package friendship

import (
	"fmt"

	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"gorm.io/gorm"
)

type Follows struct {
	gorm.Model

	ActorId     uint `gorm:"primaryKey;type:bigint;actor_id;unique"`
	ObjectId    uint `gorm:"primaryKey;type:bigint;object_id;unique"`
	IsFollower  bool `gorm:"type:boolean;is_follower"`
	IsFollowing bool `gorm:"type:boolean;is_following"`
	IsFriend    bool `gorm:"type:boolean;is_friend"`
}

type Followee interface {
	Follow() error
	UNFollow() error
	Get() ([]uint, error)
}

const (
	FollowsTableName = "follows"
	Follower         = "Follower"
	Following        = "Following"
	Friend           = "Friend"
)

func NewFollower(actorID uint, objectId uint) *Follows {
	return &Follows{
		ActorId:    actorID,
		ObjectId:   objectId,
		IsFollower: true,
	}
}

func NewFollowing(actorID uint, objectId uint) *Follows {
	return &Follows{
		ActorId:     actorID,
		ObjectId:    objectId,
		IsFollowing: true,
	}
}

func (f *Follows) Follow() error {
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
		Where("actor_id = ? AND object_id = ?", f.ActorId, f.ObjectId).
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
		Where("actor_id = ? AND object_id = ?", f.ActorId, f.ObjectId).
		First(&f); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if ok {
			return errors.New("NOT_FOUND")
		}
	}
	if f.IsFollower && f.IsFollowing {
		if err := db.Debug().Table(FollowsTableName).
			Where("actor_id = ? AND object_id = ?", f.ActorId, f.ObjectId).
			Update(field, false).
			Update("is_friend", false).
			Error; err != nil {
			return err
		}
	} else {
		if err := db.Debug().Table(FollowsTableName).
			Where("actor_id = ? AND object_id = ?", f.ActorId, f.ObjectId).
			Unscoped().
			Delete(&Follows{}).
			Error; err != nil {
			return err
		}
	}

	return nil
}

func NewFollows(actorId uint, followType string) *Follows {
	switch followType {
	case Follower:
		return &Follows{ActorId: actorId, IsFollower: true}
	case Following:
		return &Follows{ActorId: actorId, IsFollowing: true}
	case Friend:
		return &Follows{ActorId: actorId, IsFriend: true}
	}
	return nil
}

func (f *Follows) Get() ([]uint, error) {
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
		Where(fmt.Sprintf("actor_id = ? AND %s = ?", field), f.ActorId, true).
		Pluck("object_id", &followers).
		Error; err != nil {
		return nil, err
	}
	return followers, nil
}
