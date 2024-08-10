package database

import (
	"context"
	"log"

	"github.com/cortezzIP/Kacherga-News/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func ConnectToDb() {
	var err error

	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(config.Config.DBURI))
	if err != nil {
		panic(err)
	}

	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}
}

func CloseConnection() {
	err := MongoClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
}
