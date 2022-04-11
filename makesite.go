package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// content, err := os.ReadFile("first-post.txt")
	// check(err)
	// fmt.Println(string(content))

	bytesToWrite := []byte("hello\ngo")
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	check(err)
}
