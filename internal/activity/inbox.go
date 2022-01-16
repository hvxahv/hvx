package activity

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Inboxes struct {
	gorm.Model

	AccountID    uint   `gorm:"primaryKey;type:bigint;account_id"`
	FromID       uint   `gorm:"type:bigint;form_id"`
	ActivityType string `gorm:"type:text;activity_type"`
	ActivityID   string `gorm:"primaryKey;type:text;activity_id"`
	Body         string `gorm:"type:text;activity_id"`
}

func (a *Inboxes) GetInboxes() (*[]Inboxes, error) {
	db := cockroach.GetDB()
	ibx := &[]Inboxes{}
	if err := db.Debug().Table("inboxes").Where("account_id = ?", a.AccountID).Find(&ibx).Error; err != nil {
		return nil, err
	}
	return ibx, nil
}

func (a *Inboxes) DeleteByActivityID() error {
	db := cockroach.GetDB()
	if err := db.Debug().Table("inboxes").Where("activity_id = ?", a.ActivityID).Unscoped().Delete(&Inboxes{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Inboxes) GetByActivityID() (*Inboxes, error) {
	db := cockroach.GetDB()
	if err := db.Debug().Table("inboxes").Where("activity_id = ?", a.ActivityID).First(&Inboxes{}).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func (a *Inboxes) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		return err
	}
	if err := db.Debug().Table("inboxes").Create(&a).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesActorID(activityID string) *Inboxes {
	return &Inboxes{ActivityID: activityID}
}

func NewInboxesAccountID(accountID uint) *Inboxes {
	return &Inboxes{AccountID: accountID}
}

func NewInboxes(accountID uint, fromID uint, activityType string, activityID string, body string) *Inboxes {
	return &Inboxes{AccountID: accountID, FromID: fromID, ActivityType: activityType, ActivityID: activityID, Body: body}
}

type Inbox interface {
	Create() error
	GetInboxes() (*[]Inboxes, error)
	GetByActivityID() (*Inboxes, error)
	DeleteByActivityID() error
}
