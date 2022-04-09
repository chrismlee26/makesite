package main

import (
	"fmt"
	"html/template"
)

func main() {
	fmt.Println("Hello, world!")
	// Parse data
	paths := []string{"template.tmpl"}

	t := template.Must(template.New("html-tmpl").ParseFiles(paths...))
}
