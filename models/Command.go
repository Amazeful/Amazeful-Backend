package models

import (
	"github.com/Amazeful/Amazeful-Backend/util"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Command struct {
	BaseModel `bson:",inline"`

	Name       string            `bson:"name" json:"name"`
	Enabled    bool              `bson:"enabled" json:"enabled"`
	Cooldowns  Cooldown          `bson:"cooldowns" json:"cooldowns"`
	Role       UserRole          `bson:"role" json:"role"`
	Stream     StreamStatus      `bson:"stream" json:"stream"`
	Response   string            `bson:"response" json:"response"`
	Aliases    []string          `bson:"aliases" json:"aliases"`
	HasVar     bool              `bson:"hasVar" json:"-"`
	Attributes CommandAttributes `bson:"attributes,omitempty" json:"attributes"`

	Channel primitive.ObjectID `bson:"channel" json:"-"`
}

func NewCommand(collection util.ICollection) *Command {
	return &Command{
		BaseModel: BaseModel{
			collection: collection,
		},
		Enabled:   true,
		Cooldowns: Cooldown{Global: 5, User: 15},
		Role:      UserRoleGlobal,
		Stream:    StreamLive | StreamOffline,
	}
}
