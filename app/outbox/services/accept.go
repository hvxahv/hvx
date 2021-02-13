package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"hvxahv/pkg/activitypub"
	db2 "hvxahv/pkg/db"
	"hvxahv/pkg/models"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// AcceptHandler ...
func AcceptHandler(in *models.Accept) int {
	domain := viper.GetString("activitypub")

	idr := strconv.Itoa(rand.Int())
	uad := fmt.Sprintf("https://%s/u/%s", domain, in.Name)
	randId := fmt.Sprintf("https://%s/%s", domain, idr)

	obj := map[string]string {
		"id": in.RequestId,
		"type": "Follow",
		"actor": in.Actor,
		"object": uad,
	}
	p := gin.H{
		"@context": "https://www.w3.org/ns/activitystreams",
		"id": randId,
		"type": "Accept",
		"actor": uad,
		"object": obj,
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Println(err)
	}

	eib := fmt.Sprintf("%s/inbox", in.Actor)
	method := "POST"

	sa := *models.NewSendActivity(data, eib, method, in.Name, uad, in.Actor)
	r := activitypub.SendActivity(&sa)

	// 将关注者写到数据库并将关注数 +1
	db := db2.GetMongo()
	co := db.Collection("follower")
	a := new(models.Follow)
	a.Name = in.Name
	a.Actor = in.Actor
	a.Date = time.Now().UTC().Format(http.TimeFormat)

	insertResult, err := co.InsertOne(context.TODO(), a)
	if err != nil {
		log.Println("insert follower error: ", err)
	}

	log.Println("Inserted follower: ", insertResult.InsertedID)
	go func() {
		rdb := db2.GetRDB()
		v, err := redis.Int64(rdb.Do("INCR", fmt.Sprintf("%s-follower", in.Name)))
		if err != nil {
			log.Println("INCR failed:", err)
			return
		}

		log.Println("value:", v)
	}()
	return r

}
