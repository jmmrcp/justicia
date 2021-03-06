// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package quiz

import (
	gui "justicia/quiz/interface"
	"justicia/quiz/questions"
	"log"

	"github.com/jroimartin/gocui"
)

//FBInit -- Initializes the Fill in the Blank gui interface
func FBInit(g *gocui.Gui, q *questions.Question, count string) (err error) {
	//The Answers
	as := q.Answers

	//Highlight the selected view and make it green
	g.Highlight = true
	g.SelFgColor = gocui.ColorGreen

	//Add content to gui
	questionFrame := gui.NewQuestionFrame("questionFrame")
	question := gui.NewQuestion("question", count, q.Question)
	answerBlank := gui.NewAnswer(gui.BoxBlank, gui.BoxBlank, as.Answers[0].Answer)
	infoBar := gui.NewInfoBar(gui.InfoBarName, gui.InfoBarFillInBlank)

	g.SetManager(questionFrame, question, answerBlank, infoBar)

	//Add keybindings
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, gui.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, FillInAnswer); err != nil {
		log.Panicln(err)
	}

	return nil
}
