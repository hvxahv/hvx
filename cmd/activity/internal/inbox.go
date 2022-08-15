package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"gorm.io/gorm"
)

const (
	InboxTableName = "inboxes"
)

type Inboxes struct {
	gorm.Model

	ReceiverId   uint   `gorm:"primaryKey;type:bigint;receiver_id"`
	SenderAddr   string `gorm:"type:text;sender_addr"`
	ActivityId   string `gorm:"primaryKey;type:text;activity_id"`
	ActivityType string `gorm:"type:text;activity_type"`
	ActivityBody string `gorm:"type:text;activity_body"`
}

func NewInboxes(receiverId uint, senderId, activityId, activityType string, activityBody []byte) *Inboxes {
	return &Inboxes{
		ReceiverId:   receiverId,
		SenderAddr:   senderId,
		ActivityId:   activityId,
		ActivityType: activityType,
		ActivityBody: string(activityBody),
	}
}

type Ibx interface {
	Create() error
	Delete() error
	GetInbox() (*Inboxes, error)
	DeleteInbox() error
	GetInboxes() ([]*Inboxes, error)
}

func (i *Inboxes) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		return err
	}

	if err := db.Debug().
		Table(InboxTableName).
		Create(i).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesActivityId(activityId string) *Inboxes {
	return &Inboxes{ActivityId: activityId}
}

func (i *Inboxes) Delete() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("activity_id = ?", i.ActivityId).
		Unscoped().
		Delete(Inboxes{}).
		Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesIdAndActorId(id, receiverId uint) *Inboxes {
	return &Inboxes{
		Model: gorm.Model{
			ID: id,
		},
		ReceiverId: receiverId,
	}
}

func (i *Inboxes) GetInbox() (*Inboxes, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(InboxTableName).
		Where("id = ? AND receiver_id = ?", i.ID, i.ReceiverId).
		First(&i).
		Error; err != nil {
		return nil, err
	}

	return i, nil
}

func (i *Inboxes) DeleteInbox() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("id = ? AND receiver_id = ?", i.ID, i.ReceiverId).
		Unscoped().
		Delete(Inboxes{}).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesReceiverId(receiverId uint) *Inboxes {
	return &Inboxes{ReceiverId: receiverId}
}

func (i *Inboxes) GetInboxes() ([]*Inboxes, error) {
	db := cockroach.GetDB()
	var inboxes []*Inboxes
	if err := db.Debug().
		Table(InboxTableName).
		Where("receiver_id = ?", i.ReceiverId).
		Find(&inboxes).Error; err != nil {
		return nil, err
	}
	return inboxes, nil
}
