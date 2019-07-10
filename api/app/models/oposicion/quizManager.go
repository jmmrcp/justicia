package oposicion

import (
	"justicia/api/app/models/helpers"

	"go.uber.org/zap"
	"gopkg.in/mgo.v2/bson"
)

type QuizManager struct {
	Repo   *QuizRepository
	Logger *zap.Logger
}

func (m *QuizManager) GetAll() ([]*Quiz, error) {
	quizzes, err := m.Repo.FindAll()

	if quizzes == nil {
		quizzes = []*Quiz{}
	}

	if err != nil {
		m.Logger.Error(err.Error())
	}

	return quizzes, err
}

func (m *QuizManager) Get(id string) (*Quiz, error) {
	quiz, err := m.Repo.FindByID(id)

	if m.Repo.IsNotFoundErr(err) {
		return nil, helpers.NewErrNotFound("Quiz `" + id + "` does not exist.")
	}

	if err != nil {
		m.Logger.Error(err.Error())
	}

	return quiz, err
}
func (m *QuizManager) Update(id string, quiz *Quiz) (*Quiz, error) {

	quiz.ID = bson.ObjectId(id)

	err := m.Repo.Update(quiz)

	if m.Repo.IsNotFoundErr(err) {
		return nil, helpers.NewErrNotFound("Quiz `" + id + "` does not exist.")
	}

	if err != nil {
		m.Logger.Error(err.Error())
		return nil, err
	}

	return quiz, err
}
