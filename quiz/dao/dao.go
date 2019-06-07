package dao

import (
	"context"
	"justicia/quiz"
	"justicia/quiz/answers"
	"justicia/quiz/models"
	"justicia/quiz/questions"
	"log"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type (
	// QuestionsDAO data access object to DB
	QuestionsDAO struct {
		Server   string
		Database string
	}
)

var db *mongo.Database

const (
	// COLLECTION Document TABLE
	COLLECTION = "preguntas"
)

// Connect to MLAB DB
func (q *QuestionsDAO) Connect() {
	ctx := context.Background()
	client, err := mongo.NewClient(
		options.Client().ApplyURI(q.Server),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := client.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	db = client.Database(q.Database)
}

// GetAll All document Find
func (q *QuestionsDAO) GetAll() ([]models.Mlab, error) {
	var questions []models.Mlab
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{}
	defer cancel()
	c := db.Collection(COLLECTION)
	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var q models.Mlab
		if err := cursor.Decode(&q); err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

// GetTest All document Find
func (q *QuestionsDAO) GetTest() ([]models.Mlab, error) {
	var questions []models.Mlab
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{
		"test": quiz.QuestionTest,
	}
	defer cancel()
	c := db.Collection(COLLECTION)
	cursor, err := c.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var q models.Mlab
		if err := cursor.Decode(&q); err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

// GetCategory All document Find
func (q *QuestionsDAO) GetCategory(f string) ([]models.Mlab, error) {
	var questions []models.Mlab
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	filter := bson.M{
		"categoria": f,
	}
	options := options.Find()
	options.SetLimit(int64(quiz.QuestionLimit))
	defer cancel()
	c := db.Collection(COLLECTION)
	cursor, err := c.Find(ctx, filter, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var q models.Mlab
		if err := cursor.Decode(&q); err != nil {
			return nil, err
		}
		questions = append(questions, q)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}
	return questions, nil
}

// GetByID document ID Find
func (q *QuestionsDAO) GetByID(id string) (models.Mlab, error) {
	var question models.Mlab
	filter := bson.M{
		"ID": id,
	}
	c := db.Collection(COLLECTION)
	err := c.FindOne(context.TODO(), filter).Decode(&question)
	if err != nil {
		return question, err
	}
	return question, nil
}

// Create document on DB
func (q *QuestionsDAO) Create(question models.Mlab) error {
	return nil
}

// Update document on DB
func (q *QuestionsDAO) Update(id string, question models.Mlab) error {
	filter := bson.M{
		"ID": id,
	}
	update := bson.D{
		{"$inc", bson.D{
			{"box", 1},
		}},
	}
	c := db.Collection(COLLECTION)
	_, err := c.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

// CreateQuestionsDAO from the Mlabs DB
func (q *QuestionsDAO) CreateQuestionsDAO(qs questions.Questions) (questions.Questions, error) {
	var (
		//data [][]string
		err error
	)

	//Get data from mlabs
	data, err := q.GetAll()
	if err != nil {
		return qs, err
	}

	for _, qData := range data {
		l := len(qData)
		as := answers.Answers{[]*answers.Answer{}}

		if l == 7 {
			as.Answers = append(as.Answers, &answers.Answer{qData[1], true})
			as.Answers = append(as.Answers, &answers.Answer{qData[2], false})
			as.Answers = append(as.Answers, &answers.Answer{qData[3], false})
			as.Answers = append(as.Answers, &answers.Answer{qData[4], false})
		} else if l == 4 {
			as.Answers = append(as.Answers, &answers.Answer{qData[1], true})
			as.Answers = append(as.Answers, &answers.Answer{qData[2], false})
		} else if l == 3 {
			as.Answers = append(as.Answers, &answers.Answer{qData[1], true})
		}

		//Shuffle the answers
		as.Shuffle()

		ID, _ := strconv.Atoi(qData[l-1])

		qs.Questions = append(qs.Questions, &questions.Question{qData[0], as, qData[l-2], ID})
	}
	// log.Printf("%v\n", data)
	return qs, nil
}
