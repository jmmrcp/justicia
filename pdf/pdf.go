package pdf

import (
	"fmt"
	"quiz/questions"

	"github.com/jung-kurt/gofpdf"
)

// Create a PDF Document
func Create(q questions.Questions, n int) error {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetTopMargin(15)
	tr := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFuncMode(func() {
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 8)
		pdf.CellFormat(0, 4, "REP-##", "B", 0, "R", false, 0, "")
		pdf.Ln(5)
	}, true)

	pdf.SetFooterFunc(func() {
		pdf.SetY(-15)
		pdf.SetFont("Arial", "I", 8)
		pdf.CellFormat(0, 10, fmt.Sprintf("%d de {nb}", pdf.PageNo()), "TC", 0, "C", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()

	for i := 0; i < n; i++ {
		pdf.SetFont("Arial", "B", 10)
		pdf.MultiCell(0, 4, tr(fmt.Sprintf("%d)  %s", i+1, q.Questions[i].Question)), "", "", false)
		for j := 0; j < 4; j++ {
			pdf.SetX(20)
			pdf.SetFont("Arial", "", 10)
			answer := q.Questions[i].Answers.Answers[j].Answer
			pdf.MultiCell(0, 4, tr(fmt.Sprintf("%c) ", j+97)+answer), "", "", false)
		}
	}

	pdf.SetHeaderFuncMode(func() {
		pdf.SetY(5)
		pdf.SetFont("Arial", "B", 8)
		pdf.CellFormat(0, 4, "RESPUESTAS", "B", 0, "C", false, 0, "")
		pdf.Ln(5)
	}, true)

	pdf.AddPage()

	for i := 0; i < n; i++ {
		pdf.SetX(15)
		pdf.SetFont("Arial", "B", 10)
		pdf.MultiCell(0, 4, tr(fmt.Sprintf("%d)  %s", i+1, q.Questions[i].Question)), "", "", false)
		pdf.SetX(20)
		pdf.SetFont("Arial", "", 10)
		for j := 0; j < 4; j++ {
			ok := q.Questions[i].Answers.Answers[j].Correct
			answer := q.Questions[i].Answers.Answers[j].Answer
			if ok {
				pdf.MultiCell(0, 4, tr(fmt.Sprintf("%c) ", j+97)+answer), "", "", false)
			}
		}
	}

	err := pdf.OutputFileAndClose("Test.pdf")
	if err != nil {
		return err
	}
	return nil
}
