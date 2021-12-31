package util

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/consts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var repository Repository

//InitDB initializes mongo db instance
func InitDB(ctx context.Context) error {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().MongoURI))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	mongoClient = client
	repository = NewWithClient(mongoClient)

	return nil
}

//GetDB returns mongodb instance
func GetMongoClient() *mongo.Client {
	return mongoClient
}

func NewRepository(dbName consts.MongoDatabase, collection consts.MongoCollection) Repository {
	return repository.New(dbName, collection)
}

//SetRepository can be used to set a fake repository for testing
func SetRepository(r Repository) {
	repository = r
}
