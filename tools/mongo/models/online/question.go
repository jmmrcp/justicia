package online

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// Articulo schema in Db
	Articulo struct {
		Art string
		Ley string
	}
	// Mlab schema Mongo DB
	Mlab struct {
		// ID         primitive.ObjectID `bson:"_id" json:"id"`
		Test       int       `bson:"test" json:"test"`
		Categoria  string    `bson:"categoria" json:"categoria"`
		Temas      []int     `bson:"temas" json:"temas"`
		Titulo     string    `bson:"titulo" json:"titulo,"`
		Ord        int       `bson:"ord" json:"ord"`
		Pregunta   string    `bson:"pregunta" json:"pregunta"`
		Respuestas []string  `bson:"respuestas" json:"respuestas"`
		Articulos  string    `bson:"articulos" json:"articulos"`
		Fecha      time.Time `bson:"fecha" json:"fecha"`
		Contador   int       `bson:"contador" json:"contador"`
		Box        int       `bson:"box" json:"box"`
	}
)

// Questions enlaza a la tabla preguntas
func (db *DB) Questions() *mongo.Collection {
	return db.Collection("preguntas")
}

// GetAll conexion a la DB
func (db *DB) GetAll() ([]*Mlab, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{
		"tags": bson.M{
			"$elemMatch": bson.M{
				"$eq": "golang"},
		},
	}
	filter = bson.M{}
	defer cancel()
	questions := []*Mlab{}
	c := db.Questions()

	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var q Mlab
		if err := cursor.Decode(&q); err != nil {
			return nil, err
		}
		fmt.Printf("Question: %+v\n", q)
		questions = append(questions, &q)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

// GetTest devuelve test seleccionado
func (db *DB) GetTest() ([]*Mlab, error) {
	t := 4
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{
		"test": t,
	}
	options := options.Find()
	options.SetLimit(5)
	defer cancel()
	questions := []*Mlab{}
	c := db.Questions()

	cursor, err := c.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var q Mlab
		if err := cursor.Decode(&q); err != nil {
			return nil, err
		}
		// fmt.Printf("Question: %+v\n", q)
		questions = append(questions, &q)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

// InsertOne inserta en la base de datos
func (db *DB) InsertOne(data interface{}) error {
	ctx, cancel := context.WithTimeout(db.Context, 10*time.Second)
	defer cancel()
	// ctx := context.Background()
	c := db.Questions()
	res, err := c.InsertOne(ctx, data)
	if err != nil {
		return err
	}
	fmt.Printf("new question created with id: %s\n", res.InsertedID.(primitive.ObjectID).Hex())
	return nil
}

// InsertMany inserta varios al mismo tiempo
func (db *DB) InsertMany(data []interface{}) error {
	ctx, cancel := context.WithTimeout(db.Context, 10*time.Second)
	defer cancel()
	c := db.Questions()
	_, err := c.InsertMany(ctx, data)
	if err != nil {
		return err
	}
	// fmt.Printf("new questions created with ids: %s\n", res.InsertedIDs)
	return nil
}
