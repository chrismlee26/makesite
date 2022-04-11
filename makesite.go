package main

import (
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

func main() {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	check(err)

	page := Page{
		TextFilePath: "./first-post.txt",
		TextFileName: "First Post",
		HTMLPagePath: "first-post.html",
		Content:      string(fileContents),
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, err := os.Create("first-post.html")
	check(err)

	t.Execute(newFile, page)
	newFile.Close()

	// http.ListenAndServe(":8080", nil)
}
