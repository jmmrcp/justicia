// Copyright (C) 2019 José Martínez Ruiz <jmmrcp@gmail.com>.
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file.

package questions

import (
	"justicia/quiz/answers"
	"log"
)

func OrderFour(as answers.Answers, qData []string) {
	as.Answers = append(as.Answers, &answers.Answer{Answer: qData[1], Correct: true})
	as.Answers = append(as.Answers, &answers.Answer{Answer: qData[2], Correct: false})
	as.Answers = append(as.Answers, &answers.Answer{Answer: qData[3], Correct: false})
	as.Answers = append(as.Answers, &answers.Answer{Answer: qData[4], Correct: false})
}

func OrderTwo(as answers.Answers, qData []string) {
	as.Answers = append(as.Answers, &answers.Answer{Answer: qData[1], Correct: true})
	as.Answers = append(as.Answers, &answers.Answer{Answer: qData[2], Correct: false})
}

func OrdenOne(as answers.Answers, qData []string) {
	as.Answers = append(as.Answers, &answers.Answer{Answer: qData[1], Correct: true})
}

func TextQuestions() {

}

func DBQuestions(qs Questions, data [][]string) {
	for _, qData := range data {
		l := len(qData)
		as := answers.Answers{Answers: []*answers.Answer{}}

		if l >= 6 {
			as.Answers = append(as.Answers, &answers.Answer{Answer: qData[1], Correct: true})
			as.Answers = append(as.Answers, &answers.Answer{Answer: qData[2], Correct: false})
			as.Answers = append(as.Answers, &answers.Answer{Answer: qData[3], Correct: false})
			as.Answers = append(as.Answers, &answers.Answer{Answer: qData[4], Correct: false})
		} else if l == 4 {
			as.Answers = append(as.Answers, &answers.Answer{Answer: qData[1], Correct: true})
			as.Answers = append(as.Answers, &answers.Answer{Answer: qData[2], Correct: false})
		} else if l == 3 {
			as.Answers = append(as.Answers, &answers.Answer{Answer: qData[1], Correct: true})
		}

		//Shuffle the answers
		err := as.Shuffle()
		if err != nil {
			log.Fatal(err)
		}

		//ID, _ := strconv.Atoi(q[l-1])
		qs.Questions = append(qs.Questions, &Question{qData[0], as, qData[l-2], qData[l-1]})
	}
	return qs, nil
}
