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

type IDB interface {
	Collection(collection string) ICollection
	Disconnect(ctx context.Context) error
}
type DB struct {
	client *mongo.Client
}

//GetCollection returns mongodb collection
func (db *DB) Collection(collection string) ICollection {
	c := db.client.Database(dbName).Collection(collection)
	return &Collection{Collection: c}
}

//Disconnect disconnects db client
func (db *DB) Disconnect(ctx context.Context) error {
	return db.client.Disconnect(ctx)
}

var db IDB

//InitDB initializes mongo db instance
func InitDB(ctx context.Context) error {
	var err error
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.GetConfig().MongoURI))
	if err != nil {
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		return err
	}

	db = &DB{
		client: client,
	}

	return nil
}

//GetDB returns mongodb instance
func GetDB() IDB {
	return db
}

//GetCollection returns mongodb collection
func GetCollection(collection string) ICollection {
	return db.Collection(collection)
}

func SetDB(newDB IDB) {
	db = newDB
}
