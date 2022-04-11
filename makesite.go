package main

import (
	// "fmt"

	"html/template"
	"os"
)

type post struct {
	Title   string
	Content string
}

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
	// content, err := os.ReadFile("first-post.txt")
	// check(err)
	// fmt.Println(string(content))

	// bytesToWrite := []byte("hello\ngo")
	// err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	// check(err)

	fileContents := []byte("A string for demo purposes.")
	page := Page{
		TextFilePath: "filePath",
		TextFileName: "new",
		HTMLPagePath: "new.html",
		Content:      string(fileContents),
	}

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))

	newFile, err := os.Create("new.html")
	check(err)

	t.Execute(newFile, page)
}
