// article The article function is used for users to create an article or status.
// The published article or status is only visible to your friends (people who follow each other).
// It uses the activityPub protocol. You can delete or modify it,
// and your friends will be notified after publishing.

package article

import (
	"fmt"
	"github.com/hvxahv/hvxahv/pkg/cockroach"
	"github.com/lib/pq"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model

	AccountID   uint           `gorm:"primaryKey;account_id"`
	AuthorName  string         `gorm:"type:text;author_name"`
	URL         string         `gorm:"type:text;url"`
	Title       string         `gorm:"type:text;title"`
	Summary     string         `gorm:"type:text;summary"`
	Article     string         `gorm:"type:text;article"`
	Attachments pq.Int64Array  `gorm:"type:integer[];attachments"`
	TO          pq.StringArray `gorm:"type:text[];to"`
	CC          pq.StringArray `gorm:"type:text[];cc"`
	Statuses    bool           `gorm:"type:boolean;statuses"`
	NSFW        bool           `gorm:"type:boolean;nsfw"`
	Visibility  bool           `gorm:"type:boolean;visibility"`
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

	if err := db.Debug().Table("articles").Where("author_id = ?", a.AccountID).Unscoped().Delete(&Articles{}).Error; err != nil {
		return err
	}
	return nil
}

func (a *Articles) Create() error {
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
	if err := db.Debug().Table("articles").Where("author_id", a.AccountID).Find(&articles).Error; err != nil {
		return nil, err
	}
	for _, i := range articles {
		fmt.Println(i.TO)
	}
	return &articles, nil
}

func NewArticlesByAccountID(id uint) *Articles {
	return &Articles{AccountID: id}
}

type Article interface {
	// Create an article or status.
	Create() error

	// Update your article or status.
	Update() error

	// DeleteByID your article or status.
	DeleteByID() error

	DeleteByURL() error

	// FindByID Get article or status by ID.
	FindByID() (*Articles, error)

	// FindByAccountID Get the article or status by the username.
	// The article retrieved by this method is a collection instead of a certain article.
	// This collection will return all the article under this user.
	// You need to set the number of article obtained by default.
	FindByAccountID() (*[]Articles, error)

	// DeleteByAccountID Delete all article data under the account.
	DeleteByAccountID() error
}

func NewArticles(accountID uint, name, title, summary, article string, isNSFW bool) *Articles {

	return &Articles{
		AccountID:  accountID,
		AuthorName: name,
		Title:      title,
		Summary:    summary,
		Article:    article,
		Statuses:   false,
		NSFW:       isNSFW,
	}
}

func NewStatus(accountID uint, name, content string, attachmentsID []int64, to []string, cc []string, isNSFW bool) *Articles {

	return &Articles{
		AccountID:   accountID,
		AuthorName:  name,
		Article:     content,
		Attachments: attachmentsID,
		TO:          to,
		CC:          cc,
		Statuses:    true,
		NSFW:        isNSFW,
	}
}
