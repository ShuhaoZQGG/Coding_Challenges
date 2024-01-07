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

func analyzeStdInWithFlag(reader *bufio.Reader, flag string, filename string) string {
	var outputNumber int64
	var outputString string
	for {
		input, err := reader.ReadBytes('\n') // Reads until the newline or EOF
		if err == io.EOF {
			if flag == "-m" {
				outputNumber += 1
			}
			fmt.Println("reaching end of the file")
			break // Exit the loop on EOF
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			break
		}
		switch flag {
		case "-l":
			outputNumber += 1
			break
		case "-c":
			outputNumber += getNumberOfBytes(input)
		case "-w":
			outputNumber += getNumberOfWordsFromBytes(input)
			break
		case "-m":
			outputNumber += getNumberOfCharsFromBytes(input)
			break
		default:
			return gethelpMessage()
		}
	}

	outputString = fmt.Sprintf("%v %v", outputNumber, filename)
	return outputString
}

func analyzeFlagWithFilename(flag string, filename string) (outputString string) {
	var outputNumber int64
	var err error
	switch flag {
	case "-c":
		outputNumber, err = getBytesFromFile(filename)
		break
	case "-l":
		outputNumber, err = getLinesFromFile(filename)
		break
	case "-w":
		outputNumber, err = getNumberOfWordsFromFile(filename)
		break
	case "-m":
		outputNumber, err = getNumberOfCharsFromFile(filename)
		break
	default:
		return gethelpMessage()
	}

	if err != nil {
		panic(fmt.Sprintf("error get bytes from file %s with error %v", filename, err))
	}
	outputString = fmt.Sprintf("%v %v", outputNumber, filename)
	return outputString
}
