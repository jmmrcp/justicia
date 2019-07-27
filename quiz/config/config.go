package config

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// COLLECTION Document TABLE
	COLLECTION = "preguntas"
)

type (
	// DB information
	DB struct {
		*mongo.Database
		*mongo.Client
		context.Context
	}
)

// GetMlabDB concecta a la base de datos en MLab
func GetMlabDB() (*mongo.Database, error) {
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://jmmrcpsip:SK0umjgZr0qxTS3b@justice-shard-00-00-sbfoj.mongodb.net:27017,justice-shard-00-01-sbfoj.mongodb.net:27017,justice-shard-00-02-sbfoj.mongodb.net:27017/test?ssl=true&replicaSet=Justice-shard-0&authSource=admin&retryWrites=true&w=majority"),
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

// GetMongoDB concecta a la base de datos en Mongo
func GetMongoDB() (*DB, error) {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://jmmrcpsip:SK0umjgZr0qxTS3b@justice-shard-00-00-sbfoj.mongodb.net:27017,justice-shard-00-01-sbfoj.mongodb.net:27017,justice-shard-00-02-sbfoj.mongodb.net:27017/test?ssl=true&replicaSet=Justice-shard-0&authSource=admin&retryWrites=true&w=majority"),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database("justicia")
	return &DB{db, client, ctx}, nil
}

// GetCleverDB concecta a la base de datos en Mongo
func GetCleverDB() (*mongo.Database, error) {
	// ctx := context.Background()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb://u2mnmnnqwwiu6tr0ugs9:4uYFMYouheqLOLnCeipb@bhmydkfk8yl7h0i-mongodb.services.clever-cloud.com:27017/bhmydkfk8yl7h0i"),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database("bhmydkfk8yl7h0i")
	return db, nil
}

// GetAtlasDB concecta a la base de datos en Mongo
func GetAtlasDB() (*DB, error) {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	// defer cancel()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb+srv://jmmrcpsip:SK0umjgZr0qxTS3b@justice-sbfoj.mongodb.net/test?retryWrites=true&w=majority"),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database("justicia")
	return &DB{db, client, ctx}, nil
}
