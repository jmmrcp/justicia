package models

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type (
	// Mlab schema Mongo DB
	Mlab struct {
		ID         bson.ObjectId `bson:"_id" json:"id"`
		Categoria  string        `bson:"categoria" json:"categoria"`
		Test       int           `bson:"test" json:"test"`
		Ord        int           `bson:"ord" json:"ord"`
		Pregunta   string        `bson:"pregunta" json:"pregunta"`
		Respuestas []string      `bson:"respuestas" json:"respuestas"`
		Articulo   string        `bson:"articulo" json:"articulo"`
		Fecha      time.Time     `bson:"fecha" json:"fecha"`
		Box        int           `bson:"box" json:"box"`
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

// Questions enlaza a la tabla preguntas
func (db *DBData) Questions() *mgo.Collection {
	return db.C("preguntas")
}

// GetAll conexion a la DB
func (db *DBData) GetAll() ([]Mlab, error) {
	questions := []Mlab{}
	c := db.Questions()
	if err := c.Find(nil).All(&questions); err != nil {
		return nil, err
	}
	return questions, nil
}
