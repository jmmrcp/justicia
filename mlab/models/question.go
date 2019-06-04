package models

import "time"

// Question Schema de la base de datos
type Question struct {
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
}

// All funcion para recuperar todo el contenido de la BD.
func (db *DB) All() ([]*Question, error) {
	rows, err := db.Query("SELECT * FROM just")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	questions := make([]*Question, 0)
	for rows.Next() {
		question := new(Question)
		err := rows.Scan(
			&question.ID,
			&question.Test,
			&question.Tema,
			&question.Pregunta,
			&question.Respuesta1,
			&question.Respuesta2,
			&question.Respuesta3,
			&question.Respuesta4,
			&question.Articulo,
			&question.Ord,
			&question.Fecha,
			&question.Box,
		)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}
