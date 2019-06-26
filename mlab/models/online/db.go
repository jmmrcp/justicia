package online

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// DB dotos de conexion
	DB struct {
		*mongo.Database
		context.Context
	}
)

// NewDB Conexion con Mlab
func NewDB() (*DB, error) {
	ctx := context.Background()
	// uri := "mongodb://jmmrcp:J538MTUSbg3v3Vh@ds263876.mlab.com:63876/justicia"
	//uri := "mongodb://u2mnmnnqwwiu6tr0ugs9:4uYFMYouheqLOLnCeipb@bhmydkfk8yl7h0i-mongodb.services.clever-cloud.com:27017/bhmydkfk8yl7h0i"
	uri := "mongodb://jmmrcpsip:SK0umjgZr0qxTS3b@justice-shard-00-00-sbfoj.mongodb.net:27017,justice-shard-00-01-sbfoj.mongodb.net:27017,justice-shard-00-02-sbfoj.mongodb.net:27017/test?ssl=true&replicaSet=Justice-shard-0&authSource=admin&retryWrites=true&w=majority"
	client, err := mongo.NewClient(
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database("justicia")
	//db := client.Database("bhmydkfk8yl7h0i")
	return &DB{db, ctx}, nil
}
