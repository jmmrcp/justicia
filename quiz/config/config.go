// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package config

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// DB information
	DB struct {
		*mongo.Database
		*mongo.Client
		*mongo.Collection
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
		proto:    getEnv("PROTO", ""),
		user:     getEnv("USERNAME", ""),
		password: getEnv("PASSWORD", ""),
		host:     getEnv("HOST", ""),
		options:  getEnv("OPTIONS", ""),
		db:       getEnv("DB", ""),
	}
}

// Simple helper function to read an environment or return a default value
//
// TODO: What's do it?
//
// getEnv read .env file.
func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

// GetMongoDB conect DB.
func GetMongoDB() (*DB, error) {
	conf := New()
	mongoURI := newURI(conf)
	client, err := mongo.NewClient(
		options.Client().ApplyURI(mongoURI),
	)

	// if not a pre-concection with DB return a Error.
	if err != nil {
		return nil, err
	}

	// if not conect return a Error
	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		return nil, err
	}

	// conect exist, return a conection
	db := client.Database(conf.db)
	return &DB{
		db,
		client,
		db.Collection(os.Getenv("COLLECTION")),
		ctx,
	}, nil
}

func newURI(conf *Config) string {
	mongoURI := conf.proto + "://" + conf.user + ":" + conf.password + "@" + conf.host + "/" + conf.db + "?" + conf.options
	return mongoURI
}
