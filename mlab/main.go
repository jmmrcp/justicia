package main

import (
	"fmt"
	"justicia/mlab/models"
	"log"
	"time"
)

type (
	// Env alamcena toda la base de datos
	Env struct {
		db models.Datastore
	}
	allDB struct{}
)

func main() {

	db, err := models.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	env := &Env{db}

	// env.Index()

}

//Index lista todos los Datos
func (env *Env) Index() {
	questions, err := env.db.All()
	if err != nil {
		log.Fatal(err)
	}
	for _, question := range questions {
		fmt.Printf("ID : %d\nTest: %d\nTema: %s\n",
			question.ID,
			question.Test,
			question.Tema)
	}
}

// Update actualiza mlabs
func (env *Env) Update() error {
	session := models.MlabDB
	collection := session.DB("justicia").C("preguntas")
	questions, err := env.db.All()
	if err != nil {
		log.Fatal(err)
	}

	for _, question := range questions {

		mlab := new(models.Mlab)
		mlab.Categoria = question.Tema
		mlab.Test = question.Test
		mlab.Ord = question.Ord
		mlab.Pregunta = question.Pregunta
		mlab.Respuestas = []string{
			question.Respuesta1,
			question.Respuesta2,
			question.Respuesta3,
			question.Respuesta4,
		}
		mlab.Articulo = question.Articulo
		mlab.Fecha = time.Now()
		mlab.Box = question.Box

		err = collection.Insert(&mlab)
		if err != nil {
			return err
		}
	}
	return nil
}
