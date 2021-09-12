// articles The article function is used for users to create an article or status.
// The published article or status is only visible to your friends (people who follow each other).
// It uses the activityPub protocol. You can delete or modify it,
// and your friends will be notified after publishing.

package articles

import (
	"fmt"
	pb "github.com/disism/hvxahv/api/accounts/v1alpha1"
	"github.com/disism/hvxahv/pkg/microservices/client"
	"golang.org/x/net/context"
	"gorm.io/gorm"
	"log"
)

type Articles struct {
	gorm.Model
	AuthorID uint   `gorm:"primaryKey;author_id"`
	Title    string `gorm:"type:varchar(600);title"`
	Article  string `gorm:"type:varchar(600);article"`
}

func NewArticles(name, title, article string) *Articles {
	cli, conn, err := client.Accounts()
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

	acct, err := cli.Find(context.Background(), &pb.NewAccountByName{
		Username: name,
	})
	if err != nil {
		log.Printf("failed to send message to accounts server: %v", err)
	}
	fmt.Println(acct.Id)
	return &Articles{AuthorID: uint(acct.Id), Title: title, Article: article}
}

func (a *Articles) New() error {
	err := NewArticle(a)
	if err != nil {
		return err
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

	// FetchArticlesByName Get the article or status by the username.
	// The article retrieved by this method is a collection instead of a certain article.
	// This collection will return all the articles under this user.
	// You need to set the number of articles obtained by default.
	FetchArticlesByName()
}
