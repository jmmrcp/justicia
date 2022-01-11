// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package dao

import (
	"context"
	"fmt"
	"justicia/quiz/config"
	"justicia/quiz/models"
	"time"
)

// Read Lee todos los registros de la base de datos
func Read(records [][]string, view, test, tema int, cat string) ([][]string, error) {
	var (
		questions [][]string
	)

	// test filter
	if test != 0 {
		Filter = Test(test)
	} else {
		// View filter
		switch view {
		case 1:
			Filter = StageOne
		case 2:
			Filter = StageTwo
		case 3:
			Filter = StageThree
		default:
			if cat != "" || tema != 0 {
				if cat != "" {
					Filter = Category(cat)
				}
				if tema != 0 {
					Filter = Tema(tema)
				}
			} else {
				return Quick()
			}
		}
	}

	// categoty filter

	//Open the db
	db, err := config.GetMongoDB()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(db.Context, 3*time.Second)
	defer cancel()

	err = db.Client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	// Cursor Results
	cursor, err := db.Collection.Find(ctx, Filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Next result
	for cursor.Next(ctx) {
		var m *models.Mlab
		if err := cursor.Decode(&m); err != nil {
			return nil, err
		}
		q := m.Parse()
		questions = append(questions, q)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	if err = db.Client.Disconnect(ctx); err != nil {
		return nil, err
	}
	return questions, nil
}

// Update Actualiza el contenido
func Update(ids []string) error {

	update := Correct
	db, err := config.GetMongoDB()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(db.Context, 3*time.Second)
	defer cancel()
	// Check the connection
	err = db.Client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	for _, id := range ids {
		filter := IDs(id)
		updateResult, err := db.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
		fmt.Printf("Correct Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}
	err = db.Client.Disconnect(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}

// Unupdate Actualiza el contenido
func Unupdate(ids []string) error {
	update := Wrong
	db, err := config.GetMongoDB()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(db.Context, 3*time.Second)
	defer cancel()
	// Check the connection
	err = db.Client.Ping(ctx, nil)
	if err != nil {
		return err
	}
	// updateResult
	for _, id := range ids {
		filter := IDs(id)
		updateResult, err := db.Collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return err
		}
		fmt.Printf("Wrong Matched %v documents and unupdated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	}
	err = db.Client.Disconnect(ctx)
	if err != nil {
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}
