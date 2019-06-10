package config

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// COLLECTION Document TABLE
	COLLECTION = "preguntas"
)

// GetMongoDB concecta a la base de datos en MLab
func GetMongoDB() (*mongo.Database, error) {
	ctx := context.Background()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://jmmrcp:J538MTUSbg3v3Vh@ds263876.mlab.com:63876/justicia"),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database("justicia")
	return db, nil
}
