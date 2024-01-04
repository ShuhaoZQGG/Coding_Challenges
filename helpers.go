package main

import (
	"io/fs"
	"os"
	"strconv"
	"strings"
)

const root string = "."

var fileSystem fs.FS = os.DirFS(root)

func getFile(filename string) (fs.File, error) {
	return fileSystem.Open(filename)
}

func getBytesFromFile(filename string) (int64, error) {
	file, err := getFile(filename)
	if err != nil {
		return 0, err
	}
	stat, err := file.Stat()

	if err != nil {
		return 0, err
	}

	size := stat.Size()
	return size, nil
}

func getNumberOfBytes(content []byte) int64 {
	return int64(len(content))
}

func getLinesFromFile(filename string) (int64, error) {
	content, err := fs.ReadFile(fileSystem, filename)
	if err != nil {
		return 0, err
	}
	var counter int64 = 0
	for _, v := range content {
		if v == 10 {
			counter += 1
		}
	}
	return counter, nil
}

func getNumberOfWordsFromFile(filename string) (int64, error) {
	content, err := fs.ReadFile(fileSystem, filename)
	if err != nil {
		return 0, err
	}
	return getNumberOfWordsFromBytes(content), err
}

func getNumberOfWordsFromBytes(content []byte) int64 {
	contentInString := string(content)
	contentInStringSlice := strings.Fields(contentInString)
	return int64(len(contentInStringSlice))
}
func getNumberOfCharsFromFile(filename string) (int64, error) {
	content, err := fs.ReadFile(fileSystem, filename)
	if err != nil {
		return 0, err
	}

	return getNumberOfCharsFromBytes(content), nil
}

/**
*	Comparing len(string(content)) and len([]rune(string(content))):
* len(string(content)): This gives you the number of bytes in the string.
* If the string contains multibyte characters, this count will be higher than the actual number of characters.
*
* Example: For the string "Go语言", len(string(content)) returns 7 because "语言" are multibyte characters in UTF-8.
* len([]rune(string(content))): This converts the string into a slice of runes before counting.
*
* Each rune represents one Unicode character, regardless of how many bytes it takes up in the encoding.
* This gives you the actual number of characters in the string.
* Example: For the same string "Go语言", len([]rune(string(content))) returns 4, accurately reflecting the number of characters.
 */
func getNumberOfCharsFromBytes(content []byte) int64 {
	return int64(len([]rune(string(content))))
}

func getAllInformationFromFile(filename string) (string, error) {
	outputString := ""
	outputNumber, err := getBytesFromFile(filename)
	if err != nil {
		return "", err
	}
	outputString += strconv.FormatInt(outputNumber, 10) + " "
	outputNumber, err = getNumberOfWordsFromFile(filename)
	if err != nil {
		return "", err
	}
	outputString += strconv.FormatInt(outputNumber, 10) + " "
	outputNumber, err = getNumberOfCharsFromFile(filename)
	if err != nil {
		return "", err
	}
	outputString += strconv.FormatInt(outputNumber, 10) + " "
	outputString += filename
	return outputString, nil
}

func format(numbers ...int64) string {
	output := ""
	for _, v := range numbers {
		numberInString := strconv.FormatInt(v, 10)
		output += numberInString + " "
	}

	return output
}
