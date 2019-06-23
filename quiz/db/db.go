package db

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	// Sqlite 3 Import
	_ "github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
	// ID -> Identicador unico
	ID int
	// Test -> numero de Test
	Test int
	// Tema -> Categora de Test
	Tema string
	// Pregunta -> Pregunta de Test
	Pregunta string
	// Respuesta1 -> Respuesta correcta
	Respuesta1 string
	// Respuesta2 -> Respuesta falsa
	Respuesta2 string
	// Respuesta3 -> Otra falsa
	Respuesta3 string
	// Respuesta4 -> Otra de las respustas falsas
	Respuesta4 string
	// Articulo -> Art. de la Ley donde se apoya la respuesta correcta
	Articulo string
	// Ord -> orden original en el Test para modificaciones
	Ord int
	// Fecha -> Fecha de inclusion -> modifacion / fecha de acierto o fallo
	Fecha time.Time
	// Box -> Contenedor de respuetas acertadas a 7, 14, 28 dias
	Box  int
	err  error
	rows *sql.Rows
)

//Read -- Parses db file
func Read(path string, records [][]string, view int, test int, cat string) ([][]string, error) {

	//Make sure the file exists
	_, err = os.Stat(path)
	if os.IsNotExist(err) {
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

	sqlQuery := createQuery(test, view, cat)

	//Read the database

	rows, err = db.Query(sqlQuery)
	if err != nil {
		return records, err
	}

	// if test != 0 {
	// 	rows, err = db.Query("SELECT * FROM just WHERE test = ?", test)
	// 	if err != nil {
	// 		return records, err
	// 	}
	// } else {
	// 	switch view {
	// 	case 1:
	// 		rows, err = db.Query("SELECT * FROM semana < date('now', '-7 days'")
	// 		if err != nil {
	// 			return records, err
	// 		}
	// 	case 2:
	// 		rows, err = db.Query("SELECT * FROM quincena < date('now', '-14 days'")
	// 		if err != nil {
	// 			return records, err
	// 		}
	// 	case 3:
	// 		rows, err = db.Query("SELECT * FROM mes < date ('now', '-28 days'")
	// 		if err != nil {
	// 			return records, err
	// 		}
	// 	default:
	// 		rows, err = db.Query("SELECT * FROM dia")
	// 		if err != nil {
	// 			return records, err
	// 		}
	// 	}
	// }
	// if cat != "" {
	// 	rows, err = db.Query("SELECT * FROM just WHERE tema = UPPER(?)", cat)
	// 	if rows != nil {
	// 		return records, err
	// 	}
	// }

	defer rows.Close()
	//counter
	for rows.Next() {
		err = rows.Scan(&ID, &Test, &Tema, &Pregunta, &Respuesta1, &Respuesta2, &Respuesta3, &Respuesta4, &Articulo, &Ord, &Fecha, &Box)
		id := strconv.Itoa(ID)
		record := []string{Pregunta, Respuesta1, Respuesta2, Respuesta3, Respuesta4, Articulo, id}
		records = append(records, record)
	}
	return records, nil
}

// Update db
func Update(id string) error {

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

	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE just SET box = box + 1 WHERE id = ?", i)
	if err != nil {
		return err
	}
	return nil
}

// Unupdate db
func Unupdate(id string) error {

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

	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE just SET box = 0 WHERE id = ?", i)
	if err != nil {
		return err
	}
	return nil
}

func createQuery(test, box int, tema string) string {
	var sqlQuery string
	t := strconv.Itoa(test)
	if tema != "" {
		sqlQuery = `SELECT * FROM just WHERE tema = UPPER(` + "\"" + tema + "\"" + `);`
		fmt.Println(sqlQuery)
		return sqlQuery
	}
	if test != 0 {
		sqlQuery = `SELECT * FROM just WHERE test = ` + t + `AND box = 0;`
		return sqlQuery
	}
	if box != 0 {
		switch box {
		case 1:
			sqlQuery = `SELECT * FROM semana WHERE fecha < date('now', '-7 days');`
			return sqlQuery
		case 2:
			sqlQuery = `SELECT * FROM quincena WHERE fecha < date('now', '-14 days');`
			return sqlQuery
		case 3:
			sqlQuery = `SELECT * FROM mes WHERE fecha < date('now', '-28 days');`
			return sqlQuery
		}
	}
	sqlQuery = `SELECT * FROM dia ORDER BY random() LIMIT 100;`
	return sqlQuery
}
