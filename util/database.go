package util

import (
	"context"

	"github.com/Amazeful/dataful"
)

var db dataful.Database

//InitDB initializes the db
func InitDB(ctx context.Context, uri string) error {
	var err error
	db, err = dataful.NewMongoDB(ctx, uri)

	return err
}

//DB returns global DB
func DB() dataful.Database {
	return db
}
