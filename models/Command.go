package models

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/models/embeddables"
	"github.com/Amazeful/Amazeful-Backend/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Command struct {
	util.BaseModel `bson:",inline"`

	Name       string                        `bson:"name" json:"name"`
	Enabled    bool                          `bson:"enabled" json:"enabled"`
	Cooldowns  embeddables.Cooldown          `bson:"cooldowns" json:"cooldowns"`
	Role       embeddables.UserRole          `bson:"role" json:"role"`
	Stream     embeddables.StreamStatus      `bson:"stream" json:"stream"`
	Response   string                        `bson:"response" json:"response"`
	Aliases    []string                      `bson:"aliases,omitempty" json:"aliases"`
	HasVar     bool                          `bson:"hasVar" json:"-"`
	Attributes embeddables.CommandAttributes `bson:"attributes,omitempty" json:"attributes"`
	Timer      embeddables.Timer             `bson:"timer,omitempty" json:"timer"`
	Channel    primitive.ObjectID            `bson:"channel" json:"channel"`
}

func NewCommand(r util.Repository) *Command {
	return &Command{
		BaseModel: util.BaseModel{
			R: r,
		},
		Enabled:   true,
		Cooldowns: embeddables.Cooldown{Global: 5, User: 15},
		Role:      embeddables.UserRoleGlobal,
		Stream:    embeddables.StreamLive | embeddables.StreamOffline,
	}
}

func (c *Command) Create(ctx context.Context) error {
	return c.R.InsertOne(ctx, c)
}

func (c *Command) Delete(ctx context.Context) error {
	return c.R.DeleteOne(ctx, bson.M{"_id": c.ID})
}

func (c *Command) UpdateCustomFields(newCommand *Command) {
	c.Enabled = newCommand.Enabled
	c.Cooldowns = newCommand.Cooldowns
	c.Role = newCommand.Role
	c.Stream = newCommand.Stream
	c.Response = newCommand.Response
	c.Aliases = newCommand.Aliases
	c.Attributes = newCommand.Attributes
	c.Timer = newCommand.Timer
	c.Attributes = newCommand.Attributes
}
