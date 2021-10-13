// articles The article function is used for users to create an article or status.
// The published article or status is only visible to your friends (people who follow each other).
// It uses the activityPub protocol. You can delete or modify it,
// and your friends will be notified after publishing.

package articles

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type TO struct {
	TO []string
}

type Attachment struct {
	Attachment []struct {
		Type      string      `json:"type"`
		MediaType string      `json:"mediaType"`
		Url       string      `json:"url"`
		Name      interface{} `json:"name"`
		Blurhash  string      `json:"blurhash"`
		Width     int         `json:"width"`
		Height    int         `json:"height"`
	} `json:"attachment"`
}


type CC struct {
	CC []string
}

type Articles struct {
	gorm.Model

	ActivityID string `gorm:type:text;activity_id`
	AuthorID   uint   `gorm:"primaryKey;author_id"`
	URL        string `gorm:"type:text;url"`
	Title      string `gorm:"type:text;title"`
	Summary    string `gorm:"type:text;summary"`
	Article    string `gorm:"type:text;article"`

	Attachment *Attachment `gorm:"type:jsonb;attachment"`

	TO *TO `gorm:"type:jsonb;to"`
	CC *CC `gorm:"type:jsonb;cc"`

	// Whether the setting is status.
	Statuses bool `gorm:"type:boolean;statuses"`
	NSFW     bool `gorm:"type:boolean;nsfw"`

	// If it is set to the public state, the article data will be combined into data traversal
	// and sent to everyone in the follower list.
	Visibility bool `gorm:"type:boolean;visibility"`
}

func (a *Articles) DeleteByURL() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("articles").Where("url = ?", a.URL).Unscoped().Delete(&Articles{}).Error; err != nil {
		return err
	}
	return nil
}

func NewArticleURL(url string) *Articles {
	return &Articles{URL: url}
}

func (a *Articles) DeleteByAccountID() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("articles").Where("author_id = ?", a.AuthorID).Unscoped().Delete(&Articles{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Articles) New() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Articles{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("articles").Create(&a); err != nil {
		return errors.Errorf("failed to create article: %v", err)
	}

	// Save to timelines.

	return nil
}

func (a *Articles) Update() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("articles").Where("id = ?", a.ID).Updates(&a).Error; err != nil {
		return err
	}
	return nil
}

func (a *Articles) DeleteByID() error {
	db := cockroach.GetDB()

	if err := db.Debug().Table("articles").Where("id = ?", a.ID).Unscoped().Delete(&Articles{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Articles) FindByID() (*Articles, error) {
	db := cockroach.GetDB()

	if err := db.Debug().Table("articles").Where("id", a.ID).First(&a).Error; err != nil {
		return nil, err
	}
	return a, nil
}

func NewArticleID(id uint) *Articles {
	return &Articles{
		Model: gorm.Model{
			ID: id,
		},
	}
}

func (a *Articles) FindByAccountID() (*[]Articles, error) {
	db := cockroach.GetDB()

	var articles []Articles
	if err := db.Debug().Table("articles").Where("author_id", a.AuthorID).Find(&articles).Error; err != nil {
		return nil, err
	}
	return &articles, nil
}

func NewArticlesByAccountID(id uint) *Articles {
	return &Articles{AuthorID: id}
}

type Article interface {
	// New Create an article or status.
	New() error

	// Update your article or status.
	Update() error

	// DeleteByID your article or status.
	DeleteByID() error

	DeleteByURL() error

	// FindByID Get article or status by ID.
	FindByID() (*Articles, error)

	// FindByAccountID Get the article or status by the username.
	// The article retrieved by this method is a collection instead of a certain article.
	// This collection will return all the articles under this user.
	// You need to set the number of articles obtained by default.
	FindByAccountID() (*[]Articles, error)

	// DeleteByAccountID Delete all article data under the account.
	DeleteByAccountID() error
}

func NewArticles(
	authorID uint,
	title string,
	summary string,
	article string,
	statuses bool,
	NSFW bool,
) *Articles {

	return &Articles{
		Model:    gorm.Model{},
		AuthorID: authorID,
		Title:    title,
		Summary:  summary,
		Article:  article,
		Statuses: statuses,
		NSFW:     NSFW,
	}
}

func NewStatus(actorID uint, content string) *Articles {
	to := []string{"https://mas.to/users/hvturingga"}
	return &Articles{
		AuthorID: actorID,
		Article:  content,
		TO:       &TO{TO: to},
		Statuses: true,
	}
}
