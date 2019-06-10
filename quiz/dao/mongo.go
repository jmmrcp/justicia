package dao

import (
	"context"
	"fmt"
	"justicia/quiz/config"
	"justicia/quiz/models"
	"time"

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
		filter    bson.D
		questions [][]string
	)
	today := time.Now()
	//Open the db
	db, err := config.GetMongoDB()
	if err != nil {
		return nil, err
	}

	// categoty filter
	if cat != "" {
		filter = bson.D{
			{"categoria", cat},
		}
	}
	// test filter
	if test != 0 {
		filter = bson.D{
			{"test", test},
		}
	} else {
		// View filter
		switch view {
		case 1:
			date := today.AddDate(0, 0, -7)
			filter = bson.D{
				{"fecha", bson.D{
					{"$gt", date},
				}},
				{"box", 1},
			}
		case 2:
			date := today.AddDate(0, 0, -14)
			filter = bson.D{
				{"fecha", bson.D{
					{"$gt", date},
				}},
				{"box", 2},
			}
		case 3:
			date := today.AddDate(0, 0, -28)
			filter = bson.D{
				{"fecha", bson.D{
					{"$gt", date},
				}},
				{"box", 3},
			}
		default:
			filter = bson.D{}
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
		return err
	}
	fmt.Printf("ID: %v\n", v)
	db, err := config.GetMongoDB()
	if err != nil {
		return err
	}
	// Check the connection
	err = db.Client().Ping(context.TODO(), nil)
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
		{"$set", bson.D{
			{"fecha", time.Now()},
		}},
	}
	c := db.Collection(COLLECTION)
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	err = db.Client().Disconnect(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}

// Unupdate Actualiza el contenido
func Unupdate(id string) error {
	v, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	fmt.Printf("ID: %v\n", v)
	db, err := config.GetMongoDB()
	if err != nil {
		return err
	}
	// Check the connection
	err = db.Client().Ping(context.TODO(), nil)
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
	updateResult, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	fmt.Printf("Matched %v documents and unupdated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	err = db.Client().Disconnect(context.TODO())
	if err != nil {
		return err
	}
	fmt.Println("Connection to MongoDB closed.")
	return nil
}
