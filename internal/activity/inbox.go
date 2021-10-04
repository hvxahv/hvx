package activity

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Inboxes struct {
	gorm.Model

	ActorID      uint   `gorm:"type:bigint;actor_id"`
	ActivityType string `gorm:"type:text;activity_type"`
	ActivityID   string `gorm:"index;type:text;activity_id"`
	LocalActorID uint   `gorm:"primaryKey;type:bigint;local_actor_id"`
}

func (i *Inboxes) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("inboxes").Where("activity_id = ?", i.ActivityID).Unscoped().Delete(&Inboxes{}).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesActivityID(id string) *Inboxes {
	return &Inboxes{ActivityID: id}
}

func (i *Inboxes) FindInboxesByActorID() (*[]Inboxes, error) {
	db := cockroach.GetDB()

	var inboxes []Inboxes
	if err := db.Debug().Table("inboxes").Where("account_id = ?", i.LocalActorID).Find(&inboxes).Error; err != nil {
		return nil, errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return &inboxes, nil
}

func (i *Inboxes) New() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("inboxes").Create(&i).Error; err != nil {
		return errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return nil
}

type Inbox interface {
	New() error

	FindInboxesByActorID() (*[]Inboxes, error)

	Delete() error
}

func NewInbox(actorID uint, types string, eventID string, localActorID uint) (*Inboxes, error) {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		return nil, errors.New("FAILED_TO_AUTOMATICALLY_CREATE_INBOX_DATABASE")
	}

	return &Inboxes{ActorID: actorID, ActivityType: types, ActivityID: eventID, LocalActorID: localActorID}, nil
}

func NewInboxAccountID(id uint) *Inboxes {
	return &Inboxes{LocalActorID: id}
}
