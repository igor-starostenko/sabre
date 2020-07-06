package main

import (
	"fmt"
	"os"
)

// Holds references of output files
var OutputFiles []*os.File

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		reason := fmt.Sprintf("Failed to open file: %s", fileName)
		Stop(reason, 1)
	}

	return file
}

func createFile(index int) {
	fileName := fmt.Sprintf("%s_%v.%s", args.outputDir, index, args.outputFormat)
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		reason := fmt.Sprintf("Failed to create a file: %s", fileName)
		Stop(reason, 1)
	}
	OutputFiles = append(OutputFiles, outputFile)
}

func getOutputFile() *os.File {
	return OutputFiles[len(OutputFiles)-1]
}

func generateOutputFile(i int) *os.File {
	if i%args.lineLimit == 0 {
		createFile(i/args.lineLimit + 1)
	}
	return getOutputFile()
}

func handleCloseFile(file *os.File, i int) {
	if (i+1)%args.lineLimit == 0 {
		file.Close()
	}
}

func writeLine(i int, lineText string) {
	file := generateOutputFile(i)

	line := fmt.Sprintf("%s\n", lineText)
	_, err := file.WriteString(line)
	if err != nil {
		fmt.Println(err)
		file.Close()
		reason := fmt.Sprintf("Failed to write line #%v", i)
		Stop(reason, 1)
	}

	handleCloseFile(file, i)

	// log.Println(l, "bytes written successfully")
	// err = file.Close()
	// if err != nil {
	// 	fmt.Println(err)
	// 	file.Close()
	// 	return
	// }
}
