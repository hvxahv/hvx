package app

import (
	"errors"
)

func (s Status) CreateStatus() error {
	db2.AutoMigrate(Status{})
	if err := db2.Debug().Table("status").Create(&s).Error; err != nil {
		return errors.New("Failed to write new status to database... ")
	}
	return nil
}
