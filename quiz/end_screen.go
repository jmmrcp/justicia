package quiz

import (
	"fmt"
	"justicia/quiz/dao"
	"justicia/quiz/user"
	"log"
	"os"

	gui "justicia/quiz/interface"

	"github.com/jroimartin/gocui"
)

// C Correct IDs
var C []string

// I Incorrect Ids
var I []string

//ESInit -- End Screen Initialization. Presents the results
func ESInit(g *gocui.Gui, u user.Answers) (err error) {
	//End quiz when you run out of answers
	if CurrentUserAnswer > len(UserAnswers) {
		g.Close()
		box()
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

func box() {
	for _, v := range C {
		dao.Update(v)
	}
	for _, v := range I {
		dao.Unupdate(v)
	}
}
