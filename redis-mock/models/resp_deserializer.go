package models

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/redis-mock/constants"
)

type Deserializer interface {
	Deserialize(message string) (any, error)
	DeserializeSimpleStrings(message string) string
	DeserializeBulkStrings(message string) string
	DeserializeIntegers(message string) string
	DeserializeErrors(message string) string
	DeserializeArrays(message string) []string
}

type RespDeserializer struct {
}

func (respDeserializer *RespDeserializer) Deserialize(message string) (any, error) {
	var result any
	prefix := message[0:1]
	switch prefix {
	case "":
		result = respDeserializer.DeserializeArrays(message)
	case constants.SIMPLE_STRINGS_PREFIX:
		result = respDeserializer.DeserializeSimpleStrings(message)
	case constants.BULK_STRINGS_PREFIX:
		result = respDeserializer.DeserializeBulkStrings(message)
	case constants.INTEGERS_PREFIX:
		result = respDeserializer.DeserializeIntegers(message)
	case constants.ERRORS_PREFIX:
		result = respDeserializer.DeserializeErrors(message)
	case constants.ARRAYS_PREFIX:
		result = respDeserializer.DeserializeArrays(message)
	default:
		return nil, errors.New("Unable to deserialize, format unkown")
	}
	return result, nil
}

func (respDeserializer *RespDeserializer) DeserializeSimpleStrings(message string) string {
	crlf := constants.CRLF
	message = strings.TrimSuffix(message, crlf)
	result := message[1:]
	return result
}

func (respDeserializer *RespDeserializer) DeserializeBulkStrings(message string) string {
	crlf := constants.CRLF
	messageInSlice := strings.Split(message, crlf)
	result := messageInSlice[1]
	return result
}

func (respDeserializer *RespDeserializer) DeserializeIntegers(message string) string {
	crlf := constants.CRLF
	message = strings.TrimSuffix(message, crlf)
	result, _ := strconv.Atoi(message[1:])
	resultInString := strconv.Itoa(int(result))
	return resultInString
}

func (respDeserializer *RespDeserializer) DeserializeErrors(message string) string {
	return respDeserializer.DeserializeSimpleStrings(message)
}

func (respDeserializer *RespDeserializer) DeserializeArrays(message string) []string {
	if message == "*0\r\n" || message == "" {
		return []string{}
	}
	crlf := constants.CRLF
	messageInSlices := strings.Split(message, crlf)

	output := make([]string, 0)
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
			bulkStringInOutput := respDeserializer.DeserializeBulkStrings(bulkString)
			output = append(output, bulkStringInOutput)
			counter += 2
		} else {
			eachElementInOutput, err := respDeserializer.Deserialize(eachString + crlf)
			if err != nil {
				fmt.Println("Error deserializing element:", err)
				return nil
			}
			eachElementInOutputInString, _ := eachElementInOutput.(string)

			output = append(output, eachElementInOutputInString)
			counter += 1
		}
	}

	return output
}
