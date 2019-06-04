package models

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Mlab schema Mongo DB
	Mlab struct {
		Categoria  string    `bson:"categoria"`
		Test       int       `bson:"test"`
		Ord        int       `bson:"ord"`
		Pregunta   string    `bson:"pregunta"`
		Respuestas []string  `bson:"respuestas"`
		Articulo   string    `bson:"articulo"`
		Fecha      time.Time `bson:"fecha"`
		Box        int       `bson:"box"`
	}
	// DBData dotos de conexion
	DBData struct {
		*mgo.Database
	}
)

// MlabDB Conexion con Mlab
func MlabDB() (*DBData, error) {
	uri := "mongodb://jmmrcp:J538MTUSbg3v3Vh@ds263876.mlab.com:63876/justicia"
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}
	session.SetSafe(&mgo.Safe{})
	db := session.DB("justicia")

	return &DBData{db}, nil
}

// All conexion a la DB
func (data *DBData) All() {
	test := "34"
	c := data.C("preguntas")
	gamesWon := c.Find(bson.M{}).All()
	fmt.Printf("%s has won %v games.\n", test, gamesWon)
}
