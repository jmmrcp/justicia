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

	// test filter
	if test != 0 {
		filter = bson.D{
			primitive.E{
				Key:   "test",
				Value: test},
		}
	} else {
		// View filter
		switch view {
		case 1:
			date := today.AddDate(0, 0, -7)
			filter = bson.D{
				primitive.E{
					Key: "fecha",
					Value: bson.D{
						primitive.E{
							Key:   "$lt",
							Value: date},
					}},
				primitive.E{
					Key:   "box",
					Value: 1,
				},
			}
		case 2:
			date := today.AddDate(0, 0, -14)
			filter = bson.D{
				primitive.E{
					Key: "fecha",
					Value: bson.D{
						primitive.E{
							Key:   "$lt",
							Value: date,
						},
					}},
				primitive.E{
					Key:   "box",
					Value: 2,
				},
			}
		case 3:
			date := today.AddDate(0, 0, -28)
			filter = bson.D{
				primitive.E{
					Key: "fecha",
					Value: bson.D{
						primitive.E{
							Key:   "$lt",
							Value: date,
						},
					}},
				primitive.E{
					Key:   "box",
					Value: 3,
				},
			}
		default:
			//filter = bson.D{}
			return quick()
		}
	}

	// categoty filter
	if cat != "" {
		filter = bson.D{
			primitive.E{
				Key:   "categoria",
				Value: cat},
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
		primitive.E{
			Key: "$inc",
			Value: bson.D{
				primitive.E{
					Key:   "box",
					Value: 1,
				},
			}},
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "fecha",
					Value: time.Now(),
				},
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
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{
					Key:   "box",
					Value: 0,
				},
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

func quick() ([][]string, error) {
	var (
		questions [][]string
	)
	//Open the db
	db, err := config.GetMongoDB()
	if err != nil {
		return nil, err
	}
	// Cursor Results
	c := db.Collection(COLLECTION)
	ctx := context.TODO()
	pipeline := []bson.D{
		primitive.D{
			primitive.E{
				Key: "$sample",
				Value: primitive.D{
					primitive.E{
						Key:   "size",
						Value: 10,
					}}}},
		primitive.D{
			primitive.E{
				Key: "$match",
				Value: primitive.D{
					primitive.E{
						Key:   "box",
						Value: 0,
					},
				},
			},
		},
	}
	cursor, err := c.Aggregate(ctx, pipeline)
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
