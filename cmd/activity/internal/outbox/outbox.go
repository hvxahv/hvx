package outbox

import (
	"github.com/hvxahv/hvx/cmd/activity/internal/activity"
	"github.com/hvxahv/hvx/cockroach"
	"gorm.io/gorm"
	"strings"
)

const (
	TableOutboxesName = "outboxes"
)

type Outboxes struct {
	gorm.Model

	ActorId    uint   `gorm:"primaryKey;type:bigint;actor_id"`
	ActivityId string `gorm:"primaryKey;type:text;activity_id"`
	To         string `gorm:" type:text;to"`
	Cc         string `gorm:"type:text;cc"`
	Bcc        string `gorm:"type:text;bcc"`
	Bto        string `gorm:"type:text;bto"`
	Audience   string `gorm:"type:text;audience"`
	Types      string `gorm:"type:text;types"`
	Body       string `gorm:"type:text;body"`
	IsPublic   bool   `gorm:"type:boolean;is_public"`
}

func NewOutboxes(actorId uint, activityId, to, types, body string) *Outboxes {
	const private = false
	public := true
	switch types {
	case activity.Follow:
		public = private
	case activity.Accept:
		public = private
	case activity.Undo:
		public = private
	case activity.Reject:
		public = private
	}
	return &Outboxes{
		ActorId:    actorId,
		ActivityId: activityId,
		To:         to,
		Types:      types,
		Body:       body,
		IsPublic:   public,
	}
}

type Outbox interface {
	Create() error
	GetOutboxes() ([]*Outboxes, error)
	GetOutboxesPublic() ([]*Outboxes, error)
	Delete() error
}

func (o *Outboxes) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Outboxes{}); err != nil {
		return err
	}
	if err := db.Debug().
		Table(TableOutboxesName).
		Create(&o).Error; err != nil {
		return err
	}
	return nil
}

func NewOutboxesActorId(actorId uint) *Outboxes {
	return &Outboxes{ActorId: actorId}
}

func (o *Outboxes) GetOutboxes() ([]*Outboxes, error) {
	db := cockroach.GetDB()
	var outboxes []*Outboxes
	if err := db.Debug().
		Table(TableOutboxesName).
		Where("actor_id = ?", o.ActorId).
		Find(&outboxes).Error; err != nil {
		return nil, err
	}
	return outboxes, nil
}

func (o *Outboxes) GetOutboxesPublic() ([]*Outboxes, error) {
	db := cockroach.GetDB()
	var outboxes []*Outboxes
	if err := db.Debug().
		Table(TableOutboxesName).
		Where("actor_id = ? AND is_public = ?", o.ActorId, o.IsPublic).
		Find(&outboxes).Error; err != nil {
		return nil, err
	}
	return outboxes, nil
}

func NewOutboxesActivityId(activityId string) *Outboxes {
	return &Outboxes{
		ActivityId: activityId,
	}
}

func (o *Outboxes) Delete() error {
	db := cockroach.GetDB()
	var outboxes []*Outboxes
	if err := db.Debug().
		Table(TableOutboxesName).
		Where("activity_id = ?", o.ActivityId).
		Unscoped().
		Delete(&outboxes).Error; err != nil {
		return err
	}
	return nil
}

// Stos conv slice to strings.
func Stos(slice []string) string {
	return strings.Join(slice, ", ")
}
