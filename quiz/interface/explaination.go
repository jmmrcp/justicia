// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package gui

import (
	"fmt"
	texttowidth "justicia/quiz/text-to-width"

	"github.com/jroimartin/gocui"
)

//Explanation -- Gui component that holds the question
type Explanation struct {
	name     string
	result   string
	question string
	answer   string
	explain  string
}

//NewExplanation -- creates new question gui component
func NewExplanation(name, result, question, answer, explain string) *Explanation {
	return &Explanation{name: name, result: result, question: question, answer: answer, explain: explain}
}

//location -- holds the location logic
func (e *Explanation) location(g *gocui.Gui) (x, y, w, h int) {
	maxX, maxY := g.Size()
	x = int(0.1 * float32(maxX))
	y = int(0.3 * float32(maxY))
	w = int(0.9 * float32(maxX))
	h = int(0.8 * float32(maxY))
	return
}

//Layout -- Tells gocui.Gui how to display this compenent
func (e *Explanation) Layout(g *gocui.Gui) error {
	x, y, w, h := e.location(g)
	v, err := g.SetView(e.name, x, y, w, h)
	if err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		//Set to false because I am using text-to-width for word wraping
		v.Wrap = false

		//result as title -- you were right or wrong
		v.Title = e.result

		//The allowed length that the text can take
		length := w - x - 1 //The 1 seems to be needed to keep the text within the bounds of the box

		//Display question
		fmt.Fprint(v, texttowidth.Format(fmt.Sprintf("(Pregunta) -- %s\n\n", e.question), length))

		//Display answer
		fmt.Fprint(v, texttowidth.Format(fmt.Sprintf("(Respuesta) -- %s\n\n", e.answer), length))

		//Display explanation
		fmt.Fprint(v, texttowidth.Format(fmt.Sprintf("(Explicacion) -- %s", e.explain), length))
	}
	return nil
}
