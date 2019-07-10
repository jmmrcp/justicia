package oposicion

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type QuizRepository struct {
	Session *mgo.Session
}

func (repo *QuizRepository) collection() *mgo.Collection {
	return repo.Session.DB("justicia").C("preguntas")
}

func (repo *QuizRepository) FindAll() ([]*Quiz, error) {
	var quizzes []*Quiz
	err := repo.collection().Find(nil).All(&quizzes)
	return quizzes, err
}
func (repo *QuizRepository) FindByID(id string) (*Quiz, error) {
	var quiz *Quiz
	i := bson.ObjectIdHex(id)
	err := repo.collection().FindId(i).One(&quiz)
	return quiz, err
}
func (repo *QuizRepository) Update(quiz *Quiz) error {
	return repo.collection().UpdateId(quiz.ID, quiz)
}
func (repo *QuizRepository) IsNotFoundErr(err error) bool {
	return err == mgo.ErrNotFound
}
func (repo *QuizRepository) IsAlreadyExistErr(err error) bool {
	return mgo.IsDup(err)
}
