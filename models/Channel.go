package models

import (
	"context"
	"time"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/helix"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

type Channel struct {
	BaseModel `bson:",inline"`

	ChannelId       string    `bson:"channelId" json:"channelId"`
	BroadcasterName string    `bson:"broadcasterName" json:"broadcasterName"`
	Language        string    `bson:"language" json:"language"`
	GameId          string    `bson:"gameId" json:"gameId"`
	GameName        string    `bson:"gameName" json:"gameName"`
	Title           string    `bson:"title" json:"title"`
	Joined          bool      `bson:"joined" json:"joined"`
	Silenced        bool      `bson:"silenced" json:"silenced"`
	AccessToken     string    `bson:"accessToken,omitempty" json:"-"`
	RefreshToken    string    `bson:"refreshToken,omitempty" json:"-"`
	Prefix          string    `bson:"prefix" json:"prefix"`
	Live            bool      `bson:"live" json:"live"`
	Shard           int       `bson:"shard" json:"shard"`
	Authenticated   bool      `bson:"authenticated" json:"authenticated"`
	StartedAt       time.Time `bson:"startedAt,omitempty" json:"startedAt"`
	EndedAt         time.Time `bson:"endedAt,omitempty" json:"endedAt"`
	Moderator       bool      `bson:"moderator" json:"moderator"`
}

func NewChannel(collection util.ICollection) *Channel {
	return &Channel{
		BaseModel: BaseModel{
			collection: collection,
		},
	}
}

func (c *Channel) FindBylId(ctx context.Context, id primitive.ObjectID) error {
	return c.FindOne(ctx, bson.M{"_id": id}, c)
}

func (c *Channel) FindByChannelId(ctx context.Context, channelId string) error {
	return c.FindOne(ctx, bson.M{"channelId": channelId}, c)
}

func (c *Channel) Create(ctx context.Context) error {
	return c.Insert(ctx, c)
}

func (c *Channel) Update(ctx context.Context) error {
	return c.ReplaceOne(ctx, c)
}

//GetUserFromTwitch gets user info from twitch helix API.
//If a the user exists in db, it updates the user; otherwise, it creates a new user.
func (c *Channel) GetChannelByAccessToken(ctx context.Context, accessToken string) error {
	client, err := helix.NewClient(&helix.Options{
		ClientID:        config.GetTwitchConfig().ClientID,
		UserAccessToken: accessToken,
	})
	if err != nil {
		return err
	}
	user, err := client.GetChannelInformation(nil)
	if err != nil {
		return err
	}
	zap.L().Info("test", zap.Any("user", user))
	return nil
}

func (c *Channel) HydrateFromHelix(channel *helix.GetChannelInformationResponse) {
	c.ChannelId = channel.Data.BroadcasterID
	c.BroadcasterName = channel.Data.BroadcasterName
	c.Language = channel.Data.BroadcasterLanguage
	c.GameId = channel.Data.GameID
	c.GameName = channel.Data.GameName
	c.Title = channel.Data.Title
}
