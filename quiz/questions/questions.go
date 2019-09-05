package questions

import (
	"fmt"
	"justicia/quiz/answers"
	"justicia/quiz/csv"
	"justicia/quiz/dao"
	"justicia/quiz/db"
	"math/rand"
	"time"
)

func init() {
	//Seeding the random number generator
	rand.Seed(time.Now().UnixNano())
}

//Question -- the interface for a question
type Question struct {
	Question    string
	Answers     answers.Answers
	Explanation string
	ID          string
}

//CorrectAnswer -- returns the currect answer
func (q Question) CorrectAnswer() (*answers.Answer, error) {
	result, err := q.Answers.CorrectAnswer()
	if err != nil {
		return result, fmt.Errorf("No correct answer was found")
	}
	return result, nil
}

//Questions -- The container that holds the questions
type Questions struct {
	Questions []*Question
	Index     int
}

//Shuffle -- In place Shuffle of the questions
func (qs Questions) Shuffle() error {
	numOfQuestions := len(qs.Questions)

	if len(qs.Questions) < 2 {
		return fmt.Errorf("There are not enough questions to be shuffled")
	}

	for i := range qs.Questions {
		swapIndex := rand.Intn(numOfQuestions)
		tempt := qs.Questions[i]
		qs.Questions[i] = qs.Questions[swapIndex]
		qs.Questions[swapIndex] = tempt
	}
	return nil
}

//Current -- returns the current question
func (qs Questions) Current() (*Question, error) {
	if len(qs.Questions) == 0 {
		return &Question{"", answers.Answers{Answers: []*answers.Answer{}}, "", ""}, fmt.Errorf("There are no questions")
	}
	return qs.Questions[qs.Index], nil
}

//NextExist -- Checks to see if there is a next question
func (qs Questions) NextExist() bool {
	return qs.Index < len(qs.Questions)-1
}

//Next -- Moves index pointer to the next question
func (qs Questions) Next() (*Question, error) {
	var q *Question

	if !qs.NextExist() {
		return q, fmt.Errorf("There is no next question")
	}

	qs.Index++

	q, err := qs.Current()
	if err != nil {
		return q, err
	}
	return q, nil
}

//PreviousExist -- Check to see if there is a previous question
func (qs Questions) PreviousExist() bool {
	return qs.Index > 0
}

//Previous -- movies index pointer to the previous question
func (qs Questions) Previous() (*Question, error) {
	var q *Question

	if !qs.PreviousExist() {
		return q, fmt.Errorf("There is no previous question")
	}

	qs.Index--

	q, err := qs.Current()
	if err != nil {
		return q, err
	}
	return q, nil
}

//NewQuestions -- returns an empty Question container
func NewQuestions() Questions {
	return Questions{[]*Question{}, 0}
}

//CreateQuestionsCSV -- used to create questions
func CreateQuestionsCSV(qs Questions, files []string) (Questions, error) {
	var data [][]string
	var err error

	//Get data From files
	for _, file := range files {
		data, err = csv.Read(file, data)
		if err != nil {
			return qs, err
		}
	}

	for _, qData := range data {
		l := len(qData)
		as := answers.Answers{Answers: []*answers.Answer{}}

		if l == 6 {
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
		as.Shuffle()

		qs.Questions = append(qs.Questions, &Question{qData[0], as, qData[l-1], ""})
	}
	return qs, nil
}

//CreateQuestionsDB -- used to create questions
func CreateQuestionsDB(qs Questions, view, test, tema int, cat string) (Questions, error) {
	var (
		data [][]string
		err  error
	)

	//Get data from db or dbs
	data, err = db.Read("tools/mongo/data.db", data, view, test, tema, cat)
	if err != nil {
		return qs, err
	}

	for _, qData := range data {
		l := len(qData)
		as := answers.Answers{Answers: []*answers.Answer{}}

		if l == 7 {
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
		as.Shuffle()

		//ID, _ := strconv.Atoi(qData[l-1])
		qs.Questions = append(qs.Questions, &Question{qData[0], as, qData[l-2], qData[l-1]})
	}
	// log.Printf("%v\n", data)
	return qs, nil
}

//CreateQuestionsDAO -- used to create questions
func CreateQuestionsDAO(qs Questions, view, test, tema int, cat string) (Questions, error) {
	var (
		data [][]string
		err  error
	)

	//Get data from mlabs
	data, err = dao.Read(data, view, test, tema, cat)
	if err != nil {
		return qs, err
	}

	for _, qData := range data {
		l := len(qData)
		as := answers.Answers{Answers: []*answers.Answer{}}

		if l == 7 {
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
		as.Shuffle()

		//ID, _ := strconv.Atoi(q[l-1])
		qs.Questions = append(qs.Questions, &Question{qData[0], as, qData[l-2], qData[l-1]})
	}
	return qs, nil
}
