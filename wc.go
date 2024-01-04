package main

import (
	"bufio"
	"fmt"
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
	var outputString string
	reader := bufio.NewReader(os.Stdin)

	// fmt.Println(fmt.Sprintf("%v", args))
	if len(args) == 1 {
		outputString = analyzeStdIn(reader)
		fmt.Println(outputString)
	} else if len(args) == 2 {
		arg1 := args[1]
		if arg1 == "help" || arg1 == "-help" || arg1 == "--help" {
			fmt.Print(gethelpMessage())
		} else {
			_, err := getFile(arg1)
			if err != nil {
				flag = arg1
				outputString = analyzeStdInWithFlag(reader, flag, filename)
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
	} else if len(args) == 3 {
		flag = args[1]
		filename = args[2]
		outputString = analyzeFlagWithFilename(flag, filename)

		fmt.Println(outputString)
	} else {
		fmt.Print(gethelpMessage())
	}
}
