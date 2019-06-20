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

// GetMlabDB concecta a la base de datos en MLab
func GetMlabDB() (*mongo.Database, error) {
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

// GetMongoDB concecta a la base de datos en Mongo
func GetMongoDB() (*mongo.Database, error) {
	ctx := context.Background()
	client, err := mongo.NewClient(
		options.Client().ApplyURI("mongodb+srv://jmmrcpsip:SK0umjgZr0qxTS3b@justice-sbfoj.mongodb.net/admin?replicaSet=Justice-shard-0&connectTimeoutMS=10000&authSource=admin&authMechanism=SCRAM-SHA-1&3t.uriVersion=3&3t.connection.name=Justice-shard-0&3t.databases=admin,test"),
		//options.Client().ApplyURI("mongodb+srv://jmmrcpsip:SK0umjgZr0qxTS3b@justice-sbfoj.mongodb.net/test?retryWrites=true&w=majority"),
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

// GetCleverDB concecta a la base de datos en Mongo
func GetCleverDB() (*mongo.Database, error) {
	ctx := context.Background()
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
