package activity

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/net/context"
	"hvxahv/pkg/db"
	"log"
)

func GetPublicArticleById(activityId string, c *gin.Context) {
	// 从 MongoDB 取出
	db := db.GetMongo()
	f := bson.M{"id": activityId}

	co := db.Collection("articles")

	var a bson.M
	if err := co.FindOne(context.TODO(), f).Decode(&a); err != nil {
		log.Fatal(err)
	}


	c.JSON(200, gin.H{
		"res": a,
	})
}