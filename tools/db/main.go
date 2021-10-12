// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var (
	// Total Suna de todas las preguntas
	Total int
	// Key to encrypt Question and Answer One
	Key string

	preguntas  []string
	respuestas []string
	info       []string
)

func main() {
	// Borrado de la base de datoa
	err := os.Remove("../mongo/data.db")
	file, err := os.Open("test.nfo")
	if err != nil {
		log.Fatal("Base de datos no existe o no se puede borrar", err)
	}
	defer file.Close()
	// Leemos el fichero test.nfo
	key()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Si la linea esta vacio la saltamos
		if len(line) < 1 {
			fmt.Println()
			fmt.Printf("%q\n", info)
			txt(info)
			// Vaciamos las variables
			info = nil
			preguntas = nil
			respuestas = nil
		} else {
			info = append(info, line)
		}
	}
	fmt.Println("Numero de Pregunta procesadas: ", Total)
	fmt.Printf("key to encrypt/decrypt : %s\n", Key)
}

// txt -> []string
// return la base de datos terminada
func txt(info []string) {
	var (
		// filename -> nombre del fichero definido en el fichero nfo.
		filename = info[0]
		// total -> numero de preguntas definido en el fichero nfo.
		total, _ = strconv.Atoi(info[4])
	)
	// Leemoos el fichero de la carpeta txt
	file, err := os.Open("txt/" + filename + ".txt")
	if err != nil {
		log.Fatal("No se ha podido leer el fichero.", err)
	}
	defer file.Close()
	// Empezamos a leer las lineas del test
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Procesamos la linea
		quest(line)
	}
	// Numero de preguntas en el Test
	p := len(preguntas)
	// Numero de respuestas en el Test
	r := len(respuestas)
	// Mostramos las preguntas leidas.
	fmt.Println("Numero de Preguntas del test: ", p)
	// Comprobamos que todas las lineas se han leido correctamente
	// diviendo el total de respuestas entre cuatro e igualandalo al numero de preguntas
	if r/p == 4 && p == total {
		fmt.Println("Preproceso correcto.")
		// Procesamos el Test
		db(preguntas, respuestas)
		Total += p
	} else {
		fmt.Println("Numero de preguntas o respuestas incorrecto.")
		fmt.Println("Respuestas: ", r)
		// Salimos del programa
	}
}

// quest -> string
// añade a las variables las preguntas o las respuestas
func quest(line string) {
	// Separamos entre por el parentesis
	str := strings.Split(line, ") ")
	// Si tiene dos partes es una pregunta o respuesta
	if len(str) > 1 {
		// la primera parte la convertimos a entero
		_, err := strconv.Atoi(str[0])
		// Si da error es Letra -> Respuesta a,b,c,d.
		if err != nil {
			// Respuestas
			txt := str[1]
			respuestas = append(respuestas, txt)
			// Si es entero es Numero -> Pregunta 1,2,3,4,5 ....
		} else {
			// Preguntas
			txt := str[1]
			preguntas = append(preguntas, txt)
		}
	}
}

// db -> []preguntas, []respuestas
// Return -> La base datos con los datos del Test
func db(preguntas []string, respuestas []string) {

	fmt.Printf("key to encrypt/decrypt : %s\n", Key)
	// Definimos todas las variables
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
	// Convertimos el nombre del Test a entero para la DB
	TE, _ := strconv.Atoi(filename)
	// Abrimos la base de datos
	db, err = sql.Open("sqlite3", "../mongo/data.db")
	if err != nil {
		log.Fatal("No podemos abrir la base de datos.", err)
	}
	defer db.Close()

	// test connection
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// Iniciamos las transaciones.
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	// Creamos la base de datos si no existe.
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
		log.Fatal("No se ha podico crear la base de datos", err)
	}
	// Creamos el indice de temas.
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS categoria ON just (
    tema ASC
);`)
	if err != nil {
		log.Fatal("No se ha podico crear el indice de temas.", err)
	}
	// Creamos el indice de test.
	_, err = db.Exec(`CREATE INDEX IF NOT EXISTS tests ON just (
    test ASC
);`)
	if err != nil {
		log.Fatal("No de ha podido crear el indice de tests.", err)
	}
	// Creamos la Vista de Dia.
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
	// Creamos la vista de semana.
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
	// Creamos la vista de quincena.
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
	// Creamos la vista de mes.
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
	// Creamos un disparador
	// Cuando actualizamos un registro lo hacemos a fecha de la actualizacion
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
	// Leemos el fichero de los articulos de las leyes.
	art, err := os.Open("art/" + filename + ".art")
	if err != nil {
		log.Fatal(err)
	}
	defer art.Close()
	// Leemos el fichero de las soluciones.
	sol, err := os.Open("sol/" + filename + ".sol")
	if err != nil {
		log.Fatal(err)
	}
	defer sol.Close()
	// Procesamos ambos archivos
	articulos := bufio.NewScanner(art)
	soluciones := bufio.NewScanner(sol)

	for i, P := range preguntas {
		// Linea de la Ley
		articulos.Scan()
		// Linea de la solucion correcta
		soluciones.Scan()
		// Separamos la linea
		solucion := strings.Split(soluciones.Text(), ". ")
		if len(solucion) > 1 {
			// Si el archivo de soluciones tiene el numero de pregunta y la solucion
			letra = strings.ToUpper(solucion[1])
		} else {
			// Si solo tiene la solucion
			letra = strings.ToUpper(solucion[0])
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
		default:
			log.Fatalf("Datos Incorrectos en: - %v", i*4)
		}
		P = strings.TrimRight(P, "-1234567890modAGET ")
		R1 = strings.TrimRight(R1, " .")
		R2 = strings.TrimRight(R2, " .")
		R3 = strings.TrimRight(R3, " .")
		R4 = strings.TrimRight(R4, " .")

		EP := encrypt(P, Key)
		ER1 := encrypt(R1, Key)

		// INSERT OR IGNORE INTO just (
		// Insertamos el registro en la base de datos
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
      VALUES (?,?,?,?,?,?,?,?,?,?,?)`, TE, C, T, TI, EP, ER1, R2, R3, R4, A1, i+1)
		if err != nil {
			log.Fatal(err)
		}
	}
	fmt.Println("Proceso correcto.")
	tx.Commit()
}

func encrypt(stringToEncrypt string, keyString string) (encryptedString string) {

	//Since the key is in string, we need to convert decode it to bytes
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(stringToEncrypt)

	//Create a new Cipher Block from the key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	//Create a new GCM - https://en.wikipedia.org/wiki/Galois/Counter_Mode
	//https://golang.org/pkg/crypto/cipher/#NewGCM
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	//Create a nonce. Nonce should be from GCM
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		panic(err.Error())
	}

	//Encrypt the data using aesGCM.Seal
	//Since we don't want to save the nonce somewhere else in this case, we add it as a prefix to the encrypted data. The first nonce argument in Seal is the prefix.
	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return fmt.Sprintf("%x", ciphertext)
}

func key() {
	bytes := make([]byte, 32) //generate a random 32 byte key for AES-256
	if _, err := rand.Read(bytes); err != nil {
		panic(err.Error())
	}

	Key = hex.EncodeToString(bytes) //encode key in bytes to string and keep as secret, put in a vault
}
