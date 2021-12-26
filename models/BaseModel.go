package models

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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

//Insert adds new document to db and updates the ID field.
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

//FindOne finds one document from db using provided filter.
//If found, it decodes the document into provided document param.
func (bm *BaseModel) FindOne(ctx context.Context, filter primitive.M, document interface{}) error {
	err := bm.collection.FindOne(ctx, filter).Decode(document)
	if err != nil {
		return err
	}

	bm.isLoaded = true
	return nil
}

//Update finds a document using primary _id then replaces it with given values.
func (bm *BaseModel) Update(ctx context.Context, document interface{}) error {
	doc, err := bm.collection.ReplaceOne(ctx, bson.M{"_id": bm.ID}, document)
	if err != nil {
		return err
	}
	if doc.MatchedCount == 0 {
		return errors.New("update failed. zero matches returned.")
	}
	return nil
}

//Loaded returns wether document was successfully loaded from db.
func (bm *BaseModel) Loaded() bool {
	return bm.isLoaded
}
