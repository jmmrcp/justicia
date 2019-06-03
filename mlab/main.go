package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
)

type (
	pregunta struct {
		ID        int
		Test      int
		Tema      string
		Pregunta  string
		Respuesta []string
		Articulo  string
		Ord       int
		Fecha     time.Time
		Box       int
	}
)

func main() {
	uri := "mongodb://jmmrcp:J538MTUSbg3v3Vh@ds263876.mlab.com:63876/justicia"

	session, err := mgo.Dial(uri)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	session.SetSafe(&mgo.Safe{})

	collection := session.DB("justicia").C("preguntas")

	err = collection.Insert(&pregunta{
		ID: 1,
	})
	if err != nil {
		log.Fatal(err)
	}
}
