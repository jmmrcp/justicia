package main

import (
	"flag"
	"fmt"
	"io"
	"justicia/quiz"
	"justicia/quiz/pdf"
	"justicia/quiz/questions"
	"log"
	"net/http"
	"os"

	"github.com/jroimartin/gocui"
)

var (
	// TEST Numero de test
	TEST int
	// VIEW Lista a utilizar
	VIEW int
	// COUNT Numero de preguntas
	COUNT int
	// CAT Categoria
	CAT string
	// PDFS Imprime o no
	PDFS bool
)

func init() {
	const (
		defaultCat   = ""
		usageCat     = "Category.\ndefault=none"
		defaultTest  = 0
		usageTest    = "Test Number.\ndefault=All."
		defaultCount = 100
		usageCount   = "Questions Number\ndefault=100"
		defaultView  = 0
		usageView    = "Question Marks\ndefault=1\n 1 - Daily\n 2 - Weekly \n 3 - Quincenally\n 4 - Monthly"
		defaultPdf   = false
		usagePdf     = "Generate a PDF with Test and Answers"
	)
	flag.StringVar(&CAT, "Category_Type", defaultCat, usageCat)
	flag.IntVar(&TEST, "Test_number", defaultTest, usageTest)
	flag.IntVar(&COUNT, "Question_number", defaultCount, usageCount)
	flag.IntVar(&VIEW, "Mark_number", defaultView, usageView)
	flag.BoolVar(&PDFS, "PDF", defaultPdf, usagePdf)
	flag.StringVar(&CAT, "C", defaultCat, usageCat+" (shorthand)")
	flag.IntVar(&TEST, "T", defaultTest, usageTest+" (shorthand)")
	flag.IntVar(&COUNT, "Q", defaultCount, usageCount+" (shorthand)")
	flag.IntVar(&VIEW, "M", defaultView, usageView+" (shorthand)")
	flag.BoolVar(&PDFS, "P", defaultPdf, usagePdf+" (shorthand)")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    justicia -C='LOPJ' -T=67 -Q=100 -M=1 -P=false...\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()

}

func main() {
	// Parse the flags.

	quiz.QuestionLimit = COUNT

	//Get gui driver
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Need to create questions
	if fileExists("data/data.db") {
		quiz.Questions, err = questions.CreateQuestionsDB(quiz.Questions, VIEW, TEST, CAT)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		quiz.Questions, err = questions.CreateQuestionsDAO(quiz.Questions, VIEW, TEST, CAT)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Shuffle Questions
	err = quiz.Questions.Shuffle()
	if err != nil {
		log.Fatal(err)
	}

	//Create PDF
	if PDFS {
		err = pdf.Create(quiz.Questions, COUNT)
		if err != nil {
			log.Fatal(err)
		}
	}

	//Need to initialize screen
	err = quiz.Init(g)
	if err != nil {
		log.Fatal(err)
	}

	//Run main loop
	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

func downloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
