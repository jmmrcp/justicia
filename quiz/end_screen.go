// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package quiz

import (
	"fmt"
	"justicia/quiz/dao"
	"justicia/quiz/db"
	"justicia/quiz/user"
	"log"
	"os"

	gui "justicia/quiz/interface"

	"github.com/jroimartin/gocui"
)

// C is an array for Correct IDs
var C []string

// I is an array for Incorrect IDs
var I []string

//ESInit -- End Screen Initialization. Presents the results
func ESInit(g *gocui.Gui, u user.Answers) (err error) {
	//End quiz when you run out of answers
	if CurrentUserAnswer > len(UserAnswers) {
		g.Close()
		if FileExists("tools/mongo/data.db") {
			boxDb()
		} else {
			boxDao()
		}
		os.Exit(0)
	}

	//Place holder for 'correct' string that gets placed
	//in the gui
	var correct string

	numCorrect, err := u.TotalCorrect()
	if err != nil {
		log.Panicln(err)
	}
	//Score = num of correct answers over total
	score := fmt.Sprintf("%d/%d", numCorrect, u.Total())

	//Currently need a global counter...
	questionCount := fmt.Sprintf("%d", CurrentUserAnswer)

	currentUserAnswer := u[questionCount]
	answerCorrect, err := currentUserAnswer.Correct()
	if err != nil {
		log.Panicln(err)
	}
	if answerCorrect {
		C = append(C, currentUserAnswer.Question.ID)
		correct = gui.Right
	} else {
		I = append(I, currentUserAnswer.Question.ID)
		correct = gui.Wrong
	}

	//Get the correct answer
	correctAnswer, err := currentUserAnswer.Question.CorrectAnswer()
	if err != nil {
		log.Panicln(err)
	}

	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//create widgets
	scoreWidget := gui.NewScore(gui.ScoreName, score, questionCount)
	explainWidget := gui.NewExplanation(gui.Explain, correct, currentUserAnswer.Question.Question, correctAnswer.Answer, currentUserAnswer.Question.Explanation)
	infoBar := gui.NewInfoBar(gui.InfoBarName, gui.InfoBarEndScreen)

	g.SetManager(scoreWidget, explainWidget, infoBar)

	//Keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, NextUserAnswer); err != nil {
		log.Panicln(err)
	}
	return nil
}

func boxDao() {
	if len(C) > 0 {
		err := dao.Update(C)
		if err != nil {
			log.Fatal(err)
		}
	}
	if len(I) > 0 {
		err := dao.Unupdate(I)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func boxDb() {
	for _, v := range C {
		err := db.Update(v)
		if err != nil {
			log.Fatal(err)
		}
	}
	for _, v := range I {
		err := db.Unupdate(v)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// FileExists check filename
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return os.IsNotExist(err)
}
