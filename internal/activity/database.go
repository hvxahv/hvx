package activity

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
)

// NewInboxToDB Create an inbox data.
func NewInboxToDB(i *Inbox) error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Inbox{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("inboxes").Create(&i).Error; err != nil {
		return errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return nil
}

// FetchInboxCollectionByName Get the actor inbox by username.
func FetchInboxCollectionByName(name string) (*[]Inbox, error) {
	db := cockroach.GetDB()

	i := &[]Inbox{}
	if err := db.Debug().Table("inboxes").Where("username = ?", name).Find(&i).Error; err != nil {
		return nil, errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return i, nil
}