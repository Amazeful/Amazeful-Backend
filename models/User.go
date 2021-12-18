package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	BaseModel `bson:",inline"`

	UserID          int    `bson:"userID" json:"userID"`
	Login           string `bson:"login" json:"login"`
	DisplayName     string `bson:"displayName" json:"displayName"`
	Type            string `bson:"type" json:"type"`
	BroadcasterType string `bson:"broadcasterType" json:"broadcasterType"`
	GameName        string `bson:"gameName" json:"gameName"`
	Description     string `bson:"description" json:"description"`
	ProfileImageURL bool   `bson:"profileImageURL" json:"profileImageURL"`
	OfflineImageURL string `bson:"offlineImageURL" json:"offlineImageURL"`
	ViewCount       int    `bson:"viewCount" json:"viewCount"`
	Suspended       bool   `bson:"suspended" json:"suspended"`
	Admin           bool   `bson:"admin" json:"admin"`

	collection *mongo.Collection
}

func NewUser(collection *mongo.Collection) *User {
	return &User{
		collection: collection,
	}
}

func (u *User) FindByUserId(ctx context.Context, userId int) error {
	return u.collection.FindOne(ctx, bson.M{"userId": userId}).Decode(u)
}

func (u *User) Insert(ctx context.Context) error {
	u.BaseModel.Create()
	result, err := u.collection.InsertOne(ctx, u)
	if err != nil {
		return err
	}

	u.ID = result.InsertedID.(primitive.ObjectID)
	return nil
}
