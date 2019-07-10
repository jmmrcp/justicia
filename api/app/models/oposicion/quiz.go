package oposicion

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// Quiz schema Mongo DB
	Quiz struct {
		// 	ID         string    `bson:"_id" json:"id"`
		// 	Test       int       `bson:"test" json:"test"`
		// 	Categoria  string    `bson:"categoria" json:"categoria"`
		// 	Ord        int       `bson:"ord" json:"ord"`
		// 	Pregunta   string    `bson:"pregunta" json:"pregunta"`
		// 	Respuestas []string  `bson:"respuestas" json:"respuestas"`
		// 	Articulo   string    `bson:"articulo" json:"articulo"`
		// 	Fecha      time.Time `bson:"fecha" json:"fecha"`
		// 	Box        int       `bson:"box" json:"box"`
		ID         bson.ObjectId `json:"id" bson:"_id"`
		Test       int           `bson:"test" json:"test"`
		Categoria  string        `bson:"categoria" json:"categoria"`
		Tema       []string      `bson:"tema" json:"tema"`
		Titulo     string        `bson:"titulo" json:"titulo,"`
		Ord        int           `bson:"ord" json:"ord"`
		Pregunta   string        `bson:"pregunta" json:"pregunta"`
		Respuestas []string      `bson:"respuestas" json:"respuestas"`
		Articulo   string        `bson:"articulo" json:"articulo"`
		Fecha      time.Time     `bson:"fecha" json:"fecha"`
		Box        int           `bson:"box" json:"box"`
	}
)
