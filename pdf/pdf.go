package pdf

import (
	"fmt"
	"log"

	"github.com/jung-kurt/gofpdf"
)

// Create a PDF Document
func Create() error {
	data := []string{
		"Test Autogenerado",
		"REP",
	}
	pdf := newReport()
	pdf = header(pdf, data)
	pdf.Ln(12)
	pdf = footer(pdf)
	if pdf.Err() {
		log.Fatalf("Failed creating PDF report: %s\n", pdf.Error())
	}
	err := savePDF(pdf)
	if err != nil {
		return err
	}
	return nil
}

func newReport() *gofpdf.Fpdf {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	return pdf
}

func header(pdf *gofpdf.Fpdf, hdr []string) *gofpdf.Fpdf {
	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(240, 240, 240)
	for _, str := range hdr {
		pdf.CellFormat(40, 7, str, "B", 0, "R", false, 0, "")
	}
	pdf.Ln(-1)
	return pdf
}

func footer(pdf *gofpdf.Fpdf) *gofpdf.Fpdf {
	pdf.SetFont("Arial", "B", 8)
	pdf.CellFormat(0, 10, fmt.Sprintf("Page %d", pdf.PageNo()), "TC", 0, "C", false, 0, "")
	return pdf
}

func savePDF(pdf *gofpdf.Fpdf) error {
	return pdf.OutputFileAndClose("report.pdf")
}
