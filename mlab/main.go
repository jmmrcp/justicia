package main

import (
	"fmt"
	"justicia/mlab/models/lite"
	"justicia/mlab/models/mlab"
	"log"
	"time"
)

type (
	// Env alamcena toda la base de datos
	Env struct {
		db lite.Datastore
	}
	allDB struct{}
)

func main() {
	enviroment, err := initializeSQL()
	if err != nil {
		log.Fatal(err)
	}

	enviroment.Update()
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
	mlabdb, err := mlab.NewDB()
	if err != nil {
		return err
	}
	questions, err := env.db.All()

	fmt.Printf("questions: %+v\n", len(questions))

	if err != nil {
		return err
	}

	for _, question := range questions {

		mlab := new(mlab.Mlab)
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

		err = mlabdb.InsertOne(&mlab)
		if err != nil {
			return err
		}
	}
	return nil
}

func initializeSQL() (*Env, error) {
	db, err := lite.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	env := &Env{db}
	return env, nil
}

func listAll() {
	db, err := mlab.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	questions, err := db.GetAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, question := range questions {
		fmt.Println(question)
	}
}
