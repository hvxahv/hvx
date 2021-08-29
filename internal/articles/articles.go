// articles The article function is used for users to create an article or status.
// The published article or status is only visible to your friends (people who follow each other).
// It uses the activityPub protocol. You can delete or modify it,
// and your friends will be notified after publishing.

package articles

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type ArticleData struct {
	gorm.Model
	Uuid    string `gorm:"primaryKey;type:varchar(100);uuid"`
	// Name shows the author’s nickname.
	Name    string `gorm:"type:varchar(100);name"`
	Avatar  string `gorm:"type:varchar(100);avatar"`
	// The author shows the address of the user, such as xxx@xxxx.xxx.
	Author  string `gorm:"primaryKey;type:varchar(100);author"`
	Title   string `gorm:""`
	Article string `gorm:""`
}

type Articles interface {
	// New Create an article or status.
	New() error
	// Update your article or status.
	Update()
	// Delete your article or status.
	Delete()
	// GetArticleByID Get article or status by ID.
	GetArticleByID()
	// GetArticlesByName Get the article or status by the user name.
	// The article retrieved by this method is a collection instead of a certain article.
	// This collection will return all the articles under this user.
	// You need to set the number of articles obtained by default.
	GetArticlesByName()
}

func Activity() {
	addr := viper.GetString("activitypub")
	idr := strconv.Itoa(rand.Int())
	date := time.Now().UTC().Format(http.TimeFormat)

	articleId := fmt.Sprintf("gateway://%s/u/%s/article/%s", addr, "HVTURINGGA", idr)
	authorUrl := fmt.Sprintf("gateway://%s/u/%s", addr, "HVTURINGGA")

	activityId := fmt.Sprintf("gateway://%s/u/%s/%s", addr, "HVTURINGGA", idr)

	to := []string{"", "", ""}
	cc := []string{"gateway://www.w3.org/ns/activitystreams#Public"}

	obj := gin.H{
		"id":           articleId,
		"type":         "Note",
		"attributedTo": authorUrl,
		"content":      "content",
		"published":    date,
		"to":           to,
		"cc":           cc,
	}
	hd := gin.H{
		"@context":  "gateway://www.w3.org/ns/activitystreams",
		"type":      "Create",
		"id":        activityId,
		"actor":     authorUrl,
		"object":    obj,
		"published": date,
		"to":        to,
		"cc":        cc,
	}
	fmt.Println(hd)
}
