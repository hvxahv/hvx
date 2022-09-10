package outbox

import (
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
}

func NewFollowOutboxes(actorId uint, activityId, to, types, body string) *Outboxes {
	return &Outboxes{ActorId: actorId, ActivityId: activityId, To: to, Types: types, Body: body}
}

type Outbox interface {
	Create() error
	GetOutboxes() ([]*Outboxes, error)
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
