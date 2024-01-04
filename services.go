package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func analyzeStdIn(reader *bufio.Reader) string {
	var numberOfLines int64
	var numberOfWords int64
	var numberOfBytes int64
	var numberOfChars int64
	for {
		input, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading from stdin: %v\n", err)
			break
		}

		numberOfLines += 1
		numberOfBytes += getNumberOfBytes(input)

		numberOfWords += getNumberOfWordsFromBytes(input)
		numberOfChars += getNumberOfCharsFromBytes(input)
	}

	return format(numberOfLines, numberOfBytes, numberOfWords, numberOfChars+1)
}
