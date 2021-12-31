package util

import (
	"context"
	"errors"

	"github.com/Amazeful/Amazeful-Backend/consts"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository interface {
	New(db consts.MongoDatabase, collection consts.MongoCollection) Repository
	InsertOne(ctx context.Context, document Model, opts ...*options.InsertOneOptions) error
	FindOne(ctx context.Context, filter bson.M, document Model, opts ...*options.FindOneOptions) error
	ReplaceOne(ctx context.Context, filter bson.M, replacement Model, opts ...*options.ReplaceOptions) error
}

type MongoRepository struct {
	Database   consts.MongoDatabase
	Collection consts.MongoCollection
	client     *mongo.Client
}

func NewWithClient(client *mongo.Client) *MongoRepository {
	return &MongoRepository{
		client: client,
	}
}

func (r *MongoRepository) New(db consts.MongoDatabase, collection consts.MongoCollection) Repository {
	return &MongoRepository{
		Database:   db,
		Collection: collection,
		client:     r.client,
	}
}

func (r *MongoRepository) FindOne(ctx context.Context, filter bson.M, document Model, opts ...*options.FindOneOptions) error {
	err := r.client.Database(string(r.Database)).Collection(string(r.Collection)).FindOne(ctx, filter, opts...).Decode(document)
	if err == mongo.ErrNoDocuments {
		return nil
	} else if err != nil {
		return err
	}
	document.SetLoaded(true)
	return nil
}

func (r *MongoRepository) ReplaceOne(ctx context.Context, filter bson.M, replacement Model, opts ...*options.ReplaceOptions) error {
	updateResult, err := r.client.Database(string(r.Database)).Collection(string(r.Collection)).ReplaceOne(ctx, filter, replacement, opts...)
	if err != nil {
		return err
	}
	if updateResult.MatchedCount == 0 {
		return errors.New("zero matches returned")
	}
	return nil
}

func (r *MongoRepository) InsertOne(ctx context.Context, document Model, opts ...*options.InsertOneOptions) error {
	insertResult, err := r.client.Database(string(r.Database)).Collection(string(r.Collection)).InsertOne(ctx, document, opts...)
	if err != nil {
		return err
	}

	document.SetId(insertResult.InsertedID)
	document.SetLoaded(true)
	return nil
}
