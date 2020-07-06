package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var args struct {
	source       string
	outputDir    string
	outputFormat string
	lineLimit    int
	quiet        bool
	version      bool
}

type StandardError struct {
	name        string
	description string
}

// Holds references of output files
var OutputFiles []*os.File

func parseArgs() {
	flag.StringVar(&args.outputFormat, "e", defaultOutputFormat, "output file extension")
	flag.IntVar(&args.lineLimit, "l", defaultLineLimit, "Max lines sliced per file")
	flag.BoolVar(&args.quiet, "q", false, "Supress informational output")
	flag.BoolVar(&args.version, "v", false, "Print version info about sabre and exit")
	// flag.IntVar(&args.workers, "w", config.Workers, "# of workers")
	flag.Usage = usage
	flag.Parse()

	if args.version {
		Stop(version, 0)
	}

	if len(flag.Args()) <= 1 {
		usage()
		os.Exit(0)
	}

	args.source = flag.Args()[0]
	args.outputDir = flag.Args()[1]
}

func usage() {
	fmt.Println("usage: sabre [options] SOURCE OUTPUT")
	fmt.Println()
	flag.PrintDefaults()
}

func openFile(fileName string) *os.File {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func getOutputFile() *os.File {
	return OutputFiles[len(OutputFiles)-1]
}

func writeLine(i int, lineText string) {
	// fmt.Println(i, lineText)
	file := getOutputFile()

	line := fmt.Sprintf("%v  %s\n", i, lineText)
	_, err := file.WriteString(line)
	if err != nil {
		log.Fatal(err)
		file.Close()
		return
	}

	// log.Println(l, "bytes written successfully")
	// err = file.Close()
	// if err != nil {
	// 	log.Fatal(err)
	// 	file.Close()
	// 	return
	// }
}

type lineWriter func(int, string)

func parseLines(file *os.File, parse lineWriter) {
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		i += 1
		lineText := scanner.Text()
		parse(i, lineText)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	parseArgs()

	log.Printf("Reading file: %s", args.source)
	file := openFile(args.source)
	defer file.Close()

	log.Println(args)
	fileName := fmt.Sprintf("%s.%s", args.outputDir, args.outputFormat)
	outputFile, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	OutputFiles = []*os.File{outputFile}

	parseLines(file, writeLine)
}
