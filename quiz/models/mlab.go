// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	// Mlab Schema MLabs database
	Mlab struct {
		ID         primitive.ObjectID `bson:"_id" json:"id"`
		Test       int                `bson:"test" json:"test"`
		Categoria  string             `bson:"categoria" json:"categoria"`
		Temas      []int              `bson:"temas" json:"temas"`
		Titulo     string             `bson:"titulo" json:"titulo,"`
		Ord        int                `bson:"ord" json:"ord"`
		Pregunta   string             `bson:"pregunta" json:"pregunta"`
		Respuestas []string           `bson:"respuestas" json:"respuestas"`
		Articulos  string             `bson:"articulos" json:"articulos"`
		Fecha      time.Time          `bson:"fecha" json:"fecha"`
		Box        int                `bson:"box" json:"box"`
	}
)

// Parse convert Objet to Array
func (mlab *Mlab) Parse() []string {
	data := []string{mlab.Pregunta}
	data = append(data, mlab.Respuestas...)
	data = append(data, mlab.Articulos, mlab.ID.Hex())
	return data
}
