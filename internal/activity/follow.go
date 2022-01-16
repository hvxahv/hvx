package activity

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/hvxahv/hvxahv/internal/account"
	"github.com/hvxahv/hvxahv/pkg/activitypub"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Follows struct {
	gorm.Model
	ActorID  uint `gorm:"primaryKey;bigint;actor_id"`
	ObjectID uint `gorm:"primaryKey;bigint;object_id"`
}

type FollowRequests struct {
	gorm.Model
	ActivityId string `gorm:"primaryKey;activity_id"`
	ActorID    uint   `gorm:"type:bigint;actor_id"`
	ObjectID   uint   `gorm:"type:bigint;object_id"`
}

type FollowAccepts struct {
	gorm.Model
	ActivityId       string `gorm:"activity_id"`
	ActorID          uint   `gorm:"type:bigint;actor_id"`
	ObjectID         uint   `gorm:"type:bigint;object_id"`
	ObjectActivityID string `gorm:"object_activity_id"`
}

func (f *FollowAccepts) Create() error {
	//db := cockroach.GetDB()
	//
	//if err := db.AutoMigrate(&FollowAccepts{}); err != nil {
	//	return err
	//}
	//if err := db.Debug().Table("follow_accepts").Create(&f).Error; err != nil {
	//	return err
	//}
	//
	//if err := NewInboxes("Accept", f.ActorID, f.ObjectID, f.ID).Create(); err != nil {
	//	return err
	//}
	//
	//if err := NewFollows(f.ActorID, f.ObjectID).Create(); err != nil {
	//	return err
	//}

	return nil
}

func (f *FollowRequests) Create() error {
	//db := cockroach.GetDB()
	//
	//if err := db.AutoMigrate(&FollowRequests{}); err != nil {
	//	return err
	//}
	//if err := db.Debug().Table("follow_requests").Create(&f).Error; err != nil {
	//	return err
	//}
	//
	//if err := NewInboxes("Follow", f.ActorID, f.ObjectID, f.ID).Create(); err != nil {
	//	return err
	//}

	return nil
}

func (f *FollowRequests) CreateSend() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&FollowRequests{}); err != nil {
		return err
	}
	if err := db.Debug().Table("follow_requests").Create(&f).Error; err != nil {
		return err
	}

	return nil
}

func (f *FollowRequests) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("follow_requests").Where("activity_id = ?", f.ActivityId).
		First(&f).
		Unscoped().Delete(&FollowRequests{}).Error; err != nil {
		return err
	}

	if err := db.Debug().Table("inboxes").Where("source_id = ?", f.ID).Unscoped().Delete(&Inboxes{}).Error; err != nil {
		return err
	}

	return nil
}

func (f *FollowRequests) GetDetailsByID() (*FollowRequests, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("follow_requests").Where("id = ?", f.ID).First(&f).Error; err != nil {
		return nil, err
	}
	return f, nil
}

func NewFollowRequestsActivityID(activityId string) *FollowRequests {
	return &FollowRequests{ActivityId: activityId}
}

func NewFollowRequests(activityId string, actorID, objectID uint) *FollowRequests {
	return &FollowRequests{ActivityId: activityId, ActorID: actorID, ObjectID: objectID}
}

func NewFollowRequestsByID(id uint) *FollowRequests {
	return &FollowRequests{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func NewFollowAccepts(activityId string, actorID uint, objectID uint, objectActivityID string) *FollowAccepts {
	return &FollowAccepts{ActivityId: activityId, ActorID: actorID, ObjectID: objectID, ObjectActivityID: objectActivityID}
}

// NewFollows actorID is the ID of the person who requested to be followed, and object is the ID of the followed user
func NewFollows(actorID uint, objectID uint) *Follows {
	return &Follows{ActorID: actorID, ObjectID: objectID}
}

func NewByActorID(id uint) *Follows {
	return &Follows{ActorID: id}
}

func (f *Follows) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Follows{}); err != nil {
		return err
	}

	if err := db.Debug().Table("follows").
		Where("actor_id = ? AND object_id = ?", f.ActorID, f.ObjectID).
		First(&Follows{}); err != nil {
		ok := cockroach.IsNotFound(err.Error)
		if !ok {
			return errors.Errorf("ACTOR ALREADY FOLLOWED.")
		}
	}

	if err := db.Debug().Table("follows").Create(&f).Error; err != nil {
		return err
	}

	return nil
}

func (f *Follows) Remove() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("follows").
		Where("actor_id = ? AND object_id = ?", f.ActorID, f.ObjectID).
		Unscoped().Delete(&Follows{}).Error; err != nil {
		return err
	}
	return nil
}

func (f *Follows) GetFollowers() (*[]uint, error) {
	db := cockroach.GetDB()

	var followers []Follows
	if err := db.Debug().Table("follows").Where("actor_id = ?", f.ActorID).Find(&followers).Error; err != nil {
		return nil, err
	}

	var actorID []uint
	for _, i := range followers {
		actorID = append(actorID, i.ObjectID)
	}
	return &actorID, nil
}

func (f *Follows) GetFollowing() (*[]uint, error) {
	db := cockroach.GetDB()

	var following []Follows
	if err := db.Debug().Table("follows").Where("object_id = ?", f.ActorID).Find(&following).Error; err != nil {
		return nil, err
	}

	var actorID []uint
	for _, i := range following {
		actorID = append(actorID, i.ActorID)
	}
	return &actorID, nil
}

type FollowAccept interface {
	Create() error
}

type FollowRequest interface {
	Create() error
	CreateSend() error
	Delete() error

	// GetDetailsByID Get a detailed follow request by ID.
	GetDetailsByID() (*FollowRequests, error)
}

type Follow interface {
	// Create a new follower, the follower of the Actor is Object.
	// Use this method to add a following when receiving the other party's consent to follow request.
	// Use this method to create a follower when replying to agree to the other party's follow request.
	Create() error

	// Remove a follower or following.
	// When the other party sends a request to remove followers.
	Remove() error

	GetFollowers() (*[]uint, error)
	GetFollowing() (*[]uint, error)
}

// NewFoAPData Instantiate the requested follow data in ActivityPub format.
// Return the data and the inbox address of the other party.
func NewFoAPData(actor string, objectID uint) (*activitypub.Follow, string) {
	o, err := account.NewActorID(objectID).GetByActorID()
	if err != nil {
		return nil, ""
	}

	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id  = fmt.Sprintf("https://%s/%s", viper.GetString("localhost"), uuid.New().String())
	)

	return &activitypub.Follow{
		Context: ctx,
		Id:      id,
		Type:    "Follow",
		Actor:   fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), actor),
		Object:  o.Url,
	}, o.Inbox
}

// NewFoAPAccept Instantiate the ActivityPub format data that agrees to follow the request,
// return the data, and the recipient's inbox address.
// Create a follow data to receive the IDs of followers and followers
func NewFoAPAccept(actor, activityID string, objectID uint) (*activitypub.Accept, string) {
	o, err := account.NewActorID(objectID).GetByActorID()
	if err != nil {
		return nil, ""
	}

	var (
		ctx = "https://www.w3.org/ns/activitystreams"
		id  = fmt.Sprintf("https://%s/u/%s#accepts/follows/%s", viper.GetString("localhost"), actor, uuid.New().String())
		a   = fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), actor)
	)

	return &activitypub.Accept{
		Context: ctx,
		Id:      id,
		Type:    "Accept",
		Actor:   a,
		Object: struct {
			Id     string `json:"id"`
			Type   string `json:"type"`
			Actor  string `json:"actor"`
			Object string `json:"object"`
		}{
			Id:     activityID,
			Type:   "Follow",
			Actor:  o.Url,
			Object: a,
		},
	}, o.Inbox
}
