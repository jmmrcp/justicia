package dao

import (
	"context"
	"justicia/quiz/config"
	"justicia/quiz/models"
	"log"
	"time"

	"github.com/denisbrodbeck/machineid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	questions [][]string
	today     = time.Now()
	week      = today.AddDate(0, 0, -7)
	twoWeek   = today.AddDate(0, 0, -14)
	fourWeek  = today.AddDate(0, 0, -28)
	stages    = []time.Time{today, week, twoWeek, fourWeek}
	// Filter object
	Filter bson.D
	/*
		stageCero = bson.D{
			stage(0), box(0),
		}*/
	// StageOne Box #1
	StageOne = bson.D{
		stage(1), box(1),
	}
	// StageTwo Box #2
	StageTwo = bson.D{
		stage(2), box(2),
	}
	// StageThree Box #3
	StageThree = bson.D{
		stage(3), box(3),
	}
	// Correct Update a correct Question
	Correct = bson.D{
		primitive.E{
			Key: "$inc",
			Value: bson.D{
				box(1),
			}},
		primitive.E{
			Key: "$set",
			Value: bson.D{
				actual(),
			}},
		primitive.E{
			Key: "$push",
			Value: bson.D{
				id(),
			}}}
	// Wrong update a incorrect Question
	Wrong = bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				box(0),
			}},
		primitive.E{
			Key: "$push",
			Value: bson.D{
				id(),
			}}}

	result = bson.D{
		primitive.E{
			Key: "$sample",
			Value: primitive.D{
				primitive.E{
					Key:   "size",
					Value: 150,
				}}}}
)

// IDs return id bson object
func IDs(id string) bson.M {
	v, _ := primitive.ObjectIDFromHex(id)
	return bson.M{
		"_id": v,
	}
}

// box return Box number
func box(stage int) primitive.E {
	return primitive.E{
		Key:   "box",
		Value: stage,
	}
}

func stage(week int) primitive.E {
	return primitive.E{
		Key: "fecha",
		Value: bson.D{
			primitive.E{
				Key:   "$lt",
				Value: stages[week],
			}}}
}

func actual() primitive.E {
	return primitive.E{
		Key:   "fecha",
		Value: today,
	}
}

func id() primitive.E {
	id, err := machineid.ID()
	if err != nil {
		log.Fatal(err)
	}
	return primitive.E{
		Key:   id,
		Value: today,
	}
}

func pipe(box, stage primitive.E) bson.D {
	b, s := primitive.D{box}, primitive.D{stage}
	return bson.D{
		primitive.E{
			Key: "$match",
			Value: primitive.D{
				primitive.E{
					Key:   "$and",
					Value: []primitive.D{b, s},
				}}}}
}

// Test return test bson object
func Test(test int) bson.D {
	return bson.D{
		primitive.E{
			Key:   "test",
			Value: test,
		},
		box(0),
	}
}

// Category return category bson object
func Category(category string) bson.D {
	return bson.D{
		primitive.E{
			Key:   "categoria",
			Value: category,
		}}
}

// Tema return test bson object
func Tema(tema int) bson.D {
	return bson.D{
		primitive.E{
			Key:   "tema",
			Value: tema,
		},
		box(0),
	}
}

// Quick Filter results
func Quick() ([][]string, error) {

	//Open the db
	db, err := config.GetAtlasDB()
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(db.Context, 3*time.Second)
	defer cancel()

	// Cursor Results
	c := db.Collection(COLLECTION)

	p := pipe(box(0), stage(0))
	pipeline := []bson.D{p, result}

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
	if err = db.Client.Disconnect(ctx); err != nil {
		return nil, err
	}
	return questions, nil
}
