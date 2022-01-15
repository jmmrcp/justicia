// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package quiz

import (
	"justicia/quiz/answers"
	gui "justicia/quiz/interface"
	"justicia/quiz/user"
	"strconv"

	"github.com/jroimartin/gocui"
)

//AnswersToBoxViews -- used to map the user selected box view to the actual answer
var AnswersToBoxViews = map[string]*answers.Answer{}

//FillInAnswer -- Callback used for the fill in the blank answers
func FillInAnswer(g *gocui.Gui, v *gocui.View) error {
	cQuestion, err := Questions.Current()
	if err != nil {
		return err
	}

	filledInAnswer := &answers.Answer{Answer: v.Buffer(), Correct: true}

	a := user.Answer{
		Question: cQuestion,
		Answer:   filledInAnswer,
	}

	//User answers -- The plus one is so the count starts at 1
	count := len(UserAnswers)+1
	UserAnswers[strconv.Itoa(count)] = &a

	//Increment the questions index!
	Questions.Index++

	//Next Screen
	err = Init(g)
	if err != nil {
		return err
	}
	return nil
}

//SelectAnswer -- Callback used to select an answer in quiz layouts that have
//multiple answers to select from
func SelectAnswer(g *gocui.Gui, v *gocui.View) error {
	//Reset variable used for tabbing through solutions
	gui.ActiveView = 0

	cQuestion, err := Questions.Current()
	if err != nil {
		return err
	}

	selectedAnswer := AnswersToBoxViews[v.Name()]

	a := user.Answer{
		Question: cQuestion,
		Answer:   selectedAnswer,
	}

	//User answers -- The plus one is so the count starts at 1
	long := len(UserAnswers)
	convert := strconv.Itoa(long)
	convert += 1
	UserAnswers[convert] = &a


	//Increment the questions index!
	Questions.Index++

	//Next Screen
	err = Init(g)
	if err != nil {
		return err
	}
	return nil
}

//NextUserAnswer -- View next user answer
func NextUserAnswer(g *gocui.Gui, v *gocui.View) error {
	//Increment count
	CurrentUserAnswer++

	//Next Screen
	err := Init(g)
	if err != nil {
		return err
	}
	return nil
}
