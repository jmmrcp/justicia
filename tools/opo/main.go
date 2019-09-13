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
)

func main() {

	normasJSON := `[
	{"id": 1, "nombre": "Constitución Española", "categoria":"CE" , "tema":"1", "ley": "CE"},
	{"id": 2, "nombre": "Estatuto Básico del Empleado Público", "categoria": "EBEP", "tema":"7"},
	{"id": 3, "nombre": "Estatuto de los Trabajadores"},
	{"id": 4, "nombre": "IV Convenio Único Administración/Personal Laboral"},
	{"id": 5, "nombre": "Ley de Medidas para la reforma de la Función Pública"},
	{"id": 6, "nombre": "Ley del Procedimiento Administrativo Común de las AAPP"},
	{"id": 7, "nombre": "Ley de Bases del Régimen Local"},
	{"id": 9, "nombre": "Ley de Dependencia"},
	{"id": 10, "nombre": "Ley de Extranjería"},
	{"id": 11, "nombre": "Ley de Igualdad"},
	{"id": 12, "nombre": "Ley de Incompatibilidades"},
	{"id": 14, "nombre": "Ley de Subvenciones"},
	{"id": 15, "nombre": "Ley del Gobierno"},
	{"id": 16, "nombre": "Ley General Presupuestaria"},
	{"id": 17, "nombre": "Ley General Tributaria"},
	{"id": 18, "nombre": "Ley Orgánica Violencia contra la mujer"},
	{"id": 19, "nombre": "Real Decreto 364/1995 Acceso a la función Pública"},
	{"id": 20, "nombre": "Real Decreto 365/1995 Situaciones Administrativas"},
	{"id": 21, "nombre": "Ley General de Sanidad"},
	{"id": 22, "nombre": "Estatuto Marco Personal de los Servicios de Salud"},
	{"id": 23, "nombre": "Ley de Prevención de Riesgos Laborales"},
	{"id": 24, "nombre": "Ley Orgánica del Poder Judicial", "categoria":"LOPJ", "team":"8", "Ley": "LOPJ"},
	{"id": 25, "nombre": "Ley Orgánica Fuerzas y Cuerpos Seguridad del Estado"},
	{"id": 26, "nombre": "Código Penal"},
	{"id": 27, "nombre": "Ley de Enjuiciamiento Civil", "categoria":"CIVIL", "tema":"14"},
	{"id": 28, "nombre": "Ley de Transparencia"},
	{"id": 29, "nombre": "Ley de Régimen Jurídico del Sector Público"},
	{"id": 30, "nombre": "Ley Orgánica General Penitenciaria"},
	{"id": 31, "nombre": "Reglamento Penitenciario"},
	{"id": 32, "nombre": "Ley Jurisdicción Contencioso-Administrativa", "categoria":"CA", "tema":"22"},
	{"id": 33, "nombre": "Ley Orgánica de Protección de Datos Personales y garantía de los derechos digitales"},
	{"id": 34, "nombre": "Ley General de la Seguridad Social"},
	{"id": 35, "nombre": "Estatuto de Autonomía de la Comunidad Valenciana"},
	{"id": 36, "nombre": "Estatuto de Autonomía de la Comunidad de Madrid"},
	{"id": 37, "nombre": "Estatuto de Autonomía de Andalucía"},
	{"id": 38, "nombre": "Estatuto de Autonomía de Cataluña"},
	{"id": 39, "nombre": "Declaración Universal de los Derechos Humanos"},
	{"id": 40, "nombre": "Ley de Haciendas Locales"},
	{"id": 41, "nombre": "Ley de Enjuiciamiento Criminal", "categoria":"PENAL", "tema":"20,21", "ley": "LeCrim"},
	{"id": 42, "nombre": "Ley de Contratos del Sector Público"},
	{"id": 43, "nombre": "Reglamento Europeo de Protección de Datos de las personas físicas"},
	{"id": 44, "nombre": "Ley Orgánica de Universidades"},
	{"id": 45, "nombre": "Estatuto de Autonomía de Galicia"},
	{"id": 46, "nombre": "Estatuto de Autonomía de Castilla y León"},
	{"id": 47, "nombre": "Ley Orgánica de Estabilidad Presupuestaria y Sostenibilidad Financiera"},
	{"id": 48, "nombre": "Tratado de la Unión Europea", "categoria":"EUROPA", "tema":"5", "ley": "TUE"}
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
	test := 999

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
