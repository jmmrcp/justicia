package main

import (
	"fmt"
	"justicia/tools/mongo/models/lite"
	"justicia/tools/mongo/models/online"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
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
	fmt.Println("Bases de datos actualizadas.")

}

func initializeSQL() (*Env, error) {
	db, err := lite.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	env := &Env{db}
	return env, nil
}

// Update actualiza mlabs
func (env *Env) Update() error {

	questions, err := env.db.All()
	if err != nil {
		return err
	}

	fmt.Printf("Subiendo %+v questions.\n", len(questions))

	sess, err := mgo.Dial("mongodb://justicia:Ha11e_B3rr4@ds249127.mlab.com:49127/justicia")
	if err != nil {
		fmt.Printf("Can't connect to mongo, go error %v\n", err)
		os.Exit(1)
	}
	defer sess.Close()
	sess.SetSafe(&mgo.Safe{})

	collection := sess.DB("justicia").C("preguntas")
	// collection.DropCollection()
	// fmt.Println("Borrando la Base de Datos Mlab.")

	for _, question := range questions {
		var T = []int{}

		data := new(online.Mlab)

		t := strings.Split(question.Tema, ",")
		for _, v := range t {
			te, _ := strconv.Atoi(v)
			T = append(T, te)
		}

		// data.ID = primitive.NewObjectID()
		data.Test = question.Test
		data.Categoria = question.Categoria
		data.Tema = T
		data.Titulo = question.Titulo
		data.Ord = question.Ord
		data.Pregunta = question.Pregunta
		data.Respuestas = []string{
			question.Respuesta1,
			question.Respuesta2,
			question.Respuesta3,
			question.Respuesta4,
		}
		data.Articulo = question.Articulo
		data.Fecha = time.Now()
		data.Box = question.Box

		allDB = append(allDB, data)

		// data.ID = bson.NewObjectId()
		err = collection.Insert(data)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Inserted multiple documents on Mlabs.")
	fmt.Println("Connection to MongoDB Mlabs closed.")

	clever, err := online.NewClever()
	if err != nil {
		return err
	}
	clever.Collection("preguntas").Drop(nil)
	fmt.Println("Borrando la Base de Datos Clever.")

	err = clever.InsertMany(allDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents on Clever.")

	atlas, err := online.NewAtlas()
	if err != nil {
		return err
	}
	atlas.Collection("preguntas").Drop(nil)
	fmt.Println("Borrando la Base de Datos Atlas.")

	err = atlas.InsertMany(allDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents on Atlas.")

	return nil
}
