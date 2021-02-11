package services

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
	"hvxahv/pkg/activity"
	db2 "hvxahv/pkg/db"
	"hvxahv/pkg/db/mongo"
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
	r := activity.SendActivity(&sa)

	db, err := db2.GetMongo()
	if err != nil {
		log.Println(err)
	}
	go func() {
		o := mongo.MongoSave{
			DB:    db,
			Table: "follower",
		}
		co := o.MongoSaveMethod()
		save(co, in)
	}()
	return r

}

func save(co *mongo2.Collection, in *models.Accept) {
	a := new(models.Follow)
	a.Name = in.Name
	a.Actor = in.Actor
	a.Date = time.Now().UTC().Format(http.TimeFormat)

	insertResult, err := co.InsertOne(context.TODO(), a)
	if err != nil {
		log.Println("insert data error: ", err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}