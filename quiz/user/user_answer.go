// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package user

import (
	"justicia/quiz/answers"
	"justicia/quiz/questions"
	"strings"
)

//Answer -- Container used to hold a user's answer
type Answer struct {
	Question *questions.Question
	Answer   *answers.Answer
}

//Correct -- Checks to see if the user's answer was correct
func (u Answer) Correct() (bool, error) {
	correctAnswer, err := u.Question.CorrectAnswer()
	if err != nil {
		return false, err
	}

	if correctAnswer == u.Answer {
		return true, nil
	}

	//used to compare user input with answer
	lowerCA := strings.ToLower(strings.TrimSpace(correctAnswer.Answer))
	userA := strings.ToLower(strings.TrimSpace(u.Answer.Answer))
	if strings.EqualFold(lowerCA, userA) {
		return true, nil
	}
	return false, nil
}
