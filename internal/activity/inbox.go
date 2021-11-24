package activity

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Inboxes struct {
	gorm.Model
	ActivityType string `gorm:"type:text;activity_type"`
	ActorID      uint   `gorm:"type:bigint;actor_id"`
	ObjectID     uint   `gorm:"primaryKey;type:bigint;object_id"`
	SourceID     uint   `gorm:"index;type:bigint;source_id"`
}

func (i *Inboxes) Delete() error {
	panic("implement me")
}

func (i *Inboxes) GetInboxesByObjectID() (*[]Inboxes, error) {
	db := cockroach.GetDB()

	var inboxes []Inboxes
	if err := db.Debug().Table("inboxes").Where("object_id = ?", i.ObjectID).Find(&inboxes).Error; err != nil {
		return nil, errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return &inboxes, nil
}

func (i *Inboxes) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		log.Println(err)
		return err
	}

	if err := db.Debug().Table("inboxes").Create(&i).Error; err != nil {
		return errors.Errorf("an error occurred while creating the activity: %v", err)
	}

	return nil
}

func NewObjectID(id uint) *Inboxes {
	return &Inboxes{ObjectID: id}
}

func NewInboxes(activityType string, actorID uint, objectID uint, sourceID uint) *Inboxes {
	return &Inboxes{ActivityType: activityType, ActorID: actorID, ObjectID: objectID, SourceID: sourceID}
}

type Inbox interface {

	Create() error

	GetInboxesByObjectID() (*[]Inboxes, error)

	Delete() error
}
