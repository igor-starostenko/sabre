package main

import (
	"fmt"
	"os"
)

type StandardError struct {
	name        string
	description string
}

func (e *StandardError) Error() string {
	return fmt.Sprintf("Error: (%s) %s", e.name, e.description)
}

func fetchArg(i int, name string) (string, error) {
	if i <= 0 {
		return "", &StandardError{name, "Index not specified."}
	}
	if i >= len(os.Args) {
		return "", &StandardError{name, "Missing argument."}
	}
	return os.Args[i], nil
}

func readFile(fileName string) {
	fmt.Printf("Reading file: %s", fileName)
}

func main() {
	fileName, err := fetchArg(1, "fileName")
	if err != nil {
		fmt.Println(err)
		return
	}

	// outputDir, err := fetchArg(1, "outputDir")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	readFile(fileName)
}
