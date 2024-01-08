package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/redis-mock/constants"
)

func Deserialize(message string) (any, error) {
	var result any
	prefix := message[0:1]
	switch prefix {
	case constants.SIMPLE_STRINGS_PREFIX:
		result = DeserializeSimpleStrings(message)
	case constants.BULK_STRINGS_PREFIX:
		result = DeserializeBulkStrings(message)
	case constants.INTEGERS_PREFIX:
		result = DeserializeIntegers(message)
	case constants.ERRORS_PREFIX:
		result = DeserializeErrors(message)
	case constants.ARRAYS_PREFIX:
		result = DeserializeArrays(message)
	default:
		return nil, errors.New("Unable to deserialize, format unkown")
	}
	return result, nil
}

func DeserializeSimpleStrings(message string) string {
	crlf := constants.CRLF
	message = strings.TrimSuffix(message, crlf)
	result := message[1:]
	return result
}

func DeserializeBulkStrings(message string) string {
	crlf := constants.CRLF
	messageInSlice := strings.Split(message, crlf)
	result := messageInSlice[1]
	return result
}

func DeserializeIntegers(message string) int64 {
	crlf := constants.CRLF
	message = strings.TrimSuffix(message, crlf)
	result, _ := strconv.Atoi(message[1:])
	return int64(result)
}

func DeserializeErrors(message string) string {
	return DeserializeSimpleStrings(message)
}

func DeserializeArrays(message string) []any {
	if message == "*0\r\n" {
		return nil
	}
	crlf := constants.CRLF
	messageInSlices := strings.Split(message, crlf)

	output := make([]any, 0)
	counter := 1
	for counter < len(messageInSlices) {
		if len(messageInSlices[counter]) == 0 {
			return output
		}

		eachString := messageInSlices[counter]
		firstElementInSlice := string(eachString[0])

		if firstElementInSlice == constants.BULK_STRINGS_PREFIX {
			// Ensure there is a next element in the slice
			if counter+1 >= len(messageInSlices) {
				return output
			}

			bulkString := eachString + crlf + messageInSlices[counter+1] + crlf
			bulkStringInOutput, err := Deserialize(bulkString)
			if err != nil {
				fmt.Println("Error deserializing bulk string:", err)
				return nil
			}
			output = append(output, bulkStringInOutput)
			counter += 2
		} else {
			eachElementInOutput, err := Deserialize(eachString + crlf)
			if err != nil {
				fmt.Println("Error deserializing element:", err)
				return nil
			}
			output = append(output, eachElementInOutput)
			counter += 1
		}
	}

	return output
}
