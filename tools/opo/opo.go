package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
		Tema       string
		Titulo     string
		Norma      int    `json:"norma"`
		Pregunta   string `json:"enunciado"`
		Respuesta1 string `json:"a"`
		Respuesta2 string `json:"b"`
		Respuesta3 string `json:"c"`
		Respuesta4 string `json:"d"`
		Correcta   string `json:"correcta"`
		Articulo   string
		Ord        int
		Fecha      time.Time
		Cont       int
		Box        int
	}

	// Normas leyes incluidas en el JSON
	Normas struct {
		ID        int    `json:"id"`
		Titulo    string `json:"nombre"`
		Categoria string `json:"categoria"`
		Tema      string `json:"tema"`
		Ley       string `json:"ley"`
	}
)

var (
	preguntas []Pregunta
	normas    []Normas
)

func main() {

	normasJSON, err := os.Open("normas.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Opened normas file.")
	defer normasJSON.Close()

	byteValue, _ := ioutil.ReadAll(normasJSON)
	json.Unmarshal(byteValue, &normas)

	jsonFile, err := os.Open("preguntas.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Opened questions file.")
	defer jsonFile.Close()

	byteValue, _ = ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &preguntas)

	contador := 0
	test := 400

	for i := 0; i < len(preguntas)-1; i++ {
		actual := preguntas[i].Norma
		siguiente := preguntas[i+1].Norma

		categoria, tema, titulo, ley := leyes(actual)
		t := strings.Split(preguntas[i].Pregunta, "(Art. ")
		preguntas[i].Pregunta = t[0]
		preguntas[i].Test = test
		preguntas[i].Categoria = categoria
		preguntas[i].Tema = tema
		preguntas[i].Titulo = titulo
		if len(t) > 1 {
			a := t[1]
			b := len(a) - 1
			c := a[:b]
			preguntas[i].Articulo = "Art. " + c + " " + ley
		}
		contador++
		preguntas[i].Ord = contador
		preguntas[i].Fecha = time.Now()
		if actual != siguiente {
			contador = 0
			test++
		}
	}
	fmt.Printf("%+v\n", preguntas[6220])
	fmt.Printf("%+v\n", preguntas[6221])
	fmt.Printf("%+v\n", preguntas[6222])
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
