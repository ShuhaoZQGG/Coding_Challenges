package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	/**
	* TODO: Step One: flag -c outputs the number of bytes in a file
	* ccwc -c test.text
	* 342190 text.txt
	 */
	args := os.Args
	var flag string
	var filename string
	var outputNumber int64
	var err error
	var outputString string
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println(fmt.Sprintf("%v", args))
	if len(args) == 1 {
		outputString = analyzeStdIn(reader)
		fmt.Println(outputString)
	} else if len(args) == 2 {
		arg1 := args[1]
		if arg1 == "help" || arg1 == "-help" || arg1 == "--help" {
			gethelpMessage()
		} else {
			_, err := getFile(arg1)
			if err != nil {
				flag = arg1

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
						gethelpMessage()
						return
					}
				}

				outputString = fmt.Sprintf("%v %v", outputNumber, filename)

				fmt.Println(outputString)
			} else {
				filename = arg1
				outputString, err = getAllInformationFromFile(filename)
				if err != nil {
					panic(fmt.Sprintf("error get bytes from file %s with error %v", filename, err))
				}
				fmt.Println(outputString)
			}
		}
	} else {
		flag = args[1]
		filename = args[2]

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
			gethelpMessage()
			break
		}

		if err != nil {
			panic(fmt.Sprintf("error get bytes from file %s with error %v", filename, err))
		}
		outputString = fmt.Sprintf("%v %v", outputNumber, filename)
		fmt.Println(outputString)
	}
}
