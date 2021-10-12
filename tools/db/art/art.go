package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	art, err := os.Open("sample.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer art.Close()
	articulos := bufio.NewScanner(art)

	for articulos.Scan() {
		articulo := (articulos.Text())
		separate(articulo)
	}
}

func separate(articulo string) {
	var (
		art []string
		ley string
	)
	fmt.Println("********** < ARTICULO: " + articulo + " > **********")
	// comprtobamos si es compuesto
	if strings.ContainsAny(articulo, ",=-y") {
		if strings.Contains(articulo, ",") {
			fmt.Println("********** Contiene comas **********")
			// Separamos la ley de los articulos
			l := strings.Split(articulo, " ")
			// La ley debe de ser el ultimo de la lista
			z := len(l) - 1
			// Separamos por comas y comprobamos
			m := strings.Split(articulo, ",")
			for _, v := range m {
				if strings.Contains(v, " ") {
					// Pueden ser varias leyes
					if len(strings.Split(v, " ")) == 2 {
						// el formato se corresponde
					} else {
						// No lo es
					}

				}
			}
			a := strings.Split(l[0], ",")
			// b := strings.Split(m, " ")
			art = append(art, a...)
			ley = l[z]
		}
		if strings.Contains(articulo, "=") {
			fmt.Println("Contiene iguales")
		}
		if strings.Contains(articulo, "-") {
			fmt.Println("********** Contiene guiones **********")
			// Separamos la ley de los articulos
			l := strings.Split(articulo, " ")
			z := len(l) - 1
			a := strings.Split(l[0], "-")
			art = append(art, a...)
			ley = l[z]
		}
		if strings.Contains(articulo, "y") {
			fmt.Println("********** Contiene la letra y **********")
			l := strings.Split(articulo, " ")
			z := len(l) - 1
			a := strings.Split(l[0], "y")
			art = append(art, a...)
			ley = l[z]
		}
	} else {
		// fmt.Println("Normal")
		l := strings.Split(articulo, " ")
		z := len(l) - 1
		art = append(art, l[0])
		ley = l[z]
	}
	fmt.Printf("Articulo: %s\nLey: %s\n", art, ley)
}

func leyNormal(articulo string) {

}
