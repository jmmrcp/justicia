package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	preguntas  []string
	respuestas []string
	info       []string
)

func main() {
	err := os.Remove("../mongo/data.db")
	file, err := os.Open("test.nfo")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) < 1 {
			fmt.Printf("%q\n", info)
			txt(info)
			info = nil
			preguntas = nil
			respuestas = nil
		} else {
			info = append(info, line)
		}
	}

}
func txt(info []string) {
	var (
		filename = info[0]
		total, _ = strconv.Atoi(info[4])
	)
	file, err := os.Open("txt/" + filename + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		quest(line)
	}
	p := len(preguntas)
	r := len(respuestas)

	fmt.Println("Numero de Preguntas del test: ", p)

	if r/p == 4 && p == total {
		fmt.Println("Proceso correcto.")
	} else {
		os.Exit(1)
	}
	db(preguntas, respuestas)
}

func quest(line string) {
	str := strings.Split(line, ") ")
	if len(str) > 1 {
		_, err := strconv.Atoi(str[0])
		if err != nil {
			// Respuestas
			txt := str[1]
			respuestas = append(respuestas, txt)
		} else {
			// Preguntas
			txt := str[1]
			preguntas = append(preguntas, txt)
		}
	}
}

func db(preguntas []string, respuestas []string) {
	var (
		filename = info[0]
		C        = info[1]
		T        = info[2]
		TI       = info[3]
		db       *sql.DB
		err      error
		letra    string
		R1       string
		R2       string
		R3       string
		R4       string
		A1       string
	)

	TE, _ := strconv.Atoi(filename)

	db, err = sql.Open("sqlite3", "../mongo/data.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS just (
		id         INTEGER NOT NULL
						   PRIMARY KEY AUTOINCREMENT
						   UNIQUE,
		test       INTEGER NOT NULL,
		categoria  TEXT	   NOT NULL,
		tema	     TEXT	   NOT NULL,
		titulo	   TEXT	   NOT NULL,
		pregunta   TEXT    NOT NULL,
		respuesta1 TEXT    NOT NULL,
		respuesta2 TEXT    NOT NULL,
		respuesta3 TEXT    NOT NULL,
		respuesta4 TEXT    NOT NULL,
		articulo   TEXT    NOT NULL,
		ord        INTEGER NOT NULL,
		fecha      DATE    DEFAULT CURRENT_TIMESTAMP,
		cont       INTEGER DEFAULT 0,
		box        INTEGER DEFAULT 0
	);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS categoria ON just (
    tema ASC
);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS tests ON just (
    test ASC
);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`
  CREATE VIEW IF NOT EXISTS dia AS
		SELECT id,
		test,
		categoria,
		tema,
		titulo, 
		pregunta, 
		respuesta1, 
		respuesta2, 
		respuesta3, 
		respuesta4, 
		articulo,
		ord,
		fecha,
		cont,
		box
      FROM just
		 WHERE box = 0;`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
  CREATE VIEW IF NOT EXISTS semana AS
		SELECT id,
		test,
		categoria,
		tema,
		titulo, 
		pregunta, 
		respuesta1, 
		respuesta2, 
		respuesta3, 
		respuesta4, 
		articulo,
		ord,
		fecha,
		cont,
		box
      FROM just
	 WHERE box = 1;
	 `)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
  CREATE VIEW IF NOT EXISTS quincena AS
		SELECT id,
		test,
		categoria,
		tema,
		titulo, 
		pregunta, 
		respuesta1, 
		respuesta2, 
		respuesta3, 
		respuesta4, 
		articulo,
		ord,
		fecha,
		cont,
		box
      FROM just
	 WHERE box = 2;
	 `)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
  CREATE VIEW IF NOT EXISTS mes AS
		SELECT id,
		test,
		categoria,
		tema,
		titulo, 
		pregunta, 
		respuesta1, 
		respuesta2, 
		respuesta3, 
		respuesta4, 
		articulo,
		ord,
		fecha,
		cont,
		box
      FROM just
	 WHERE box = 3;
	 `)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`
  CREATE TRIGGER IF NOT EXISTS actualiza
         AFTER UPDATE
            ON just
BEGIN
    UPDATE just
       SET fecha = CURRENT_TIMESTAMP, cont = cont + 1
     WHERE id = OLD.id;
END;
	`)
	if err != nil {
		log.Fatal(err)
	}

	art, err := os.Open("art/" + filename + ".art")
	if err != nil {
		log.Fatal(err)
	}
	defer art.Close()

	sol, err := os.Open("sol/" + filename + ".sol")
	if err != nil {
		log.Fatal(err)
	}
	defer sol.Close()

	articulos := bufio.NewScanner(art)
	soluciones := bufio.NewScanner(sol)

	for i, P := range preguntas {
		articulos.Scan()
		soluciones.Scan()

		solucion := strings.Split(soluciones.Text(), ". ")
		if len(solucion) > 1 {
			letra = solucion[1]
		} else {
			letra = solucion[0]
		}

		switch letra {
		case "A":
			p := i * 4
			R1 = respuestas[p]
			R2 = respuestas[p+1]
			R3 = respuestas[p+2]
			R4 = respuestas[p+3]
			A1 = articulos.Text()
		case "B":
			p := i * 4
			R1 = respuestas[p+1]
			R2 = respuestas[p]
			R3 = respuestas[p+2]
			R4 = respuestas[p+3]
			A1 = articulos.Text()
		case "C":
			p := i * 4
			R1 = respuestas[p+2]
			R2 = respuestas[p+1]
			R3 = respuestas[p]
			R4 = respuestas[p+3]
			A1 = articulos.Text()
		case "D":
			p := i * 4
			R1 = respuestas[p+3]
			R2 = respuestas[p+1]
			R3 = respuestas[p+2]
			R4 = respuestas[p]
			A1 = articulos.Text()
		}
		P = strings.TrimRight(P, "-1234567890modAGET ")
		R1 = strings.TrimRight(R1, " .")
		R2 = strings.TrimRight(R2, " .")
		R3 = strings.TrimRight(R3, " .")
		R4 = strings.TrimRight(R4, " .")
		_, err = db.Exec(`
    INSERT INTO just (
			test,
			categoria,
			tema,
			titulo, 
      pregunta, 
      respuesta1, 
      respuesta2, 
      respuesta3, 
      respuesta4, 
			articulo,
			ord
      ) 
      VALUES (?,?,?,?,?,?,?,?,?,?,?)`, TE, C, T, TI, P, R1, R2, R3, R4, A1, i+1)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

func listar(db *sql.DB) {
	var (
		rows       *sql.Rows
		ID         int
		Test       int
		Tema       string
		Pregunta   string
		Respuesta1 string
		Respuesta2 string
		Respuesta3 string
		Respuesta4 string
		Articulo   string
		Ord        int
		Fecha      time.Time
		Cont       int
		Box        int
		data       [][]string
		err        error
	)
	rows, err = db.Query(`SELECT * FROM just`)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		fmt.Println(rows)
		err = rows.Scan(&ID, &Test, &Tema, &Pregunta, &Respuesta1, &Respuesta2, &Respuesta3, &Respuesta4, &Articulo, &Ord, &Fecha, &Cont, &Box)
		id := strconv.Itoa(ID)
		t := strconv.Itoa(Test)
		o := strconv.Itoa(Ord)
		base := []string{id, t, Tema, Pregunta, Respuesta1, Respuesta2, Respuesta3, Respuesta4, Articulo, o}
		data = append(data, base)
	}
	fmt.Printf("%v\n", data)
	rows.Close()
}