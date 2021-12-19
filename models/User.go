package models

import (
	"context"

	"github.com/nicklaw5/helix"
	"go.mongodb.org/mongo-driver/bson"
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
}

func NewUser(collection *mongo.Collection) *User {
	return &User{
		BaseModel: BaseModel{
			collection: collection,
		},
	}
}

func (u *User) FindByUserId(ctx context.Context, userId int) error {
	return u.FindOne(ctx, bson.M{"userId": userId}, u)

}

func (u *User) Create(ctx context.Context) error {
	return u.Insert(ctx, u)
}

func (u *User) GetUserFromTwitch(accessToken string) error {
	client, err := helix.NewClient(&helix.Options{
		accessToken: accessToken,
	})

	if err != nil {
		return err
	}

	return nill
}
