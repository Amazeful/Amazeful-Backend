package consts

type MongoCollection string
type MongoDatabase string

const (
	DBAmazeful        MongoDatabase   = "Amazeful"
	CollectionChannel MongoCollection = "channel"
	CollectionUser    MongoCollection = "user"
	CollectionCommand MongoCollection = "command"
)
