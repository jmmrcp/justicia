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
	DBData struct {
		*mgo.Session
		*mgo.Collection
	}
)

// MlabDB Conexion con Mlab
func MlabDB() DBData {
	uri := "mongodb://jmmrcp:J538MTUSbg3v3Vh@ds263876.mlab.com:63876/justicia"
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetSafe(&mgo.Safe{})
	collection := session.DB("justicia").C("preguntas")

	return DBData{
		session,
		collection,
	}
}

// All() conexion a la DB
func (data *DBData) All() {
	test := "34"
	c := data.Collection
	gamesWon, err := c.Find(bson.M{"test": test}).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s has won %d games.\n", test, gamesWon)
}
