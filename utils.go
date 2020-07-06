package main

import (
	"fmt"
	"os"
)

func Stop(message string, code int) {
	fmt.Fprintf(os.Stderr, message+"\n")
	os.Exit(code)
}
