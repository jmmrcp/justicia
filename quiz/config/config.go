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
	ctx := context.Background()
	conf := New()
	uri := conf.proto + "://" +
		conf.user + ":" +
		conf.password + "@" +
		conf.host + "/" +
		conf.options
	client, err := mongo.NewClient(
		options.Client().ApplyURI(uri),
	)
	if err != nil {
		return nil, err
	}
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}
	db := client.Database(conf.db)
	return &DB{db, client, ctx}, nil
}
