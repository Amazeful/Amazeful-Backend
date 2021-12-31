package models

import (
	"context"
	"time"

	"github.com/Amazeful/Amazeful-Backend/util"
	"github.com/Amazeful/helix"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Channel struct {
	util.BaseModel `bson:",inline"`

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

func NewChannel(r util.Repository) *Channel {
	return &Channel{
		BaseModel: util.BaseModel{
			R: r,
		},
	}
}

func (c *Channel) FindBylId(ctx context.Context, id primitive.ObjectID) error {
	return c.R.FindOne(ctx, bson.M{"_id": id}, c)
}

func (c *Channel) FindByChannelName(ctx context.Context, name string) error {
	return c.R.FindOne(ctx, bson.M{"broadcasterName": name}, c)
}

func (c *Channel) FindByChannelId(ctx context.Context, channelId string) error {
	return c.R.FindOne(ctx, bson.M{"channelId": channelId}, c)
}

func (c *Channel) Create(ctx context.Context) error {
	return c.R.InsertOne(ctx, c)
}

func (c *Channel) Update(ctx context.Context) error {
	return c.R.ReplaceOne(ctx, bson.M{"_id": c.ID}, c)
}

func (c *Channel) HydrateFromHelix(channel *helix.GetChannelInformationResponse) {
	c.ChannelId = channel.Data.BroadcasterID
	c.BroadcasterName = channel.Data.BroadcasterName
	c.Language = channel.Data.BroadcasterLanguage
	c.GameId = channel.Data.GameID
	c.GameName = channel.Data.GameName
	c.Title = channel.Data.Title
}

func (c *Channel) UpdateCustomFields(newChannel *Channel) {
	c.Joined = newChannel.Joined
	c.Silenced = newChannel.Silenced
	c.Prefix = newChannel.Prefix
}
