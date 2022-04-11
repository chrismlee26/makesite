package main

import (
	"flag"
	"html/template"
	"io/ioutil"
	"os"
	// "fmt"
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
	// fileContents, err := ioutil.ReadFile("first-post.txt")
	// check(err)

	// page := Page{
	// 	TextFilePath: "./first-post.txt",
	// 	TextFileName: "First Post",
	// 	HTMLPagePath: "first-post.html",
	// 	Content:      string(fileContents),
	// }

	// t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	// newFile, err := os.Create("first-post.html")
	// check(err)

	// t.Execute(newFile, page)

	// fmt.Println(string(fileContents))

	// newFile.Close()

	newFile := flag.String("file", "default", "Txt file to read:")
	flag.Parse()
	save(*newFile)

	// http.ListenAndServe(":8080", nil)
}

func save(inputFile string) {
	fileContents, err := ioutil.ReadFile("./" + inputFile + ".txt")
	check(err)
	outputFile := inputFile + ".html"

	page := Page{
		TextFilePath: "./" + inputFile,
		TextFileName: inputFile,
		HTMLPagePath: outputFile,
		Content:      string(fileContents),
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, err := os.Create("first-post.html")
	check(err)

	t.Execute(newFile, page)

}
