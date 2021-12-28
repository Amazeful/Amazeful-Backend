package models

import (
	"context"

	"github.com/Amazeful/helix"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	BaseModel `bson:",inline"`

	UserID          string             `bson:"userId" json:"userId"`
	Login           string             `bson:"login" json:"login"`
	DisplayName     string             `bson:"displayName" json:"displayName"`
	AccessToken     string             `bson:"accessToken" json:"-"`
	RefreshToken    string             `bson:"refreshToken" json:"-"`
	Type            string             `bson:"type" json:"type"`
	BroadcasterType string             `bson:"broadcasterType" json:"broadcasterType"`
	Description     string             `bson:"description" json:"description"`
	ProfileImageURL string             `bson:"profileImageURL" json:"profileImageURL"`
	OfflineImageURL string             `bson:"offlineImageURL" json:"offlineImageURL"`
	ViewCount       int                `bson:"viewCount" json:"viewCount"`
	Suspended       bool               `bson:"suspended" json:"suspended"`
	Admin           bool               `bson:"admin" json:"admin"`
	Channel         primitive.ObjectID `bson:"channel" json:"channel"`
}

func NewUser(collection *mongo.Collection) *User {
	return &User{
		BaseModel: BaseModel{
			collection: collection,
		},
	}
}

func (u *User) FindBylId(ctx context.Context, id primitive.ObjectID) error {
	return u.FindOne(ctx, bson.M{"_id": id}, u)
}

func (u *User) FindByUserId(ctx context.Context, userId string) error {
	return u.FindOne(ctx, bson.M{"userId": userId}, u)

}

func (u *User) Create(ctx context.Context) error {
	return u.Insert(ctx, u)
}

func (u *User) Update(ctx context.Context) error {
	return u.ReplaceOne(ctx, u)
}

func (u *User) HydrateFromHelix(user *helix.UserResponse) {
	u.UserID = user.Data.ID
	u.Login = user.Data.Login
	u.DisplayName = user.Data.DisplayName
	u.Type = user.Data.Type
	u.BroadcasterType = user.Data.BroadcasterType
	u.Description = user.Data.Description
	u.ProfileImageURL = user.Data.ProfileImageURL
	u.OfflineImageURL = user.Data.OfflineImageURL
	u.ViewCount = user.Data.ViewCount
}
