package util

import (
	"context"

	"github.com/Amazeful/Amazeful-Backend/config"
	"github.com/Amazeful/dataful"
)

var db dataful.Database

//InitDB initializes the db
func InitDB(ctx context.Context) error {
	database, err := dataful.NewMongoDB(ctx, config.GetConfig().ServerConfig.MongoURI)
	if err != nil {
		return err
	}

	db = database
	return nil
}

//GetDB returns global DB
func GetDB() dataful.Database {
	return db
}
