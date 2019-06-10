package dao

import (
	"context"
	"fmt"
	"justicia/quiz/config"
	"justicia/quiz/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// COLLECTION Document TABLE
	COLLECTION = "preguntas"
)

// Read Lee todos los registros de la base de datos
func Read(records [][]string, view int, test int, cat string) ([][]string, error) {
	var (
		filter    bson.M
		questions [][]string
	)
	//Open the db
	db, err := config.GetMongoDB()
	if err != nil {
		return nil, err
	}

	// categoty filter
	if cat != "" {
		filter = bson.M{
			"categoria": cat,
		}
	}
	// test filter
	if test != 0 {
		filter = bson.M{
			"test": test,
		}
	} else {
		// View filter
		switch view {
		case 1:
			filter = bson.M{
				"box": 1,
			}
		case 2:
			filter = bson.M{
				"box": 2,
			}
		case 3:
			filter = bson.M{
				"box": 3,
			}
		default:
			filter = bson.M{}
		}
	}

	// Cursor Results
	c := db.Collection(COLLECTION)
	ctx := context.TODO()
	cursor, err := c.Find(ctx, filter)
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

	return questions, nil
}

// Update Actualiza el contenido
func Update(id string) error {
	v, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	db, err := config.GetMongoDB()
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": v,
	}
	update := bson.D{
		{"$inc", bson.D{
			{"box", 1},
		}},
	}
	c := db.Collection(COLLECTION)
	s, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println(s.ModifiedCount)
	return nil
}

// Unupdate Actualiza el contenido
func Unupdate(id string) error {
	v, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Fatal(err)
	}
	db, err := config.GetMongoDB()
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": v,
	}
	update := bson.D{
		{"$set", bson.D{
			{"box", 0},
		}},
	}
	c := db.Collection(COLLECTION)
	s, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Println(s.ModifiedCount)
	return nil
}
