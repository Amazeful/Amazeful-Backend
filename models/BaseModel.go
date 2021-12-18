package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type BaseModel struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	CreatedAt time.Time          `bson:"createdAt" json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt" json:"updatedAt"`

	isLoaded   bool
	collection *mongo.Collection
}

func (bm *BaseModel) created() {
	bm.CreatedAt = time.Now().UTC()
	bm.UpdatedAt = time.Now().UTC()
}

func (bm *BaseModel) updated() {
	bm.UpdatedAt = time.Now().UTC()
}

func (bm *BaseModel) Insert(ctx context.Context, document interface{}) error {
	bm.created()
	result, err := bm.collection.InsertOne(ctx, document)
	if err != nil {
		return err
	}
	bm.ID = result.InsertedID.(primitive.ObjectID)
	bm.isLoaded = true
	return nil
}

func (bm *BaseModel) FindOne(ctx context.Context, filter primitive.M, document interface{}) error {
	err := bm.collection.FindOne(ctx, filter).Decode(document)
	if err != nil {
		return err
	}

	bm.isLoaded = true
	return nil
}

func (bm *BaseModel) Loaded() bool {
	return bm.isLoaded
}
