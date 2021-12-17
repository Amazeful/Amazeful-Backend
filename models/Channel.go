package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Channel struct {
	BaseModel `bson:",inline"`

	ChannelId    int                  `bson:"channelId" json:"channelId"`
	Login        string               `bson:"login" json:"login"`
	DisplayName  string               `bson:"displayName" json:"displayName"`
	Language     string               `bson:"language" json:"language"`
	GameId       string               `bson:"gameId" json:"gameId"`
	GameName     string               `bson:"gameName" json:"gameName"`
	Title        string               `bson:"title" json:"title"`
	Joined       bool                 `bson:"joined" json:"joined"`
	Silenced     bool                 `bson:"silenced" json:"silenced"`
	AccessToken  string               `bson:"accessToken,omitempty" json:"accessToken"`
	RefreshToken string               `bson:"refreshToken,omitempty" json:"refreshToken"`
	Prefix       string               `bson:"prefix" json:"prefix"`
	Live         bool                 `bson:"live" json:"live"`
	Shard        int                  `bson:"shard" json:"shard"`
	StartedAt    time.Time            `bson:"startedAt,omitempty" json:"startedAt"`
	EndedAt      time.Time            `bson:"endedAt,omitempty" json:"endedAt"`
	Moderator    bool                 `bson:"moderator" json:"moderator"`
	Owner        primitive.ObjectID   `bson:"owner" json:"-"`
	Managers     []primitive.ObjectID `bson:"managers" json:"-"`

	collection *mongo.Collection
}

func NewChannel(collection *mongo.Collection) *Channel {
	return &Channel{
		collection: collection,
	}
}

func (c *Channel) FindByChannelId(ctx context.Context, channelId int) error {
	return c.collection.FindOne(ctx, bson.M{"channelId": channelId}).Decode(c)
}
