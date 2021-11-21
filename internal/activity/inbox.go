package activity

import (
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"log"
)

type Inboxes struct {
	gorm.Model

	ActivityID       string `gorm:"index;type:text;activity_id"`
	ActivityType     string `gorm:"type:text;activity_type"`
	ActorID          uint   `gorm:"type:bigint;actor_id"`
	LocalActorID     uint   `gorm:"primaryKey;type:bigint;local_actor_id"`
	TargetActivityID string `gorm:"type;type:text;target_activity_id"`
	ReqActor         string `gorm:"type;type:text;req_actor"`

	// If it is true, it is a mention, and there is a pointing ID in the Article.
	Mention bool `gorm:"type:boolean;mention"`

	// If it is true, it is a dialogue, and there is a pointing ID in the Article.
	Reply bool `gorm:"type:boolean;reply"`

	// The ID of the article
	ArticleID uint `gorm:"type:bigint;article_id"`
}

func (i *Inboxes) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("inboxes").Where("activity_id = ?", i.ActivityID).Unscoped().Delete(&Inboxes{}).Error; err != nil {
		return err
	}
	return nil
}

func NewInboxesActivityID(activityID string) *Inboxes {
	return &Inboxes{ActivityID: activityID}
}

func (i *Inboxes) FindInboxesByActorID() (*[]Inboxes, error) {
	db := cockroach.GetDB()

	var inboxes []Inboxes
	if err := db.Debug().Table("inboxes").Where("account_id = ?", i.LocalActorID).Find(&inboxes).Error; err != nil {
		return nil, errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return &inboxes, nil
}

func (i *Inboxes) New() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Inboxes{}); err != nil {
		log.Println(err)
		return err
	}
	if err := db.Debug().Table("inboxes").Create(&i).Error; err != nil {
		return errors.Errorf("an error occurred while creating the activity: %v", err)
	}
	return nil
}

type Inbox interface {
	New() error

	FindInboxesByActorID() (*[]Inboxes, error)

	Delete() error
}

func NewMention(activityID string, activityType string, actorID uint, localActorID uint, targetActivityID string, mention bool, articleID uint) *Inboxes {
	return &Inboxes{
		ActivityID:       activityID,
		ActivityType:     activityType,
		ActorID:          actorID,
		LocalActorID:     localActorID,
		TargetActivityID: targetActivityID,
	}
}

func NewReply(activityID string, activityType string, actorID uint, localActorID uint, targetActivityID string, reply bool, articleID uint) *Inboxes {
	return &Inboxes{
		ActivityID:       activityID,
		ActivityType:     activityType,
		ActorID:          actorID,
		LocalActorID:     localActorID,
		TargetActivityID: targetActivityID,
	}
}

func NewAccept(activityID string, activityType string, actorID uint, localActorID uint, targetActivityID string) *Inboxes {
	return &Inboxes{
		ActivityID:       activityID,
		ActivityType:     activityType,
		ActorID:          actorID,
		LocalActorID:     localActorID,
		TargetActivityID: targetActivityID,
	}
}

func NewFollow(activityID string, activityType string, actorID uint, localActorID uint) *Inboxes {
	return &Inboxes{
		ActivityID:   activityID,
		ActivityType: activityType,
		ActorID:      actorID,
		LocalActorID: localActorID,
	}
}
