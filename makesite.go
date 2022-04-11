package main

import (
	"fmt"
	"os"
	// "log"
	// "html/template"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	content, err := os.ReadFile("first-post.txt")
	check(err)
	fmt.Println(string(content))
}
