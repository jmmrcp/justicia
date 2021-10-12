package main

import (
	"flag"
	"fmt"
	"justicia/quiz"
	"justicia/quiz/pdf"
	"justicia/quiz/questions"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/jroimartin/gocui"
)

var (
	// TEMA numero de tema
	TEMA int
	// TEST Numero de test
	TEST int
	// VIEW Lista a utilizar
	VIEW int
	// COUNT Numero de preguntas
	COUNT int
	// CAT Categoria
	CAT string
	// PDFS Imprime o no
	PDFS    bool
	version string
	date    string
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}

	const (
		defaultTema  = 0
		usageTema    = "Tema.\ndefault = none."
		defaultCat   = ""
		usageCat     = "Category.\ndefault = none."
		defaultTest  = 0
		usageTest    = "Test Number.\ndefault = All."
		defaultCount = 10
		usageCount   = "Questions Number.\ndefault = 10."
		defaultView  = 0
		usageView    = "Question Marks.\ndefault = 0.\n  0 - Daily.\n  1 - Weekly.\n  2 - Quincenally.\n  3 - Monthly."
		defaultPdf   = false
		usagePdf     = "Generate a PDF with Test and Answers."
	)
	if version == "" {
		version = "no version"
	}
	if date == "" {
		date = "(Mon YYYY)"
	}
	// flag.StringVar(&CAT, "Category_Type", defaultCat, usageCat)
	// flag.IntVar(&TEST, "Test_number", defaultTest, usageTest)
	// flag.IntVar(&COUNT, "Question_number", defaultCount, usageCount)
	// flag.IntVar(&VIEW, "Mark_number", defaultView, usageView)
	// flag.BoolVar(&PDFS, "PDF", defaultPdf, usagePdf)
	flag.StringVar(&CAT, "C", defaultCat, usageCat)
	flag.IntVar(&TEST, "T", defaultTest, usageTest)
	flag.IntVar(&COUNT, "Q", defaultCount, usageCount)
	flag.IntVar(&VIEW, "V", defaultView, usageView)
	flag.BoolVar(&PDFS, "P", defaultPdf, usagePdf)
	flag.IntVar(&TEMA, "R", defaultTema, usageTema)

	flag.Usage = func() {
		fmt.Printf("\njusticia version %s %s.\n", version, date)
		fmt.Printf("  Usage:\n")
		fmt.Printf("    justicia -R=0 -C=CE -T=4 -Q=10 -M=0 -P=false ... (by default)\n\n")
		flag.PrintDefaults()
		os.Exit(0)
	}
	flag.Parse()
}

func main() {
	// Parse the flags.
	CAT = strings.ToUpper(CAT)
	quiz.QuestionLimit = COUNT

	//Get gui driver
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	//Need to create questions
	/* if quiz.FileExists("tools/mongo/data.db") {
		quiz.Questions, err = questions.CreateQuestionsDB(quiz.Questions, VIEW, TEST, TEMA, CAT)
		if err != nil {
			log.Fatal(err)
		}
	} else { */
	quiz.Questions, err = questions.CreateQuestionsDAO(quiz.Questions, VIEW, TEST, TEMA, CAT)
	if err != nil {
		log.Fatal(err)
	}
	// }

	//Shuffle Questions
	err = quiz.Questions.Shuffle()
	if err != nil {
		log.Fatal(err)
	}

	//Create PDF
	if PDFS {
		err = pdf.Create(quiz.Questions, COUNT, TEST)
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

/*
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
*/
