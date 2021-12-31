package embeddables

type Cooldown struct {
	Global int `bson:"global" json:"global"`
	User   int `bson:"user" json:"user"`
}
