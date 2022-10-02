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
	"gorm.io/gorm"
)

const (
	ArticleTable = "articles"
)

type Articles struct {
	gorm.Model

	ActorId uint   `gorm:"primaryKey;actor_id"`
	Title   string `gorm:"type:text;title"`
	Summary string `gorm:"type:text;summary"`

	// Article The content of the article or status,
	// it needs to store the text in html format, but please be aware of XSS attacks.
	Article string `gorm:"type:text;article"`

	// Tags When the content is an article (status is false), the tag of the article should be selected.
	//Tags pq.StringArray `gorm:"type:text[];tags"`

	// TODO - AttachmentType The design of the attachment type consistency should be addressed. Attachment marker attachment type.
	// AUDIO / VIDEO / IMAGES / FILE.
	AttachmentType string `gorm:"type:text;attachment_type"`

	// Attachments Compose an array of attached storage addresses.
	//Attachments pq.StringArray `gorm:"type:text[];attachments"`

	// TO This field stores a slice of the recipient's address, which can be multiple addresses or a single address.
	//// If the slice is empty, it is sent to all followers.
	//TO pq.Int64Array `gorm:"type:integer[];to"`
	//
	//// CC This field stores the sliced inbox address of the person being copied.
	//CC pq.Int64Array `gorm:"type:integer[];cc"`
	//
	//BTO      pq.Int64Array `gorm:"type:integer[];bto"`
	//Audience pq.Int64Array `gorm:"type:integer[];audience"`

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
	//Visibility uint `gorm:"type:bigint;visibility"`
}

type Article interface {
	// Create is a method for creating an article.
	Create() (*Articles, error)

	// Get is a method for getting an article.
	Get(actorId uint) (*Articles, error)

	// GetArticles is a method for getting all articles by actor.
	GetArticles() ([]*Articles, error)

	// Update is a method for updating an article.
	Update(articleId, accountId uint) error

	// Delete is a method for deleting an article.
	Delete() error

	// DeleteArticles is a method for deleting all articles by actor.
	DeleteArticles() error
}

type Editor interface {
	// EditTitle is a method for editing the title of an article.
	EditTitle(title string) *Articles

	// EditSummary is a method for editing the summary of an article.
	EditSummary(summary string) *Articles

	// EditArticle is a method for editing the article of an article.
	EditArticle(article string) *Articles

	// EditTags is a method for editing the tags of an article.
	EditTags(tags []string) *Articles

	// EditAttachmentType is a method for editing the attachment article.
	EditAttachmentType(attachmentType string) *Articles

	// EditAttachments is a method for editing the attachments of an article.
	EditAttachments(attachments []string) *Articles

	// EditNSFW is a method for editing the nsfw of an article.
	EditNSFW(nsfw bool) *Articles

	// EditVisibility is a method for editing the visibility of an article.
	EditVisibility(visibility uint) *Articles
}

//
//// NewArticles is a constructor for Articles.
//func NewArticles(
//	actorId uint,
//	title string,
//	summary string,
//	article string,
//	tags []string,
//	attachmentType string,
//	attachments []string,
//	TO []int64,
//	CC []int64,
//	NSFW bool,
//	visibility uint,
//) *Articles {
//	if visibility >= 2 {
//		TO = nil
//		CC = nil
//	}
//	return &Articles{
//		ActorId: actorId,
//		Title:   title,
//		Summary: summary,
//		Article: article,
//		//Tags:           tags,
//		AttachmentType: attachmentType,
//		//Attachments:    attachments,
//		//TO:             TO,
//		//CC:             CC,
//		Statuses:   false,
//		NSFW:       NSFW,
//		Visibility: visibility,
//	}
//}
//
//// NewStatus is a constructor for Status.
//func NewStatus(
//	actorId uint,
//	article string,
//	tags []string,
//	attachmentType string,
//	attachments []string,
//	TO []int64,
//	CC []int64,
//	NSFW bool,
//	visibility uint,
//) *Articles {
//	if visibility >= 2 {
//		TO = nil
//		CC = nil
//	}
//	return &Articles{
//		ActorId:        actorId,
//		Article:        article,
//		Tags:           tags,
//		AttachmentType: attachmentType,
//		Attachments:    attachments,
//		TO:             TO,
//		CC:             CC,
//		Statuses:       true,
//		NSFW:           NSFW,
//		Visibility:     visibility,
//	}
//}

func (a *Articles) Create() (*Articles, error) {
	db := cockroach.GetDB()
	if err := db.AutoMigrate(&Articles{}); err != nil {
		return nil, errors.NewDatabaseCreate(ArticleTable)
	}
	if err := db.Debug().
		Table(ArticleTable).
		Create(&a).
		Error; err != nil {
		return nil, err
	}
	return a, nil
}

//
//// NewArticlesId is a constructor for ArticlesId.
//func NewArticlesId(articleId uint) *Articles {
//	return &Articles{
//		Model: gorm.Model{
//			ID: articleId,
//		},
//	}
//}
//
//func (a *Articles) Get(actorId uint) (*Articles, error) {
//	db := cockroach.GetDB()
//	if err := db.Debug().
//		Table(ArticleTable).
//		Where("id = ?", a.ID).
//		First(&a).Error; err != nil {
//		return nil, err
//	}
//
//	var isExist []int64
//	for _, v := range a.TO {
//		if int64(actorId) == v {
//			isExist = append(isExist, v)
//		}
//	}
//	for _, v := range a.CC {
//		if int64(actorId) == v {
//			isExist = append(isExist, v)
//		}
//	}
//
//	if len(isExist) < 1 && a.ActorId != actorId {
//		return nil, errors.New(errors.ErrNoPermission)
//	}
//	return a, nil
//}
//
//// NewArticlesActorId is a constructor for ArticlesActorId.
//func NewArticlesActorId(actorId uint) *Articles {
//	return &Articles{
//		ActorId: actorId,
//	}
//}
//
//func (a *Articles) GetArticles() ([]*Articles, error) {
//	db := cockroach.GetDB()
//	var articles []*Articles
//	if err := db.Debug().
//		Table(ArticleTable).
//		Where("actor_id = ?", a.ActorId).
//		Find(&articles).
//		Error; err != nil {
//		return nil, err
//	}
//	return articles, nil
//}
//
//func (a *Articles) EditTitle(title string) *Articles {
//	a.Title = title
//	return a
//}
//
//func (a *Articles) EditSummary(summary string) *Articles {
//	a.Summary = summary
//	return a
//}
//
//func (a *Articles) EditArticle(article string) *Articles {
//	a.Article = article
//	return a
//}
//
//func (a *Articles) EditTags(tags []string) *Articles {
//	a.Tags = tags
//	return a
//}
//
//func (a *Articles) EditAttachmentType(attachmentType string) *Articles {
//	a.AttachmentType = attachmentType
//	return a
//}
//
//func (a *Articles) EditAttachments(attachments []string) *Articles {
//	a.Attachments = attachments
//	return a
//}
//
//func (a *Articles) EditNSFW(nsfw bool) *Articles {
//	a.NSFW = nsfw
//	return a
//}
//
//func (a *Articles) EditVisibility(visibility uint) *Articles {
//	a.Visibility = visibility
//	return a
//}
//
//func (a *Articles) Update(articleId, actorId uint) error {
//	db := cockroach.GetDB()
//	if err := db.Debug().
//		Table(ArticleTable).
//		Where("id = ? AND actor_id = ?", articleId, actorId).
//		Updates(a).
//		Error; err != nil {
//		return nil
//	}
//
//	return nil
//}
//
//func NewArticlesDelete(articleId, actorId uint) *Articles {
//	return &Articles{
//		Model: gorm.Model{
//			ID: articleId,
//		},
//		ActorId: actorId,
//	}
//}
//
//func (a *Articles) Delete() error {
//	db := cockroach.GetDB()
//
//	if err := db.Debug().
//		Table(ArticleTable).
//		Where("id = ? AND actor_id = ?", a.ID, a.ActorId).
//		Unscoped().
//		Delete(&Articles{}).
//		Error; err != nil {
//		return err
//	}
//
//	return nil
//}
//
//func (a *Articles) DeleteArticles() error {
//	db := cockroach.GetDB()
//
//	if err := db.Debug().
//		Table(ArticleTable).
//		Where("actor_id = ?", a.ActorId).
//		Unscoped().
//		Delete(&Articles{}).
//		Error; err != nil {
//		return err
//	}
//
//	return nil
//}
