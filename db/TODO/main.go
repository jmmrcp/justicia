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
	spc        string
	t          string
)

func main() {

	//Check passed in files
	if len(os.Args) < 2 {
		fmt.Println("Please set a file name.")
		return
	}
	filename := os.Args[1]

	file, err := os.Open(filename + ".txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Printf("Categoria del Test: ")
	fmt.Scanf("%s", &t)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		quest(line)
	}
	p := len(preguntas)
	r := len(respuestas)

	if r/p == 4 {
		fmt.Println("Proceso correcto.")
	} else {
		os.Exit(1)
	}

	db(preguntas, respuestas, filename)
}

func quest(line string) {
	str := strings.Split(line, ") ")
	if len(str) > 1 {
		_, err := strconv.Atoi(str[0])
		if err != nil {
			// Respuestas
			respuestas = append(respuestas, str[1])
		} else {
			// Preguntas
			preguntas = append(preguntas, str[1])
		}
	}
}

func db(preguntas []string, respuestas []string, filename string) {
	var (
		db    *sql.DB
		err   error
		letra string
		R1    string
		R2    string
		R3    string
		R4    string
		A1    string
	)

	test, _ := strconv.Atoi(filename)

	db, err = sql.Open("sqlite3", "data.db")
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
		test       INTEGER NOT NULL
							 DEFAULT 1,
		tema			 TEXT		 NOT NULL,
		pregunta   TEXT    NOT NULL,
		respuesta1 TEXT    NOT NULL,
		respuesta2 TEXT    NOT NULL,
		respuesta3 TEXT    NOT NULL,
		respuesta4 TEXT    NOT NULL,
		articulo   TEXT    NOT NULL,
		ord				 INTEGER NOT NULL,
		fecha      DATE    DEFAULT CURRENT_TIMESTAMP,
		box        INTEGER DEFAULT 0
	);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS temas ON just (
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
           tema,
           pregunta,
           respuesta1,
           respuesta2,
           respuesta3,
           respuesta4,
           articulo,
					 ord,
					 fecha,
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
           tema,
           pregunta,
           respuesta1,
           respuesta2,
           respuesta3,
           respuesta4,
           articulo,
					 ord,
					 fecha,
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
           tema,
           pregunta,
           respuesta1,
           respuesta2,
           respuesta3,
           respuesta4,
           articulo,
					 ord,
					 fecha,
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
           tema,
           pregunta,
           respuesta1,
           respuesta2,
           respuesta3,
           respuesta4,
           articulo,
					 ord,
					 fecha,
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
       SET fecha = CURRENT_TIMESTAMP
     WHERE id = OLD.id;
END;
	`)
	if err != nil {
		log.Fatal(err)
	}

	art, err := os.Open(filename + ".art")
	if err != nil {
		log.Fatal(err)
	}
	defer art.Close()

	sol, err := os.Open(filename + ".sol")
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
		_, err = db.Exec(`
    INSERT INTO just (
			test,
      tema, 
      pregunta, 
      respuesta1, 
      respuesta2, 
      respuesta3, 
      respuesta4, 
			articulo,
			ord
      ) 
      VALUES (?,?,?,?,?,?,?,?,?)`, test, &t, P, R1, R2, R3, R4, A1, i+1)
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
		err = rows.Scan(&ID, &Test, &Tema, &Pregunta, &Respuesta1, &Respuesta2, &Respuesta3, &Respuesta4, &Articulo, &Ord, &Fecha, &Box)
		id := strconv.Itoa(ID)
		t := strconv.Itoa(Test)
		o := strconv.Itoa(Ord)
		base := []string{id, t, Tema, Pregunta, Respuesta1, Respuesta2, Respuesta3, Respuesta4, Articulo, o}
		data = append(data, base)
	}
	fmt.Printf("%v\n", data)
	rows.Close()
}

func mlab() {
	// user - jmmrcp
	// pass - J538MTUSbg3v3Vh
	// conect mongodb://<dbuser>:<dbpassword>@ds263876.mlab.com:63876/justicia

	_, err := mgo.Dial("mongodb://jmmrcp:J538MTUSbg3v3Vh@ds263876.mlab.com:63876/justicia")
	if err != nil {
		log.Fatal(err)
	}
}
