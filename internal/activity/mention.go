package activity

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Mentions struct {
	gorm.Model
	ActivityId string `gorm:"activity_id"`
	ActorID    uint   `gorm:"type:bigint;actor_id"`
	ObjectID   uint   `gorm:"type:bigint;object_id"`
	ArticleID  uint   `gorm:"type:bigint;article_id"`
}

func (m *Mentions) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Mentions{}); err != nil {
		return err
	}

	if err := db.Debug().Table("mentions").Create(&m).Error; err != nil {
		return err
	}

	if err := NewInboxes("Mention", m.ActorID, m.ObjectID, m.ID).Create(); err != nil {
		return err
	}

	return nil
}

func (m *Mentions) Delete() error {
	panic("implement me")
}

func NewMentions(activityId string, actorID uint, objectID uint, articleID uint) *Mentions {
	return &Mentions{ActivityId: activityId, ActorID: actorID, ObjectID: objectID, ArticleID: articleID}
}

type Mention interface {
	Create() error
	Delete() error
}
