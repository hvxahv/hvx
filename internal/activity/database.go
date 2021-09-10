package activity

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
)

func NewInboxToDB(i *Inbox) error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Inbox{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("activity").Create(&i).Error; err != nil {
		return errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return nil
}
