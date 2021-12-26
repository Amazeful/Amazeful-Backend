package models

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/helix"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	BaseModel `bson:",inline"`

	UserID          string `bson:"userID" json:"userID"`
	Login           string `bson:"login" json:"login"`
	DisplayName     string `bson:"displayName" json:"displayName"`
	AccessToken     string `bson:"accessToken" json:"accessToken"`
	RefreshToken    string `bson:"refreshToken" json:"refreshToken"`
	Type            string `bson:"type" json:"type"`
	BroadcasterType string `bson:"broadcasterType" json:"broadcasterType"`
	Description     string `bson:"description" json:"description"`
	ProfileImageURL string `bson:"profileImageURL" json:"profileImageURL"`
	OfflineImageURL string `bson:"offlineImageURL" json:"offlineImageURL"`
	ViewCount       int    `bson:"viewCount" json:"viewCount"`
	Suspended       bool   `bson:"suspended" json:"suspended"`
	Admin           bool   `bson:"admin" json:"admin"`
	Channel         string `bson:"channel" json:"channel"`
}

func NewUser(collection *mongo.Collection) *User {
	return &User{
		BaseModel: BaseModel{
			collection: collection,
		},
	}
}

func (u *User) FindByUserId(ctx context.Context, userId string) error {
	return u.FindOne(ctx, bson.M{"userId": userId}, u)

}

func (u *User) Create(ctx context.Context) error {
	return u.Insert(ctx, u)
}

//GetUserFromTwitch gets user info from twitch helix API.
//If a the user exists in db, it updates the user; otherwise, it creates a new user.
func (u *User) GetUserFromTwitch(ctx context.Context) error {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        config.GetTwitchConfig().ClientID,
		UserAccessToken: u.AccessToken,
	})
	if err != nil {
		return err
	}
	user, err := client.GetMe()
	if err != nil {
		return err
	}

	u.FindByUserId(ctx, user.Data.ID)
	u.hydrateFromHelix(user)

	if u.Loaded() {
		err = u.Update(ctx, u)
	} else {
		err = u.Create(ctx)
	}
	if err != nil {
		return err
	}
	return nil
}

func (u *User) hydrateFromHelix(user *helix.UserResponse) {
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
