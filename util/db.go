package util

import (
	"context"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "Amazeful"
)

var client *mongo.Client

//InitDB initializes mongo db instance
func InitDB() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var err error

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().MongoURI))
	if err != nil {
		return err
	}

	return nil
}

//GetDB returns mongodb instance
func GetDB() *mongo.Client {
	return client
}

//GetCollection returns mongodb collection
func GetCollection(collection string) *mongo.Collection {
	return client.Database(dbName).Collection(collection)
}
