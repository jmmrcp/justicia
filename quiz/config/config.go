package config

import (
	"context"
	"os"

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
	// Config information
	Config struct {
		proto    string
		user     string
		password string
		host     string
		options  string
		db       string
	}
)

// New returns a new Config struct
func New() *Config {
	return &Config{
		proto:    getEnv("PROTO", "mongodb"),
		user:     getEnv("USERNAME", ""),
		password: getEnv("PASSWORD", ""),
		host:     getEnv("HOST", ""),
		options:  getEnv("OPTIONS", ""),
		db:       getEnv("DB", ""),
	}
}

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// GetMongoDB concecta a la base de datos en Mongo
func GetMongoDB() (*DB, error) {
	conf := New()
	// mongoURI := fmt.Sprintf("%s://%s:%s@%s/%s", conf.proto, conf.user, conf.password, conf.host, conf.options)
	mongoURI := "mongodb://jmmrcpsip:SK0umjgZr0qxTS3b@justice-shard-00-00-sbfoj.mongodb.net:27017,justice-shard-00-01-sbfoj.mongodb.net:27017,justice-shard-00-02-sbfoj.mongodb.net:27017/test?ssl=true&replicaSet=Justice-shard-0&authSource=admin&retryWrites=true&w=majority"
	client, err := mongo.NewClient(
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		return nil, err
	}
	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database(conf.db)
	return &DB{db, client, ctx}, nil
}
