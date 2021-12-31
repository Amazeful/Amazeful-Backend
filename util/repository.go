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
	InsertOne(ctx context.Context, document Model, opts ...*options.InsertOneOptions) error
	FindOne(ctx context.Context, filter bson.M, document Model, opts ...*options.FindOneOptions) error
	ReplaceOne(ctx context.Context, filter bson.M, replacement Model, opts ...*options.ReplaceOptions) error
	FindAll(ctx context.Context, filter bson.M, results []Model, opts ...*options.FindOptions) error
	DeleteOne(ctx context.Context, filter bson.M, opts ...*options.DeleteOptions) error
}

type MongoRepository struct {
	Database   consts.MongoDatabase
	Collection consts.MongoCollection
	client     *mongo.Client
}

func NewMongoRepository(client *mongo.Client, db consts.MongoDatabase, collection consts.MongoCollection) *MongoRepository {
	return &MongoRepository{
		Database:   db,
		Collection: collection,
		client:     client,
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

func (r *MongoRepository) FindAll(ctx context.Context, filter bson.M, results []Model, opts ...*options.FindOptions) error {
	cursor, err := r.client.Database(string(r.Database)).Collection(string(r.Collection)).Find(ctx, filter, opts...)
	if err != nil {
		return err
	}
	err = cursor.All(ctx, results)
	for _, result := range results {
		result.SetLoaded(true)
	}
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

func (r *MongoRepository) DeleteOne(ctx context.Context, filter bson.M, opts ...*options.DeleteOptions) error {
	result, err := r.client.Database(string(r.Database)).Collection(string(r.Collection)).DeleteOne(ctx, filter, opts...)
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("zero matches returned")
	}
	return nil
}
