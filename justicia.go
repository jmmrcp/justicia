package main

import (
	"flag"
	"fmt"
	"io"
	"justicia/quiz"
	"log"
	"net/http"
	"os"

	"justicia/quiz/pdf"
	"justicia/quiz/questions"

	"github.com/jroimartin/gocui"
)

var (
	test  int
	view  int
	count int
	pdfs  bool
)

func init() {
	const (
		defaultTest  = 0
		usageTest    = "Test Number.\ndefault=All."
		defaultCount = 100
		usageCount   = "Questions Number\ndefault=100"
		defaultView  = 0
		usageView    = "Question Marks\ndefault=1\n 1 - Daily\n 2 - Weekly \n 3 - Quincenally\n 4 - Monthly"
		defaultPdf   = false
		usagePdf     = "Generate a PDF with Test and Answers"
	)
	flag.IntVar(&test, "Test_number", defaultTest, usageTest)
	flag.IntVar(&count, "Question_number", defaultCount, usageCount)
	flag.IntVar(&view, "Mark_number", defaultView, usageView)
	flag.BoolVar(&pdfs, "PDF", defaultPdf, usagePdf)
	flag.IntVar(&test, "T", defaultTest, usageTest+" (shorthand)")
	flag.IntVar(&count, "Q", defaultCount, usageCount+" (shorthand)")
	flag.IntVar(&view, "M", defaultView, usageView+" (shorthand)")
	flag.BoolVar(&pdfs, "P", defaultPdf, usagePdf+" (shorthand)")

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    test -T=67 -Q=100 -M=1 -P=false...\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
}

func main() {
	// Parse the flags.

	quiz.QuestionLimit = count
	quiz.QuestionMode = view
	quiz.QuestionTest = test
	quiz.QuestionPdf = pdfs

	fileUrl := "https://computerwizards.es/.well-known/data.db"

	if !fileExists("data/data.db") {
		if err := downloadFile("data/data.db", fileUrl); err != nil {
			panic(err)
		}
	}

	//Get gui driver
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Need to create questions
	quiz.Questions, err = questions.CreateQuestionsDB(quiz.Questions, quiz.QuestionMode, quiz.QuestionTest)
	if err != nil {
		log.Fatal(err)
	}

	//Shuffle Questions
	err = quiz.Questions.Shuffle()
	if err != nil {
		log.Fatal(err)
	}

	//Create PDF
	err = pdf.Create(quiz.Questions, count)
	if err != nil {
		log.Fatal(err)
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
