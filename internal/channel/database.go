package channel

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
)

// NewChannel create channel function.
func NewChannel(c *Channels) error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Channels{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("channels").Create(&c).Error; err != nil {
		return errors.Errorf("failed to create channel: %v", err)

	}
	return nil
}
