package db

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var (
	MongoDB *mongo.Database
)

// InitMongoDB Initialize MongoDB.
func InitMongoDB() error {
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

	log.Println("Connected to MongoDB!")
	collection := cli.Database(name)
	MongoDB = collection
	return err
}

func GetMongo() *mongo.Database {
	return MongoDB
}
