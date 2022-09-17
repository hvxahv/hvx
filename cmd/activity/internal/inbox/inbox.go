package inbox

import (
	"github.com/hvxahv/hvx/cockroach"
	"gorm.io/gorm"
)

const (
	InboxTableName = "inboxes"
)

type Inboxes struct {
	gorm.Model

	ActorId    uint   `gorm:"primaryKey;type:bigint;actor_id"`
	ActivityId string `gorm:"primaryKey;type:text;activity_id"`
	From       string `gorm:"type:text;sender_addr"`
	Types      string `gorm:"type:text;types"`
	Body       string `gorm:"type:text;body"`
	Viewed     bool   `gorm:"type:boolean;viewed"`
}

func NewInboxes(actorId uint, activityId, from, types, body string) *Inboxes {
	return &Inboxes{
		ActorId:    actorId,
		ActivityId: activityId,
		From:       from,
		Types:      types,
		Body:       body,
		Viewed:     false,
	}
}

type Ibx interface {
	Create() error
	DeleteByActivityId() error
	GetInbox() (*Inboxes, error)
	DeleteInbox() error
	GetInboxes() ([]*Inboxes, error)
	SetViewed() error
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

func NewInboxesDeleteByActivityId(actorId uint, activityId string) *Inboxes {
	return &Inboxes{
		ActorId:    actorId,
		ActivityId: activityId,
	}
}

func (i *Inboxes) DeleteByActivityId() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("actor_id = ? AND activity_id = ?", i.ActorId, i.ActivityId).
		Unscoped().
		Delete(Inboxes{}).
		Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesIdAndActorId(id, actorId uint) *Inboxes {
	return &Inboxes{
		Model: gorm.Model{
			ID: id,
		},
		ActorId: actorId,
	}
}

func (i *Inboxes) GetInbox() (*Inboxes, error) {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(InboxTableName).
		Where("id = ? AND actor_id = ?", i.ID, i.ActorId).
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
		Where("id = ? AND actor_id = ?", i.ID, i.ActorId).
		Unscoped().
		Delete(Inboxes{}).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesReceiverId(actorId uint) *Inboxes {
	return &Inboxes{ActorId: actorId}
}

func (i *Inboxes) GetInboxes() ([]*Inboxes, error) {
	db := cockroach.GetDB()
	var inboxes []*Inboxes
	if err := db.Debug().
		Table(InboxTableName).
		Where("actor_id = ?", i.ActorId).
		Find(&inboxes).Error; err != nil {
		return nil, err
	}
	return inboxes, nil
}

func NewSetViewed(actorId, inboxId uint) *Inboxes {
	return &Inboxes{
		Model: gorm.Model{
			ID: inboxId,
		},
		ActorId: actorId,
		Viewed:  true,
	}
}
func (i *Inboxes) SetViewed() error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(InboxTableName).
		Where("actor_id = ? AND id = ?", i.ActorId, i.ID).
		Updates(&i).Error; err != nil {
		return err
	}
	return nil
}
