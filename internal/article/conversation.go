package article

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"gorm.io/gorm"
)

type Conversations struct {
	gorm.Model

	ActivityID string `gorm:"type:text;activity_id"`
	ActorID    uint   `gorm:"index;type:bigint;actor_id"`
	ArticleURl string `gorm:"primaryKey;type:text;article_url"`
	Content    string `gorm:"type:text;content"`
	ToActorID  uint   `gorm:"type:bigint;actor_id"`
}

func (c *Conversations) Create() error {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Conversations{}); err != nil {
		return err
	}
	if err := db.Debug().Table("conversations").Create(&c).Error; err != nil {
		return err
	}
	return nil
}

func NewConversations(activityID string, actorID uint, articleURl string, content string, toActorID uint) *Conversations {
	return &Conversations{ActivityID: activityID, ActorID: actorID, ArticleURl: articleURl, Content: content, ToActorID: toActorID}
}

type Conversation interface {
	Create() error
}
