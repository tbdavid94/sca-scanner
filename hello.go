package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"os"
	"strings"
)

func GenerateIndexFile() {
	resultReport, err := os.ReadFile("reports/index.html") // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	fmt.Println(resultReport) // print the content as 'bytes'
	resultReportStr := string(resultReport) // convert content to a 'string'
	fmt.Println(resultReportStr) // print the content as a 'string'

	strReader := strings.NewReader(resultReportStr)

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(strReader)
	if err != nil {
		log.Fatal(err)
	}

	// Find element need remove
	doc.Find("#test").Remove()
	//elementString := string(element)
	//fmt.Print(string(element))
}

func main() {
	GenerateIndexFile()
}