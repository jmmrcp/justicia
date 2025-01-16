// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package quiz

import (
	"log"

	"github.com/jroimartin/gocui"
)

//Init -- The Init function decides which sub Init function should be called
//ABCDInint, TFInit or FBInit
func Init(g *gocui.Gui) (err error) {
	//Have we reached the question limit?
	if Questions.Index >= QuestionLimit {
		//Need to call End Screen
		err := ESInit(g, UserAnswers)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	//Have we ran out of questions?
	if Questions.Index >= len(Questions.Questions) {
		//Need to call End Screen
		err := ESInit(g, UserAnswers)
		if err != nil {
			log.Fatal(err)
		}
		return nil
	}

	//Get current question
	q, err := Questions.Current()
	if err != nil {
		return err
	}

	//Use number of answers to figure out which Init function to use
	numOfAnswers := len(q.Answers.Answers)

	var long uint8 = len(UserAnswer)
	long++
	// var convert string = strconv.Itoa(long)
	var count string = fmt.Sprintf("Pregunta  + %d", long)

	// count := "Pregunta " + convert // + " (" + q.ID + ")"

	if numOfAnswers == 4 {
		err = ABCDInit(g, q, count)
		if err != nil {
			return err
		}
	} else if numOfAnswers == 2 {
		err = TFInit(g, q, count)
		if err != nil {
			return err
		}
	} else if numOfAnswers == 1 {
		err = FBInit(g, q, count)
		if err != nil {
			return err
		}
	}
	return nil
}
