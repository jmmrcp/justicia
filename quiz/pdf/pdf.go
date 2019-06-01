package pdf

import (
	"fmt"
	"justicia/quiz"
	"justicia/quiz/questions"
	"math/rand"
	"strconv"

	"github.com/jung-kurt/gofpdf"
)

// Create a PDF Document
func Create(q questions.Questions, n int) error {
	if quiz.QuestionPdf {
		pdf := gofpdf.New("P", "mm", "A4", "")
		pdfs := gofpdf.New("P", "mm", "A4", "")
		pdf.SetTopMargin(15)
		pdfs.SetTopMargin(15)
		tr := pdf.UnicodeTranslatorFromDescriptor("")
		pdf.SetHeaderFuncMode(func() {
			pdf.SetY(5)
			pdf.SetFont("Arial", "B", 8)
			pdf.CellFormat(0, 4, "REP-##", "B", 0, "R", false, 0, "")
			pdf.Ln(5)
		}, true)
		pdfs.SetHeaderFuncMode(func() {
			pdfs.SetY(5)
			pdfs.SetFont("Arial", "B", 8)
			pdfs.CellFormat(0, 4, "RESPUESTAS", "B", 0, "C", false, 0, "")
			pdfs.Ln(5)
		}, true)

		pdf.SetFooterFunc(func() {
			pdf.SetY(-15)
			pdf.SetFont("Arial", "I", 8)
			pdf.CellFormat(0, 10, fmt.Sprintf("%d de {nb}", pdf.PageNo()), "TC", 0, "C", false, 0, "")
		})
		pdfs.SetFooterFunc(func() {
			pdfs.SetY(-15)
			pdfs.SetFont("Arial", "I", 8)
			pdfs.CellFormat(0, 10, fmt.Sprintf("%d de {nb}", pdfs.PageNo()), "TC", 0, "C", false, 0, "")
		})

		pdf.AliasNbPages("")
		pdf.AddPage()
		pdfs.AliasNbPages("")
		pdfs.AddPage()

		realquestions := len(q.Questions)

		if realquestions < n {
			n = realquestions
		}

		for i := 0; i < n; i++ {
			pdf.SetFont("Arial", "B", 10)
			pdfs.SetFont("Arial", "B", 10)
			pdf.SetX(10)
			pdf.MultiCell(0, 4, tr(fmt.Sprintf("%d)  %s", i+1, q.Questions[i].Question)), "", "", false)
			pdfs.SetX(10)
			pdfs.MultiCell(0, 4, tr(fmt.Sprintf("%d)  %s", i+1, q.Questions[i].Question)), "", "", false)
			for j := 0; j < 4; j++ {
				pdf.SetX(20)
				pdfs.SetX(20)
				pdf.SetFont("Arial", "", 10)
				pdfs.SetFont("Arial", "", 10)
				ok := q.Questions[i].Answers.Answers[j].Correct
				answer := q.Questions[i].Answers.Answers[j].Answer
				pdf.MultiCell(0, 4, tr(fmt.Sprintf("%c) ", j+97)+answer), "", "", false)
				if ok {
					pdfs.MultiCell(0, 4, tr(fmt.Sprintf("%c) ", j+97)+answer), "", "", false)
				}
			}
		}
		r := rand.New(rand.NewSource(99))
		n := strconv.Itoa(r.Intn(100))
		err := pdf.OutputFileAndClose("Test - " + n + ".pdf")
		if err != nil {
			return err
		}
		err = pdfs.OutputFileAndClose("Respuestas - " + n + ".pdf")
		if err != nil {
			return err
		}
	}
	return nil
}
