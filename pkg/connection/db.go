package connection

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var db *mongo.Client

func initialize(dbConnectionString string) {
	ctx, _ := context.WithTimeout(context.Background(), initTimeout)

	clientOptions := options.Client().ApplyURI(dbConnectionString)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		panic("WDB:init(): Cannot connect to the DB. (/¯◡ ‿ ◡)/¯ ~ ┻━┻")
	}

	db = client
}

var inited = false

func Db(dbName, dbConnectionString string) *mongo.Database {
	if !inited {
		initialize(dbConnectionString)
		inited = true
	}

	return db.Database(dbName)
}
