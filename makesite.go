package main

import (
	"fmt"
	"io/ioutil"
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

func main() {
	// Define directory with variable name
	directory := "."
	files, err := ioutil.ReadDir(directory)
	check(err)

	// **Flag to search specified directory for "txt" files

	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "txt" {
			fmt.Println(file.Name(), "\n", "----")
		}
	}

	// dir := flag.String("dir", "", "Directory to read:")
	// flag.Parse()
}

// **Take the outputs from the above function and create new HTML pages with the data.
// **Refactor save function below.

// func save(inputFile string) {
// 	fileContents, err := ioutil.ReadFile("./" + inputFile + ".txt")
// 	check(err)
// 	outputFile := inputFile + ".html"

// 	page := Page{
// 		TextFilePath: "/" + inputFile,
// 		TextFileName: inputFile,
// 		HTMLPagePath: outputFile,
// 		Content:      string(fileContents),
// 	}

// 	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

// 	newFile, err := os.Create(outputFile)
// 	check(err)

// 	t.Execute(newFile, page)

// }
