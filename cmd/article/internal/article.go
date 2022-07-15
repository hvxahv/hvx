/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package internal

import (
	"github.com/hvxahv/hvx/cockroach"
	"github.com/hvxahv/hvx/errors"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

const (
	ArticleTable = "articles"
)

type Articles struct {
	gorm.Model

	AccountId uint   `gorm:"primaryKey;account_id"`
	Title     string `gorm:"type:text;title"`
	Summary   string `gorm:"type:text;summary"`

	// Article The content of the article or status. It needs to store the
	// text in HTML format, But please be aware of XSS attacks.
	Article string `gorm:"type:text;article"`

	// Tags The tag setting of the article, when statuses is false, then
	// the tag of the article should be selected.
	Tags pq.StringArray `gorm:"type:text[];tags"`

	// TODO - AttachmentType The design of the attachment type consistency
	// should be addressed.
	// AttachmentType Mark the type of attachment,
	// AUDIO / VIDEO / IMAGES / FILE.
	AttachmentType string `gorm:"type:text;attachment_type"`

	// Attachments Form the storage addresses of the attachments into an array.
	Attachments pq.StringArray `gorm:"type:text[];attachments"`

	// TO This field stores the address of the inbox sent to a specific user.
	// Note! When implementing the sending system, if the TO field is nil,
	// it means that it is not set to send to the specified user's inbox.
	// So, it should be sent to the inboxes of all users of that user's
	// Follower by default.
	TO pq.Int64Array `gorm:"type:bigint[];to"`

	// CC This field stores the array of inbox addresses of the copied users.
	CC pq.Int64Array `gorm:"type:bigint[];cc"`

	// Statuses sets whether the content is an article or a status.
	// Long text in HVXAHV means article, which will have to have title,
	// summary, and tags. Short articles are represented as statuses,
	// and statuses will not allow title summaries and tags to be added.
	Statuses bool `gorm:"type:boolean;statuses"`

	// NSFW Sets whether the content of a post or status is NSFW.
	NSFW bool `gorm:"type:boolean;nsfw"`

	// Visibility setting for the article or status content, this field is
	// set to int type.
	// 0 is visible to everyone.
	// 1 is only visible to friends who follow each other.
	// 2 is only visible to yourself.
	Visibility uint `gorm:"type:bigint;visibility"`
}

type Article interface {
	Create() error
	Get(actorId uint) (*Articles, error)
	GetArticles() (*[]Articles, error)
	Update(articleId, accountId uint) error
	Delete() error
	DeleteArticles() error
}

type Editor interface {
	EditTitle(title string) *Articles
	EditSummary(summary string) *Articles
	EditArticle(article string) *Articles
	EditTags(tags []string) *Articles
	EditAttachmentType(attachmentType string) *Articles
	EditAttachments(attachments []string) *Articles
	EditNSFW(nsfw bool) *Articles
	EditVisibility(visibility uint) *Articles
}

// NewArticles is a constructor for Articles.
func NewArticles(
	accountId uint,
	title string,
	summary string,
	article string,
	tags []string,
	TO []int64,
	CC []int64,
	NSFW bool,
	visibility uint,
) *Articles {
	if visibility >= 2 {
		TO = nil
		CC = nil
	}
	return &Articles{
		AccountId:  accountId,
		Title:      title,
		Summary:    summary,
		Article:    article,
		Tags:       tags,
		TO:         TO,
		CC:         CC,
		Statuses:   false,
		NSFW:       NSFW,
		Visibility: visibility,
	}
}

// NewStatus is a constructor for Status.
func NewStatus(
	accountId uint,
	article string,
	tags []string,
	attachmentType string,
	attachments []string,
	TO []int64,
	CC []int64,
	NSFW bool,
	visibility uint,
) *Articles {
	if visibility >= 2 {
		TO = nil
		CC = nil
	}
	return &Articles{
		AccountId:      accountId,
		Article:        article,
		Tags:           tags,
		AttachmentType: attachmentType,
		Attachments:    attachments,
		TO:             TO,
		CC:             CC,
		Statuses:       true,
		NSFW:           NSFW,
		Visibility:     visibility,
	}
}

func (a *Articles) Create() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Articles{}); err != nil {
		return errors.NewDatabaseCreate(ArticleTable)
	}

	if err := db.Debug().
		Table(ArticleTable).
		Create(&a).
		Error; err != nil {
		return err
	}
	return nil
}

func NewGetArticle(articleId uint) *Articles {
	return &Articles{
		Model: gorm.Model{
			ID: articleId,
		},
	}
}

func (a *Articles) Get(actorId uint) (*Articles, error) {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(ArticleTable).
		Where("id = ?", a.ID).
		First(&a).Error; err != nil {
		return nil, err
	}

	// Find out if the actor exists in CC and TO,
	// and return no permission to view if it does not exist.
	var isExist []int64
	for _, v := range a.TO {
		if int64(actorId) == v {
			isExist = append(isExist, v)
		}
	}
	for _, v := range a.CC {
		if int64(actorId) == v {
			isExist = append(isExist, v)
		}
	}

	if len(isExist) < 1 {
		return nil, errors.New(errors.ErrNoPermission)
	}

	return a, nil
}

func NewArticlesAccountId(accountId uint) *Articles {
	return &Articles{
		AccountId: accountId,
	}
}

func (a *Articles) GetArticles() (*[]Articles, error) {
	db := cockroach.GetDB()
	var articles []Articles
	if err := db.Debug().
		Table(ArticleTable).
		Where("account_id = ?", a.AccountId).
		Find(&articles).
		Error; err != nil {
		return nil, err
	}
	return &articles, nil
}

func (a *Articles) EditTitle(title string) *Articles {
	a.Title = title
	return a
}

func (a *Articles) EditSummary(summary string) *Articles {
	a.Summary = summary
	return a
}

func (a *Articles) EditArticle(article string) *Articles {
	a.Article = article
	return a
}

func (a *Articles) EditTags(tags []string) *Articles {
	a.Tags = tags
	return a
}

func (a *Articles) EditAttachmentType(attachmentType string) *Articles {
	a.AttachmentType = attachmentType
	return a
}

func (a *Articles) EditAttachments(attachments []string) *Articles {
	a.Attachments = attachments
	return a
}

func (a *Articles) EditNSFW(nsfw bool) *Articles {
	a.NSFW = nsfw
	return a
}

func (a *Articles) EditVisibility(visibility uint) *Articles {
	a.Visibility = visibility
	return a
}

func (a *Articles) Update(articleId, accountId uint) error {
	db := cockroach.GetDB()
	if err := db.Debug().
		Table(ArticleTable).
		Where("id = ? AND account_id = ?", a.ID, a.AccountId).
		Updates(a).
		Error; err != nil {
		return nil
	}

	return nil
}

func NewArticlesDelete(articleId, accountId uint) *Articles {
	return &Articles{
		Model: gorm.Model{
			ID: articleId,
		},
		AccountId: accountId,
	}
}

func (a *Articles) Delete() error {
	db := cockroach.GetDB()

	if err := db.Debug().
		Table(ArticleTable).
		Where("id = ? AND account_id = ?", a.ID, a.AccountId).
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
		Where("account_id = ?", a.AccountId).
		Unscoped().
		Delete(&Articles{}).
		Error; err != nil {
		return err
	}

	return nil
}
