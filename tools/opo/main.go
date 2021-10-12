// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"justicia/tools/mongo/models/online"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type (
	// Preguntas array de pregunta
	Preguntas struct {
		Preguntas []Pregunta
	}
	// Pregunta datos para incluir en SQLite
	Pregunta struct {
		ID         int `json:"id"`
		Test       int
		Categoria  string
		Temas      string
		Titulo     string
		Norma      int    `json:"norma"`
		Pregunta   string `json:"enunciado"`
		Respuesta1 string `json:"a"`
		Respuesta2 string `json:"b"`
		Respuesta3 string `json:"c"`
		Respuesta4 string `json:"d"`
		Correcta   string `json:"correcta"`
		Articulos  string
		Ord        int
		Fecha      time.Time
		Cont       int
		Box        int
	}

	// Normas leyes incluidas en el JSON
	Normas struct {
		ID        int    `json:"id,omitempty"`
		Titulo    string `json:"nombre,omitempty"`
		Categoria string `json:"categoria,omitempty"`
		Tema      string `json:"tema,omitempty"`
		Ley       string `json:"ley,omitempty"`
	}
)

var (
	preguntas []Pregunta
	normas    []Normas
	allDB     []interface{}
)

func main() {

	normasJSON := `[
	{
		"id": 1,
		"nombre": "Constitución Española",
		"categoria": "CE",
		"tema": "1",
		"ley": "CE"
	},
	{
		"id": 2,
		"nombre": "Estatuto Básico del Empleado Público",
		"categoria": "EBEP",
		"tema": "7",
		"ley": "EBEP"
	},
	{
		"id": 7,
		"nombre": "Ley de Bases del Régimen Local",
		"categoria": "GOBIERNO",
		"tema": "4",
		"ley": "Ley 7/85"
	},
	{
		"id": 11,
		"nombre": "Ley de Igualdad",
		"categoria": "IGUALDAD",
		"tema": "2",
		"ley": "LO 3/07"
	},
	{
		"id": 15,
		"nombre": "Ley del Gobierno",
		"categoria": "GOBIERNO",
		"tema": "3",
		"ley": "LG"
	},
	{
		"id": 18,
		"nombre": "Ley Orgánica Violencia contra la mujer",
		"categoria": "IGUALDAD",
		"tema": "2",
		"ley": "LO 1/04"
	},
	{
		"id": 24,
		"nombre": "Ley Orgánica del Poder Judicial",
		"categoria": "LOPJ",
		"tema": "6,7,8",
		"ley": "LOPJ"
	},
	{
		"id": 27,
		"nombre": "Ley de Enjuiciamiento Civil",
		"categoria": "CIVIL",
		"tema": "16,17,18",
		"ley": "LEC"
	},
	{
		"id": 29,
		"nombre": "Ley de Régimen Jurídico del Sector Público",
		"categoria": "GOBIERNO",
		"tema": "3",
		"ley": "LSP"
	},
	{
		"id": 32,
		"nombre": "Ley Jurisdicción Contencioso-Administrativa",
		"categoria": "CA",
		"tema": "22",
		"ley": "LJCA"
	},
	{
		"id": 41,
		"nombre": "Ley de Enjuiciamiento Criminal",
		"categoria": "PENAL",
		"tema": "20,21",
		"ley": "LECRIM"
	},
	{
		"id": 48,
		"nombre": "Tratado de la Unión Europea",
		"categoria": "EUROPA",
		"tema": "5",
		"ley": "TUE"
	}
	]`

	json.Unmarshal([]byte(normasJSON), &normas)

	jsonFile, err := os.Open("preguntas.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Opened json file.")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &preguntas)

	contador := 0
	test := 400

	for i := 0; i < len(preguntas)-1; i++ {
		actual := preguntas[i].Norma
		siguiente := preguntas[i+1].Norma

		// Correct on 0 position
		letra := strings.ToUpper(preguntas[i].Correcta)
		switch letra {
		case "B":
			preguntas[i].Respuesta2, preguntas[i].Respuesta1 = preguntas[i].Respuesta1, preguntas[i].Respuesta2
		case "C":
			preguntas[i].Respuesta3, preguntas[i].Respuesta1 = preguntas[i].Respuesta1, preguntas[i].Respuesta3
		case "D":
			preguntas[i].Respuesta4, preguntas[i].Respuesta1 = preguntas[i].Respuesta1, preguntas[i].Respuesta4
		}
		categoria, tema, titulo, ley := leyes(actual)
		t := strings.Split(preguntas[i].Pregunta, "(Art. ")
		preguntas[i].Pregunta = t[0]
		preguntas[i].Test = test
		preguntas[i].Categoria = categoria
		preguntas[i].Temas = tema
		preguntas[i].Titulo = titulo
		if len(t) > 1 {
			a := t[1]
			b := len(a) - 1
			c := a[:b]
			preguntas[i].Articulos = "Art. " + c + " " + ley
		}
		contador++
		preguntas[i].Ord = contador
		preguntas[i].Fecha = time.Now()
		if actual != siguiente {
			contador = 0
			test++
		}
	}
	/*
		fmt.Printf("%+v\n", preguntas[6220])
		fmt.Printf("%+v\n", preguntas[6221])
		fmt.Printf("%+v\n", preguntas[6222])

		fmt.Printf("Subiendo %+v questions.\n", len(preguntas))
		sess, err := mgo.Dial("mongodb://justicia:Ha11e_B3rr4@ds249127.mlab.com:49127/justicia")
		if err != nil {
			fmt.Printf("Can't connect to mongo, go error %v\n", err)
			os.Exit(1)
		}
		defer sess.Close()
		sess.SetSafe(&mgo.Safe{})

		collection := sess.DB("justicia").C("preguntas")
		collection.DropCollection()
		fmt.Println("Borrando la Base de Datos Mlab.")
	*/

	for _, question := range preguntas {
		var T = []int{}

		data := new(online.Mlab)

		t := strings.Split(question.Temas, ",")
		for _, v := range t {
			te, _ := strconv.Atoi(v)
			T = append(T, te)
		}
		// data.ID = primitive.NewObjectID()
		data.Test = question.Test
		data.Categoria = question.Categoria
		data.Temas = T
		data.Titulo = question.Titulo
		data.Ord = question.Ord
		data.Pregunta = question.Pregunta
		data.Respuestas = []string{
			question.Respuesta1,
			question.Respuesta2,
			question.Respuesta3,
			question.Respuesta4,
		}
		data.Articulos = question.Articulos
		data.Fecha = time.Now()
		data.Box = question.Box

		if data.Categoria != "" {
			allDB = append(allDB, data)
			/*
				data.ID = bson.NewObjectId()
				err = collection.Insert(data)
				if err != nil {
					log.Fatal(err)
				}
			*/
		}
	}
	fmt.Printf("Subidas %+v questions.\n", len(allDB))
	/*
		// fmt.Println("Inserted multiple documents on Mlabs.")
		// fmt.Println("Connection to MongoDB Mlabs closed.")
		clever, err := online.NewClever()
		if err != nil {
					log.Fatal(err)
				}
				//clever.Collection("preguntas").Drop(nil)
				//fmt.Println("Borrando la Base de Datos Clever.")

				err = clever.InsertMany(allDB)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Println("Inserted multiple documents on Clever.")
	*/

	atlas, err := online.NewAtlas()
	if err != nil {
		log.Fatal(err)
	}
	atlas.Collection("preguntas").Drop(nil)
	fmt.Println("Borrando la Base de Datos Atlas.")

	err = atlas.InsertMany(allDB)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted multiple documents on Atlas.")
}
func leyes(norma int) (string, string, string, string) {
	var (
		categoria string
		tema      string
		titulo    string
		ley       string
	)
	for _, v := range normas {
		if norma == v.ID {
			categoria = v.Categoria
			tema = v.Tema
			titulo = v.Titulo
			ley = v.Ley
		}
	}
	return categoria, tema, titulo, ley
}
