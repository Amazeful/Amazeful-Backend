package util

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/consts"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client
var NewRepository GetRepo = defaultRepoGetter

type GetRepo func(dbName consts.MongoDatabase, collection consts.MongoCollection) Repository

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

	return nil
}

//GetDB returns mongodb instance
func GetMongoClient() *mongo.Client {
	return mongoClient
}

func defaultRepoGetter(dbName consts.MongoDatabase, collection consts.MongoCollection) Repository {
	return NewMongoRepository(mongoClient, dbName, collection)
}

//SetMockRepoGetter replaces the default repo getter with a fake one.
//Should ONLY be used in test files.
func SetMockRepoGetter(r Repository) {
	NewRepository = func(_ consts.MongoDatabase, _ consts.MongoCollection) Repository {
		return r
	}
}
