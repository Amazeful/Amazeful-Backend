package models

type Cooldown struct {
	Global int `bson:"global" json:"global"`
	User   int `bson:"user" json:"user"`
}
