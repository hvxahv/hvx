/*
 * MIT License
 *
 * Copyright (c) 2022 The hvxahv Authors.
 *
 */

package article

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model

	AccountID uint   `gorm:"primaryKey;account_id"`
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
	TO pq.StringArray `gorm:"type:text[];to"`

	// CC This field stores the array of inbox addresses of the copied users.
	CC pq.StringArray `gorm:"type:text[];cc"`

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

func (a *Articles) EditArticleTitle(title string) *Articles {
	a.Title = title
	return a
}

func (a *Articles) EditArticleSummary(summary string) *Articles {
	a.Summary = summary
	return a
}

func (a *Articles) EditArticleArticle(article string) *Articles {
	a.Article = article
	return a
}

func (a *Articles) EditArticleTags(tags []string) *Articles {
	a.Tags = tags
	return a
}

func (a *Articles) EditArticleAttachmentType(attachmentType string) *Articles {
	a.AttachmentType = attachmentType
	return a
}

func (a *Articles) EditArticleAttachments(attachments []string) *Articles {
	a.Attachments = attachments
	return a
}

func (a *Articles) EditArticleNSFW(nsfw bool) *Articles {
	a.NSFW = nsfw
	return a
}

func (a *Articles) EditArticleVisibility(visibility uint) *Articles {
	a.Visibility = visibility
	return a
}

//
//func (a *article) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*pb.CreateArticleResponse, error) {
//	db := cockroach.GetDB()
//
//	aid, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	v, err := strconv.Atoi(in.Visibility)
//	if err != nil {
//		return nil, err
//	}
//
//	var s = false
//	if in.State {
//		s = true
//	}
//
//	articles := &Articles{
//		AccountID:      uint(aid),
//		Title:          in.Title,
//		Summary:        in.Summary,
//		Article:        in.Article,
//		Tags:           in.Tags,
//		AttachmentType: in.AttachmentType,
//		Attachments:    in.Attachments,
//		TO:             in.To,
//		CC:             in.Cc,
//		Statuses:       s,
//		NSFW:           in.Nsfw,
//		Visibility:     uint(v),
//	}
//
//	if err := db.AutoMigrate(&Articles{}); err != nil {
//		return nil, err
//	}
//
//	if err := db.Debug().
//		Table("articles").
//		Create(&articles).Error; err != nil {
//		return nil, err
//	}
//
//	return &pb.CreateArticleResponse{Code: "200", Reply: "ok"}, nil
//}
//
//func (a *article) GetArticle(ctx context.Context, in *pb.GetArticleRequest) (*pb.GetArticleResponse, error) {
//	db := cockroach.GetDB()
//
//	id, err := strconv.Atoi(in.Id)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.Debug().
//		Table("articles").
//		Where("id = ?", uint(id)).
//		First(&a.Articles).
//		Error; err != nil {
//		return nil, err
//	}
//
//	data := &pb.GetArticleResponse{
//		Code: "",
//		Id:   in.Id,
//		Article: &pb.CreateArticleRequest{
//			AccountId:      strconv.Itoa(int(a.Articles.AccountID)),
//			Title:          a.Articles.Title,
//			Summary:        a.Articles.Summary,
//			Article:        a.Articles.Article,
//			Tags:           a.Articles.Tags,
//			AttachmentType: a.Articles.AttachmentType,
//			Attachments:    a.Articles.Attachments,
//			To:             a.Articles.TO,
//			Cc:             a.Articles.CC,
//			State:          a.Articles.Statuses,
//			Nsfw:           a.Articles.NSFW,
//			Visibility:     strconv.Itoa(int(a.Articles.Visibility)),
//		},
//	}
//
//	return &pb.GetArticleResponse{Code: "200", Article: data.Article}, nil
//}
//
//func (a *article) GetArticlesByAccountID(ctx context.Context, in *pb.GetArticlesByAccountIDRequest) (*pb.GetArticlesByAccountIDResponse, error) {
//	db := cockroach.GetDB()
//
//	var articles []Articles
//
//	if err := db.Debug().
//		Table("articles").
//		Where("account_id = ?", in.AccountId).
//		Find(&articles).
//		Error; err != nil {
//		return nil, err
//	}
//
//	var r []*pb.CreateArticleRequest
//	for _, v := range articles {
//		r = append(r, &pb.CreateArticleRequest{
//			Id:             strconv.Itoa(int(v.ID)),
//			AccountId:      strconv.Itoa(int(v.AccountID)),
//			Title:          v.Title,
//			Summary:        v.Summary,
//			Article:        v.Article,
//			Tags:           v.Tags,
//			AttachmentType: v.AttachmentType,
//			Attachments:    v.Attachments,
//			To:             v.TO,
//			Cc:             v.CC,
//			State:          v.Statuses,
//			Nsfw:           v.NSFW,
//			Visibility:     strconv.Itoa(int(v.Visibility)),
//		})
//	}
//
//	return &pb.GetArticlesByAccountIDResponse{Code: "200", Articles: r}, nil
//}
//
//func (a *article) UpdateArticle(ctx context.Context, in *pb.UpdateArticleRequest) (*pb.UpdateArticleResponse, error) {
//	db := cockroach.GetDB()
//
//	id, err := strconv.Atoi(in.Id)
//	if err != nil {
//		return nil, err
//	}
//
//	aid, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	articles := new(Articles)
//	switch {
//	case in.Title != "":
//		articles.EditArticleTitle(in.Title)
//	case in.Summary != "":
//		articles.EditArticleSummary(in.Summary)
//	case in.Article != "":
//		articles.EditArticleArticle(in.Article)
//	case len(in.Tags) != 0:
//		articles.EditArticleTags(in.Tags)
//	case in.AttachmentType != "":
//		articles.EditArticleAttachmentType(in.AttachmentType)
//	case len(in.Attachments) != 0:
//		articles.EditArticleAttachments(in.Attachments)
//	case in.Nsfw:
//		articles.EditArticleNSFW(in.Nsfw)
//	case in.Visibility != "":
//		v, err := strconv.Atoi(in.Visibility)
//		if err != nil {
//			return nil, err
//		}
//		articles.EditArticleVisibility(uint(v))
//	}
//
//	if err := db.Debug().
//		Table("articles").
//		Where("id = ? AND account_id = ?", id, aid).
//		Updates(articles).
//		Error; err != nil {
//		return nil, err
//	}
//
//	return &pb.UpdateArticleResponse{Code: "200", Reply: "ok"}, nil
//}
//
//func (a *article) DeleteArticle(ctx context.Context, in *pb.DeleteArticleRequest) (*pb.DeleteArticleResponse, error) {
//	db := cockroach.GetDB()
//
//	id, err := strconv.Atoi(in.Id)
//	if err != nil {
//		return nil, err
//	}
//
//	aid, err := strconv.Atoi(in.AccountId)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := db.Debug().
//		Table("articles").
//		Where("id = ? AND account_id = ?", id, aid).
//		Unscoped().
//		Delete(&Articles{}).
//		Error; err != nil {
//		return nil, err
//	}
//
//	return &pb.DeleteArticleResponse{Code: "200", Reply: "ok"}, nil
//}
