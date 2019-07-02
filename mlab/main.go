package main

import (
	"fmt"
	"justicia/mlab/models/lite"
	"justicia/mlab/models/online"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"log"
	"time"
)

type (
	// Env alamcena toda la base de datos
	Env struct {
		db lite.Datastore
	}
)

var allDB []interface{}

func main() {
	local, err := initializeSQL()
	if err != nil {
		log.Fatal(err)
	}
	err = local.Update()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Base de datos MLabs actualizada.")
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
	mlabdb, err := online.NewDB()
	if err != nil {
		return err
	}
	questions, err := env.db.All()
	if err != nil {
		return err
	}

	fmt.Println("Borrando la Base de Datos.")
	// mlabdb.Collection("preguntas").Drop(nil)
	fmt.Printf("Subiendo questions: %+v\n", len(questions))

	for _, question := range questions {

		mlab := new(online.Mlab)

		mlab.ID = primitive.NewObjectID()
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

		allDB = append(allDB, mlab)
	}
	// err = mlabdb.InsertOne(&mlab)
	// if err != nil {
	// 	return err
	// }

	// trainers := []interface{}{misty, brock}

	err = mlabdb.InsertMany(allDB)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Inserted multiple documents: ")

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
	db, err := online.NewDB()
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

func labs() {
	enviroment, err := initializeSQL()
	if err != nil {
		log.Fatal(err)
	}

	enviroment.Update()
}

func list(db []*online.Mlab) {
	for _, question := range db {
		fmt.Println("Categoria:", question.Categoria)
		fmt.Println("Test:", question.Test)
		fmt.Println("ID:", question.ID.Hex())
		fmt.Println("Pregunta:", question.Pregunta)
		for i, respuesta := range question.Respuestas {
			fmt.Println(i, ") "+respuesta)
		}
		fmt.Println("Articulo: ", question.Articulo)
	}
}

func nyFuck() {
	mlabdb, err := online.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	questions, err := mlabdb.GetTest()
	if err != nil {
		log.Fatal(err)
	}
	list(questions)
}
