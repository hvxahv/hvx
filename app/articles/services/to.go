package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/activity"
	db2 "hvxahv/pkg/db"
	"hvxahv/pkg/models"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// SendActivity 将活动组合成 json 发送给用户的关注者
func SendActivity(in *pb.ArticleData) {
	addr := viper.GetString("activitypub")
	idr := strconv.Itoa(rand.Int())
	date := time.Now().UTC().Format(http.TimeFormat)

	articleId := fmt.Sprintf("https://%s/u/%s/article/%s", addr, in.Author, idr)
	authorUrl := fmt.Sprintf("https://%s/u/%s", addr, in.Author)

	activityId := fmt.Sprintf("https://%s/u/%s/%s", addr, in.Author, idr)

	to := activity.GetFollow(in.Author,"follower")
	cc := []string{"https://www.w3.org/ns/activitystreams#Public", "https://mstdn.social/users/hvturingga"}

	obj := gin.H{
		"id": articleId,
		"type": "Note",
		"attributedTo": authorUrl,
		"content": in.Content,
		"published": date,
		"to": to,
		"cc": cc,
	}
	hd := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"type": "Create",
		"id": activityId,
		"actor": authorUrl,
		"object": obj,
		"published": date,
		"to": to,
		"cc": cc,

	}
	log.Println("活动ID ", activityId)
	log.Println("文章 ID", articleId)
	data, err := json.Marshal(hd)
	if err != nil {
		log.Println(err)
	}

	// 将 json 数据持久化到 mongodb
	go func() {
		db := db2.GetMongo()
		co := db.Collection("articles")
		ia, err := co.InsertOne(context.TODO(), &hd)
		if err != nil {
			log.Println("insert data in inbox error: ", err)
		}
		log.Println("Inserted to inbox: ", ia.InsertedID)
	}()

	for _, i := range to {
		eib := fmt.Sprintf("%s/inbox", i)
		method := "POST"

		sa := *models.NewSendActivity(data, eib, method, in.Author, authorUrl, i)
		r := activity.SendActivity(&sa)
		log.Println("-------------发送创建事件到远程服务器--------------->", r)
	}

}


