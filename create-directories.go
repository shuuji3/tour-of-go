// Create template directories and copy main.go into each directory.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		PrintUsage()
		os.Exit(1)
	}

	// Parse params
	sectionName := os.Args[1]
	numberOfDirectories, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic("invalid param: number-of-directories")
	}

	// Create directories & main.go
	for i := 0; i < numberOfDirectories; i++ {
		directoryName := fmt.Sprintf("%s/%d/", sectionName, i+1)
		err := os.MkdirAll(directoryName, 0755)
		makeMainGo(directoryName)
		if err != nil {
			panic("directory creation error")
		}
	}
}

func makeMainGo(directoryName string) {
	// Read file
	filename := "main.go"
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("file read error: " + filename)
	}

	filename = directoryName + "main.go"
	if err != nil {
		panic("file write error: " + filename)
	}
	err = ioutil.WriteFile(filename, input, 0644)
}

func PrintUsage() {
	fmt.Println("Usage: go run create-directories.go <section-name> <number-of-directories>")
}
