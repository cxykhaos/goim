package db

import (
	"context"
	"fmt"
	"log"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var c *mongo.Client
var M *mongo.Database

func InitMongo() {
	host := viper.GetString("db.mongo.host")
	port := viper.GetString("db.mongo.port")
	user := viper.GetString("db.mongo.user")
	password := viper.GetString("db.mongo.pwd")
	db := viper.GetString("db.mongo.db")
	url := fmt.Sprintf("mongodb://%s:%s@%s:%s", user, password, host, port)
	clientOptions := options.Client().ApplyURI(url)
	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
	M = client.Database(db)
}

func CloseMongo() {
	c.Disconnect(context.TODO())
}
