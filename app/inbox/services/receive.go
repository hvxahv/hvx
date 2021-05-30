package services

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1alpha1"
	"hvxahv/internal/inbox"
	"hvxahv/pkg/mongo"
	redis2 "hvxahv/pkg/redis"
	"log"
	"net/http"
	"time"
)

// ReceiveInbox 接收 Inbox 数据
func ReceiveInbox(in *pb.InboxData) string {
	i := &inbox.Inbox{
		Actor:     in.Actor,
		RequestId: in.RequestId,
		EventType: in.EventType,
		Name:      in.Name,
		Date:      time.Now().UTC().Format(http.TimeFormat),
	}

	k := fmt.Sprintf("%s-inbox", in.Name)

	// 将数据存储到缓存，返回 ok ，然后将缓存中的数据用 goroutine 持久化到 MongoDB
	rdb := redis2.GetRDB()
	v, _ := json.Marshal(&i)
	if _, err := rdb.Do("SETNX", k, v); err != nil {
		log.Printf("Inbox 持久化到缓存失败: %s", err)
	}

	// 如果是 Accept 的请求类型就把数据保存到 following
	log.Println(in.EventType)

	// 将收到的数据写道 inbox 的收件箱中
	go func() {
		db := mongo.GetMongo()

		co := db.Collection("inbox")
		inbox := inbox.NewInboxStructs(i)
		ir, err := co.InsertOne(context.TODO(), &inbox)
		if err != nil {
			log.Println("insert data in inbox error: ", err)
		}

		log.Println("Inserted to inbox: ", ir.InsertedID)

	}()

	if in.EventType == "Accept" {
		go followingSave(in.Name, in.Actor)

		go func() {
			rdb := redis2.GetRDB()
			v, err := redis.Int64(rdb.Do("INCR", fmt.Sprintf("%s-following", in.Name)))
			if err != nil {
				log.Println("INCR failed:", err)
			}
			log.Println("value:", v)
		}()
	}

	return "ok"
}

func followingSave(name, actor string) {
	db := mongo.GetMongo()
	co := db.Collection("following")
	a := new(inbox.Follow)
	a.Name = name
	a.Actor = actor
	a.Date = time.Now().UTC().Format(http.TimeFormat)

	insertResult, err := co.InsertOne(context.TODO(), a)
	if err != nil {
		log.Println("insert follower error: ", err)
	}

	log.Println("Inserted following: ", insertResult.InsertedID)

}
