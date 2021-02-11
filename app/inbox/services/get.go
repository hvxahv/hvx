package services

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	pb "hvxahv/api/hvxahv/v1"
	db2 "hvxahv/pkg/db"
	"log"
)

// GetInboxData ...
func GetInboxData(in *pb.Name) []*pb.InboxData {
	// 从 MongoDB 取出
	db, err := db2.GetMongo()
	if err != nil {
		log.Println(err)
	}
	f := bson.M{"name": in.Name}

	co := db.Collection("inbox")
	var i []*pb.InboxData
	findR, err := co.Find(context.TODO(), f, nil)
	if err != nil {
		log.Println(err)
	}
	for findR.Next(context.TODO()) {
		var el pb.InboxData
		if err := findR.Decode(&el); err != nil {
			log.Println(err)
		}
		i = append(i, &el)
	}
	if err := findR.Err(); err != nil {
		log.Println(err)
	}
	_ = findR.Close(context.TODO())
	return i
}
