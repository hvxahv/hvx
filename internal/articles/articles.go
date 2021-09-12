// articles The article function is used for users to create an article or status.
// The published article or status is only visible to your friends (people who follow each other).
// It uses the activityPub protocol. You can delete or modify it,
// and your friends will be notified after publishing.

package articles

import (
	"github.com/disism/hvxahv/pkg/cockroach"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type Articles struct {
	gorm.Model
	AuthorID uint   `gorm:"primaryKey;author_id"`
	Title    string `gorm:"type:varchar(600);title"`
	Article  string `gorm:"type:varchar(3000);article"`
}

func (a *Articles) New() error {
	db := cockroach.GetDB()

	if err := db.AutoMigrate(&Articles{}); err != nil {
		return errors.Errorf("failed to automatically create database: %v", err)
	}

	if err := db.Debug().Table("articles").Create(&a); err != nil {
		return errors.Errorf("failed to create article: %v", err)
	}

	return nil
}

func (a *Articles) Update() {
	panic("implement me")
}

func (a *Articles) Delete() {
	panic("implement me")
}

func (a *Articles) FetchArticleByID() {
	panic("implement me")
}

func (a *Articles) FetchArticlesByName() {
	panic("implement me")
}

type Article interface {
	// New Create an article or status.
	New() error

	// Update your article or status.
	Update()

	// Delete your article or status.
	Delete()

	// FetchArticleByID Get article or status by ID.
	FetchArticleByID()

	// FetchLisByName Get the article or status by the username.
	// The article retrieved by this method is a collection instead of a certain article.
	// This collection will return all the articles under this user.
	// You need to set the number of articles obtained by default.
	FetchLisByName()
}

func NewArticles(name, title, article string) *Articles {
	return &Articles{
		AuthorID: client.FetchAccountIdByName(name),
		Title:    title,
		Article:  article,
	}
}