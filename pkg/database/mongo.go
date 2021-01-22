package database

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InitDB 初始化 MongoDB
func InitMongoDB() (*mongo.Database, error) {
	url := viper.GetString("mongo.address")
	username := viper.GetString("mongo.username")
	password := viper.GetString("mongo.password")
	name := viper.GetString("mongo.name")

	address := fmt.Sprintf("mongodb://%s:%s@%s", username, password, url)
	option := options.Client().ApplyURI(address)

	cli, err := mongo.Connect(context.TODO(), option)
	if err != nil {
		log.Fatal("MongoDB connect error: ", err)
	}

	err = cli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")
	collection := cli.Database(name)
	return collection, err
}
