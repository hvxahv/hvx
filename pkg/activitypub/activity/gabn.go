package activity

import (
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"hvxahv/pkg/mongo"
	"log"
)

// GetArticleByName ...
func GetArticleByName(name string) []*map[string]interface{} {
	// 从 MongoDB 取出
	db := mongo.GetMongo()
	f := bson.M{"actor": name}

	co := db.Collection("articles")
	var i []*map[string]interface{}
	findA, err := co.Find(context.TODO(), f, nil)
	if err != nil {
		log.Println(err)
	}
	for findA.Next(context.TODO()) {
		var el map[string]interface{}
		if err := findA.Decode(&el); err != nil {
			log.Println(err)
		}
		i = append(i, &el)
	}
	if err := findA.Err(); err != nil {
		log.Println(err)
	}
	_ = findA.Close(context.TODO())

	return i
}

