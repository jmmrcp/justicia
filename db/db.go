package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db         *sql.DB
	ID         int
	Tema       int
	Pregunta   string
	Respuesta1 string
	Respuesta2 string
	Respuesta3 string
	Respuesta4 string
	Articulo   string
	Fecha      time.Time
	Box        int
	err        error
	rows       *sql.Rows
)

//Read -- Parses db file
func Read(path string, records [][]string, view int, test int) ([][]string, error) {

	//Make sure the file exists
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
		err = fmt.Errorf("file: %s does not exist", path)
		return records, err
	}
	//Open the db
	db, err = sql.Open("sqlite3", path)
	if err != nil {
		return records, err
	}
	// test connection
	err = db.Ping()
	if err != nil {
		return records, err
	}
	defer db.Close()

	//Read the database
	if test != 0 {
		rows, err = db.Query("SELECT * FROM just WHERE tema = ?", test)
			if err != nil {
				return records, err
			}
	} else {
		switch view {
		case 1:
			rows, err = db.Query("SELECT * FROM semana")
			if err != nil {
				return records, err
			}
		case 2:
			rows, err = db.Query("SELECT * FROM quincena")
			if err != nil {
				return records, err
			}
		case 3:
			rows, err = db.Query("SELECT * FROM mes")
			if err != nil {
				return records, err
			}
		default:
			rows, err = db.Query("SELECT * FROM dia")
			if err != nil {
				return records, err
			}
		}
	}
	//counter
	for rows.Next() {
		err = rows.Scan(&ID, &Tema, &Pregunta, &Respuesta1, &Respuesta2, &Respuesta3, &Respuesta4, &Articulo, &Fecha, &Box)
		id := strconv.Itoa(ID)
		record := []string{Pregunta, Respuesta1, Respuesta2, Respuesta3, Respuesta4, Articulo, id}
		records = append(records, record)
	}
	rows.Close()
	return records, nil
}

// Update db
func Update(id int) error {

	db, err = sql.Open("sqlite3", "data/data.db")
	if err != nil {
		return err
	}
	// test connection
	err = db.Ping()
	if err != nil {
		return err
	}
	defer db.Close()

	_, err := db.Exec("UPDATE just SET box = ? WHERE id = ?", 1, id)
	if err != nil {
		return err
	}
	return nil
}
