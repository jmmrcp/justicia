package lite

import "time"

// Question Schema de la base de datos
type Question struct {
	ID         int
	Test       int
	Categoria  string
	Tema       string
	Titulo     string
	Pregunta   string
	Respuesta1 string
	Respuesta2 string
	Respuesta3 string
	Respuesta4 string
	Articulo   string
	Ord        int
	Contador   int
	Fecha      time.Time
	Box        int
}

// All funcion para recuperar todo el contenido de la BD.
func (db *DB) All() ([]*Question, error) {
	rows, err := db.Query("SELECT * FROM just ORDER BY test;")
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
			&question.Categoria,
			&question.Tema,
			&question.Titulo,
			&question.Pregunta,
			&question.Respuesta1,
			&question.Respuesta2,
			&question.Respuesta3,
			&question.Respuesta4,
			&question.Articulo,
			&question.Ord,
			&question.Fecha,
			&question.Contador,
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
