package activitypub

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"hvxahv/pkg/mongo"
	"log"
)

type Follow struct {
	ID    string `json:"_id"`
	Actor string `json:"actor"`
	Date  string `json:"date"`
	Name  string `json:"name"`
}

// GetFollow 获取关注的方法，返回一个数组
func GetFollow(name, collection string) []string {
	db := mongo.GetMongo()
	f := bson.M{"name": name}
	log.Println(name)
	co := db.Collection(collection)
	var i []string
	findA, err := co.Find(context.TODO(), f, nil)
	if err != nil {
		log.Println(err)
	}
	for findA.Next(context.TODO()) {
		var el Follow
		if err := findA.Decode(&el); err != nil {
			log.Println(err)
		}
		i = append(i, el.Actor)
	}
	if err := findA.Err(); err != nil {
		log.Println(err)
	}
	_ = findA.Close(context.TODO())

	return i
}
