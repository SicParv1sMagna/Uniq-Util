package main

import (
	"fmt"
	"os"

	"uniq/uniq"
)

func main() {
	options, err := uniq.ParseOptions()
	if err != nil {
		switch err {
		case uniq.ErrTooMuchFlagArgs, uniq.ErrUncombinedParams:
			fmt.Fprintln(os.Stderr, "Error: ", err)
			os.Exit(1)
		default:
			fmt.Fprintln(os.Stderr, "Error reading flags: ", err)
			os.Exit(1)
		}
	}

	if options.InputFile != "" {
		inputFile, err := os.Open(options.InputFile)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening input file:", err)
			os.Exit(1)
		}
		defer inputFile.Close()
		os.Stdin = inputFile
	}

	if options.OutputFile != "" {
		outputFile, err := os.OpenFile(options.OutputFile, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error opening output file:", err)
			os.Exit(1)
		}
		defer outputFile.Close()
		os.Stdout = outputFile
	}

	uniq.Functional(options, os.Stdin, os.Stdout)
}
