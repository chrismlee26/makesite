package main

import (
	"fmt"
	"log"
	"os"
	// "os"
	// "html/template"
)

func main() {
	content, err := os.ReadFile("first-post.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(content))
}
