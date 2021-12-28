package util

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ICollection interface {
	FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) ISingleResult
	ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error)
	InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error)
}

type ISingleResult interface {
	Decode(v interface{}) error
}

type SingleResult struct {
	*mongo.SingleResult
}

func (r *SingleResult) Decode(v interface{}) error {
	return r.SingleResult.Decode(v)
}

type Collection struct {
	*mongo.Collection
}

func (c *Collection) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) ISingleResult {
	singleResult := c.Collection.FindOne(ctx, filter, opts...)
	return &SingleResult{SingleResult: singleResult}
}

func (c *Collection) ReplaceOne(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return c.Collection.ReplaceOne(ctx, filter, replacement, opts...)
}

func (c *Collection) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return c.Collection.InsertOne(ctx, document, opts...)
}
