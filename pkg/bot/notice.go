package bot

import (
	"fmt"
	"golang.org/x/net/context"
	db2 "hvxahv/pkg/db"
	"log"
	"time"
)

// NewAccountNotice 新增用户通知
func NewAccountNotice(str string)  {
	if err := SendNotice(str); err != nil {
		log.Println(err)
	}
	// 将新增用户消息持久化到 mongo
	db, err := db2.InitMongoDB()
	if err != nil {
		log.Println("Use InitDB error: ", err)
	}
	coll := db.Collection("new_account")

	a := new(AccountNotice)
	a.Name = str
	a.Data = time.Now()

	insertResult, err := coll.InsertOne(context.TODO(), a)
	if err != nil {
		log.Println("insert data error: ", err)
	}

	log.Println("Inserted a single document: ", insertResult.InsertedID)
}

// ServicesRunningNotice 服务启动通知
func ServicesRunningNotice(srvname string, port string) {
	str := fmt.Sprintf("%s services is running..., port: %s", srvname, port)
	if err := SendNotice(str); err != nil {
		log.Println(err)
	}

	//// 将服务启动的消息持久化到 mongo
	//db, err := mongo.InitMongoDB()
	//if err != nil {
	//	log.Println("Use InitDB error: ", err)
	//}
	//coll := db.Collection("service_startup_log")
	//
	//a := new(ServicesRunNotice)
	//a.Name = str
	//a.Port = port
	//a.Data = time.Now()
	//
	//insertResult, err := coll.InsertOne(context.TODO(), a)
	//if err != nil {
	//	log.Println("insert data error: ", err)
	//}
	//
	//log.Println("Inserted a single document: ", insertResult.InsertedID)

}