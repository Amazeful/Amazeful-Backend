package util

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	dbName = "Amazeful"
)

var client *mongo.Client

//InitDB initializes mongo db instance
func InitDB(ctx context.Context) error {
	var err error

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().MongoURI))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
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
