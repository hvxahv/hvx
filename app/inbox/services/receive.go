package services

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1"
	"hvxahv/pkg/db"
	"hvxahv/pkg/models"
	"log"
	"net/http"
	"time"
)

// ReceiveInbox 接收 Inbox 数据
func ReceiveInbox(in *pb.InboxData) string {
	i := &models.Inbox{
		Actor: in.Actor,
		RequestId: in.RequestId,
		EventType: in.EventType,
		Name: in.Name,
		Date: time.Now().UTC().Format(http.TimeFormat),
	}

	k := fmt.Sprintf("%s-inbox", in.Name)

	// 将数据存储到缓存，返回 ok ，然后将缓存中的数据用 goroutine 持久化到 MongoDB
	rdb := db.GetRDB()
	v, _ := json.Marshal(&i)
	if _, err := rdb.Do("SETNX", k, v); err != nil {
		log.Printf("Inbox 持久化到缓存失败: %s", err)
	}

	go func() {
		db, err := db.GetMongo()
		if err != nil {
			log.Println(err)
		}
		co := db.Collection("inbox")
		inbox := models.NewInboxStructs(i)
		ir, err := co.InsertOne(context.TODO(), &inbox)
		if err != nil {
			log.Println("insert data in inbox error: ", err)
		}

		log.Println("Inserted a single document: ", ir.InsertedID)

	}()
	return "ok"
}
