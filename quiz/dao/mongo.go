package dao

import (
	"context"
	"fmt"
	"justicia/quiz/config"
	"justicia/quiz/models"
	"time"
)

const (
	// COLLECTION Document TABLE
	COLLECTION = "preguntas"
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
	db, err := config.GetAtlasDB()
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
	c := db.Collection(COLLECTION)
	cursor, err := c.Find(ctx, Filter)
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
	db, err := config.GetAtlasDB()
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
	c := db.Collection(COLLECTION)

	for _, id := range ids {
		filter := IDs(id)
		updateResult, err := c.UpdateOne(ctx, filter, update)
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
func Unupdate(id string) error {

	filter := IDs(id)
	update := Wrong

	db, err := config.GetAtlasDB()
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

	c := db.Collection(COLLECTION)
	// updateResult
	_, err = c.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	// fmt.Printf("Wrong Matched %v documents and unupdated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	err = db.Client.Disconnect(ctx)
	if err != nil {
		return err
	}
	// fmt.Println("Connection to MongoDB closed.")
	return nil
}
