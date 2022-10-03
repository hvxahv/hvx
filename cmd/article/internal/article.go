/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	ArticleTable = "articles"
)

type Articles struct {
	gorm.Model

	ActorId uint   `gorm:"primaryKey;type:bigint;actor_id"`
	Title   string `gorm:"type:text;title"`
	Summary string `gorm:"type:text;summary"`

	// Article The content of the article or status,
	// it needs to store the text in html format,
	// but please be aware of XSS attacks.
	Article string `gorm:"type:text;article"`

	// Tags When the content is an article (status is false),
	// the tag of the article should be selected.
	Tags pq.StringArray `gorm:"type:text[];tags"`

	// Statuses sets whether the content is an article or a status.
	// Long text in HVXAHV means article, which will have to have title,
	// summary, and tags. Short articles are represented as statuses,
	// and statuses will not allow title summaries and tags to be added.
	Statuses bool `gorm:"type:boolean;statuses"`

	//// NSFW Sets whether the content of a post or status is NSFW.
	NSFW bool `gorm:"type:boolean;nsfw"`

	Attachments pq.StringArray `gorm:"type:text[];attachments"`

	// TO This field stores a slice of the recipient's address,
	// which can be multiple addresses or a single address.
	// If the slice is empty, it is sent to all followers.
	TO pq.StringArray `gorm:"type:text[];to"`

	// CC This field stores the sliced inbox address of the person being copied.
	CC pq.StringArray `gorm:"type:text[];cc"`

	BTO pq.StringArray `gorm:"type:text[];bto"`

	// Audience means the audience sent to the channel,
	// when submitting a post, if it needs to be synced to the channel,
	// the synced channel id will be sent to this field.
	Audience pq.Int64Array `gorm:"type:integer[];audience"`

	// Visibility setting for the article or status content, this field is
	// set to int type.
	// 0 is visible to everyone.
	// 1 is only visible to friends who follow each other.
	// 2 is only visible to yourself.
	Visibility int64 `gorm:"type:bigint;visibility"`
}

type Article interface {
	// Create is a method for creating an article.
	Create() (*Articles, error)

	// Get is a method for getting an article.
	Get() (*Articles, error)

	// GetArticles is a method for getting all articles by actor.
	GetArticles() ([]*Articles, error)

	Edit() error

	// Delete is a method for deleting an article.
	Delete() error

	// DeleteArticles is a method for deleting all articles by actor.
	DeleteArticles() error
}

func (a *Articles) Create() (*Articles, error) {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(ArticleTable).
		Create(&a).
		Error; err != nil {
		return nil, err
	}
	return a, nil
}

// NewArticlesId is a constructor for ArticlesId.
func NewArticlesId(articleId uint) *Articles {
	return &Articles{
		Model: gorm.Model{
			ID: articleId,
		},
	}
}

func (a *Articles) Get() (*Articles, error) {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(ArticleTable).
		Where("id = ?", a.ID).
		First(&a).Error; err != nil {
		return nil, err
	}

	return a, nil
}

// NewArticlesActorId is a constructor for ArticlesActorId.
func NewArticlesActorId(actorId uint) *Articles {
	return &Articles{
		ActorId: actorId,
	}
}

func (a *Articles) GetArticles() ([]*Articles, error) {
	db := cockroach.GetDB()
	var articles []*Articles
	if err := db.Debug().
		Table(ArticleTable).
		Where("actor_id = ?", a.ActorId).
		Find(&articles).
		Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func NewArticlesDelete(articleId, actorId uint) *Articles {
	return &Articles{
		Model: gorm.Model{
			ID: articleId,
		},
		ActorId: actorId,
	}
}

func (a *Articles) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ArticleTable).
		Where("id = ? AND actor_id = ?", a.ID, a.ActorId).
		Unscoped().
		Delete(&Articles{}).
		Error; err != nil {
		return err
	}

	return nil
}

func (a *Articles) DeleteArticles() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ArticleTable).
		Where("actor_id = ?", a.ActorId).
		Unscoped().
		Delete(&Articles{}).
		Error; err != nil {
		return err
	}

	return nil
}
