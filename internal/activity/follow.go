package activity

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
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
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&FollowAccepts{}); err != nil {
		return err
	}
	if err := db.Debug().Table("follow_accepts").Create(&f).Error; err != nil {
		return err
	}

	if err := NewInboxes("Accept", f.ActorID, f.ObjectID, f.ID).Create(); err != nil {
		return err
	}

	return nil
}

func (f *FollowRequests) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&FollowRequests{}); err != nil {
		return err
	}
	if err := db.Debug().Table("follow_requests").Create(&f).Error; err != nil {
		return err
	}

	if err := NewInboxes("Follow", f.ActorID, f.ObjectID, f.ID).Create(); err != nil {
		return err
	}

	return nil
}

func (f *FollowRequests) Delete() error {
	panic("implement me")
}

func NewFollowRequests(activityId string, actorID, objectID uint) *FollowRequests {
	return &FollowRequests{ActivityId: activityId, ActorID: actorID, ObjectID: objectID}
}

func NewFollowAccepts(activityId string, actorID uint, objectID uint, objectActivityID string) *FollowAccepts {
	return &FollowAccepts{ActivityId: activityId, ActorID: actorID, ObjectID: objectID, ObjectActivityID: objectActivityID}
}

// NewFollows actorID is the ID of the person who requested to be followed, and object is the ID of the followed user
func NewFollows(actorID uint, objectID uint) *Follows {
	return &Follows{ActorID: actorID, ObjectID: objectID}
}

func NewObjectByID(id uint) *Follows {
	return &Follows{ObjectID: id}
}

func NewActorByID(id uint) *Follows {
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
	panic("implement me")
}

func (f *Follows) GetFollowers() (*[]uint, error) {
	db := cockroach.GetDB()

	var followers []Follows
	if err := db.Debug().Table("follows").Where("object_id = ?", f.ObjectID).Find(&followers).Error; err != nil {
		return nil, err
	}

	var actorID []uint
	for _, i := range followers {
		actorID = append(actorID, i.ActorID)
	}
	return &actorID, nil
}

func (f *Follows) GetFollowing() (*[]uint, error)  {
	db := cockroach.GetDB()

	var following []Follows
	if err := db.Debug().Table("follows").Where("actor_id = ?", f.ActorID).Find(&following).Error; err != nil {
		return nil, err
	}

	var actorID []uint
	for _, i := range following {
		actorID = append(actorID, i.ObjectID)
	}
	return &actorID, nil
}

type FollowAccept interface {
	Create() error
}

type FollowRequest interface {
	Create() error
	Delete() error
}

type Follow interface {
	Create() error
	Remove() error
	GetFollowers() (*[]uint, error)
	GetFollowing() (*[]uint, error)
}

// NewFollowAccept
// name: LOCAL ACTOR NAME,
// actor: REMOTE ACTOR LINK,
// oid: CONTEXT ID,
// object: LOCAL ACTOR LINK.
//func NewFollowAccept(name, object, activityID string, remoteActorID, localActorID uint) *activitypub.Accept {
//	var (
//		ctx = "https://www.w3.org/ns/activitystreams"
//		id  = fmt.Sprintf("https://%s/u/%s#accepts/follows/%s", viper.GetString("localhost"), name, uuid.New().String())
//	)
//
//	nf := accounts.NewFollows(remoteActorID, localActorID)
//	if err := nf.New(); err != nil {
//		log.Println(err)
//	}
//
//	return &activitypub.Accept{
//		Context: ctx,
//		Id:      id,
//		Type:    "Accept",
//		Actor:   object,
//		Object: struct {
//			Id     string `json:"id"`
//			Type   string `json:"type"`
//			Actor  string `json:"actor"`
//			Object string `json:"object"`
//		}{
//			Id:     activityID,
//			Type:   "Follow",
//			Actor:  "",
//			Object: object,
//		},
//	}
//}
//
//
//func FollowAccept(id uint, name string) {
//	db := cockroach.GetDB()
//	var ibx Inboxes
//	if err := db.Debug().Table("inboxes").Where("id = ?", id).First(&ibx).Error; err != nil {
//		log.Println(err)
//	}
//
//	actor := "https://mas.to/users/hvturingga"
//
//	object := fmt.Sprintf("https://%s/u/%s", viper.GetString("localhost"), name)
//
//
//	na := NewFollowAccept(name, object, ibx.ActivityID,  ibx.ActorID, ibx.LocalActorID)
//
//	data, err := json.Marshal(na)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//
//	nar := NewActivityRequest(object, actor, data, []byte(getPrivk()))
//	nar.Accept()
//}
