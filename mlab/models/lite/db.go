package lite

import (
	"database/sql"
	// Importacion de sqlite
	_ "github.com/mattn/go-sqlite3"
)

type (
	// Datastore es el almacen para toda la base de datos
	Datastore interface {
		All() ([]*Question, error)
	}
	// DB Base de datos SQLITE
	DB struct {
		*sql.DB
	}
)

//NewDB Crea la conexion con la base de datos
func NewDB() (*DB, error) {
	db, err := sql.Open("sqlite3", "data.db")
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DB{db}, nil
}
