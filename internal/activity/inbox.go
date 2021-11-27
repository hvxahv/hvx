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

func (i *Inboxes) GetInboxesByID() (*[]Inboxes, error) {
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

func NewInboxDetails(a string, id uint) *Inboxes {
	return &Inboxes{ActivityType: a, ObjectID: id}
}

type Inbox interface {
	// Create an inbox data, SourceID corresponds to the specific primary key ID of the table data.
	Create() error

	// GetInboxesByID Get the user's inbox list through the inbox ID.
	GetInboxesByID() (*[]Inboxes, error)

	// Delete a piece of inbox data.
	Delete() error
}
