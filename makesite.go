package main

import (
	"flag"
	"fmt"
	"html/template"

	// "html/template"
	"io/ioutil"
	"os"
)

type Page struct {
	TextFilePath string
	TextFileName string
	HTMLPagePath string
	Content      string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func save(inputFile string, outputFile string) {
	slice := Page{
		TextFilePath: "/" + inputFile,
		TextFileName: inputFile,
		HTMLPagePath: outputFile,
		Content:      string(fileContents),
	}
	page, _ := template.ParseFiles("template.tmpl")

	newFile, _ := os.Create(outputFile)
	err := page.Execute(newFile, slice)
	check(err)

}

func main() {
	fileContents, err := ioutil.ReadFile("./" + inputFile + ".txt")
	check(err)
	outputFile := inputFile + ".html"

	// Define directory with variable name
	directory := "."
	files, err := ioutil.ReadDir(directory)
	check(err)
	flag.Parse()

	// **Flag to search specified directory for "txt" files

	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "txt" {
			fmt.Println(file.Name(), "\n", "----")
		}
	}

	dir := flag.String("dir", "", "Directory to read:")
	flag.Parse()
}

// **Take the outputs from the above function and create new HTML pages with the data.
// **Refactor save function below.
