package mongo

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	MongoDB *mongo.Database
)

// InitMongoDB 初始化 MongoDB
func InitMongoDB() error {
	viper.SetConfigFile("./configs/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

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