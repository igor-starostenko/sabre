package main

import (
	"fmt"
	"os"
)

type ArgumentError struct {
	name        string
	description string
}

func (e *ArgumentError) Error() string {
	return fmt.Sprintf("ArgumentError: (%s) %s", e.name, e.description)
}

func fetchArg(i int, name string) (string, error) {
	if i <= 0 {
		return "", &ArgumentError{name, "Index not specified."}
	}
	if i >= len(os.Args) {
		return "", &ArgumentError{name, "Missing argument."}
	}
	return os.Args[i], nil
}

func main() {
	fileName, err := fetchArg(1, "fileName")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Reading file: %s", fileName)
	}
}
