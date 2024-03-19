package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

func NewMongodbConnection(context context.Context) (*mongo.Database, error) {
	mongoUrl := os.Getenv("MONGODB_URL")
	mongoDatabase := os.Getenv("MONGODB_DATABASE")

	client, err := mongo.Connect(context, options.Client().ApplyURI(mongoUrl))

	if err != nil {
		return nil, err
	}

	if err := client.Ping(context, nil); err != nil {
		return nil, err
	}

	return client.Database(mongoDatabase), nil
}
