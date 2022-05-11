package main

import (
	"flag"
	"fmt"
	"html/template"
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

func save(inputFile string) {
	fileContents, err := ioutil.ReadFile(inputFile + ".txt")
	slice := Page{
		TextFilePath: "/" + inputFile,
		TextFileName: inputFile,
		HTMLPagePath: inputFile + ".html",
		Content:      string(fileContents),
	}
	page, _ := template.ParseFiles("template.tmpl")

	newFile, _ := os.Create(inputFile + ".html")
	err = page.Execute(newFile, slice)
	check(err)
}

func main() {

	// Render specified text file to HTML page
	// ex command: go run makesite.go --file first-post
	// NOTE make sure not to include extension of text file (first-post vs first-post.txt)
	inputFilePtr := flag.String("file", "", "File to render to HTML")
	inputFile := *inputFilePtr
	if inputFile != "" {
		save(inputFile)
	}

	// Render text files in specified directory to HTML pages
	// ex command: go run makesite.go --dir .
	dirPtr := flag.String("dir", "", "Directory to read")
	flag.Parse()
	dir := *dirPtr

	if dir != "" {
		files, _ := ioutil.ReadDir(dir)

		for _, file := range files {
			if file.Name()[len(file.Name())-3:] == "txt" {
				fmt.Println(file.Name(), "\n", "----")
				save(file.Name()[:len(file.Name())-4])
			}
		}
	}
}
