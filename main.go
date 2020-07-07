package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"time"
)

var args struct {
	source       string
	outputDir    string
	outputFormat string
	lineLimit    int
	scan         bool
	quiet        bool
	version      bool
}

func parseArgs() {
	flag.StringVar(&args.outputFormat, "e", defaultOutputFormat, "Output file extension")
	flag.IntVar(&args.lineLimit, "l", defaultLineLimit, "Max lines sliced per file")
	flag.BoolVar(&args.quiet, "q", false, "Supress informational output")
	flag.BoolVar(&args.scan, "s", false, "Use Scan method. Has max 4096 byte buffer limit for a line.")
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

type lineWriter func(int, string)

func scanLines(file *os.File, parse lineWriter) {
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		lineText := scanner.Text()
		parse(i, lineText)
		i += 1
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		reason := fmt.Sprintf("Failed to parse line #%v.", i)
		Stop(reason, 1)
	}
}

func readLines(file *os.File, parse lineWriter) {
	reader := bufio.NewReader(file)
	var err error
	i := 0

	for {
		var buffer bytes.Buffer
		var l []byte
		var isPrefix bool

		for {
			l, isPrefix, err = reader.ReadLine()
			buffer.Write(l)

			// If we've reached the end of the line, stop reading.
			if !isPrefix {
				break
			}

			// If we're just at the EOF, break
			if err != nil {
				break
			}
		}

		if err == io.EOF {
			break
		}

		lineText := buffer.String()
		parse(i, lineText)
		i += 1
	}

	if err != io.EOF {
		fmt.Println(err)
		reason := fmt.Sprintf("Failed to parse line #%v.", i)
		Stop(reason, 1)
	}
}

func main() {
	parseArgs()

	t := time.Now()
	OutputFiles = []*os.File{}

	if !args.quiet {
		fmt.Printf("Reading file: %s\n", args.source)
	}

	file := openFile(args.source)
	defer file.Close()

	if args.scan {
		scanLines(file, writeLine)
	} else {
		readLines(file, writeLine)
	}

	if !args.quiet {
		fmt.Printf("Slicing complete: %d seconds elapsed\n", int(time.Since(t).Seconds()))
	}
}
