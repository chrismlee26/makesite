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
	// newFile := flag.String("file", "", "Txt file to read:")
	// flag.Parse()
	// save(*newFile)

	directory := "."
	files, err := ioutil.ReadDir(directory)
	check(err)

	for _, file := range files {
		if file.Name()[len(file.Name())-3:] == "txt" {
			fmt.Println(file.Name(), "----")
			// return
		}
	}

	// dir := flag.String("dir", "", "Directory to read:")
	// flag.Parse()
}

func checkTxt(inputFile string) {

}

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

// func ReadDir(dirname string) ([]os.FileInfo, error)

// func ReadDir(dirname string) ([]os.FileInfo, error) {
// 	files, err := ioutil.ReadDir(".")
// 	check(err)

// 	for _, file := range files {
// 		fmt.Println("File:", file.Name())
// 	}
// 	fmt.Println(
// 		"\n", "Files in directory:", len(files),
// 	)
// 	return files, err
// }
