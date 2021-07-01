package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func Connect(dbConnectionString string) {
	ctx, _ := context.WithTimeout(context.Background(), initTimeout)

	clientOptions := options.Client().ApplyURI(dbConnectionString)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic("WDB:init(): Cannot connect to the DB. (/¯◡ ‿ ◡)/¯ ~ ┻━┻")
	}

	db = client
}

func Db(dbName string) *mongo.Database {
	if db == nil {
		return nil
	}

	return db.Database(dbName)
}
