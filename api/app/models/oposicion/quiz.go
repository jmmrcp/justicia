package oposicion

import (
	"time"
)

type (
	// Quiz schema Mongo DB
	Quiz struct {
		ID         string    `bson:"_id" json:"id"`
		Categoria  string    `bson:"categoria" json:"categoria"`
		Test       int       `bson:"test" json:"test"`
		Ord        int       `bson:"ord" json:"ord"`
		Pregunta   string    `bson:"pregunta" json:"pregunta"`
		Respuestas []string  `bson:"respuestas" json:"respuestas"`
		Articulo   string    `bson:"articulo" json:"articulo"`
		Fecha      time.Time `bson:"fecha" json:"fecha"`
		Box        int       `bson:"box" json:"box"`
	}
)
