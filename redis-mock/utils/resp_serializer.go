// TODO: Move it under resp model

package utils

import (
	"errors"
	"fmt"

	"github.com/redis-mock/constants"
	"github.com/redis-mock/models"
)

func Serialize[T any](message T) (string, error) {
	serializedOutput := ""
	switch v := any(message).(type) {

	case models.SimpleString:
		serializedOutput = SerializeSimpleStrings(v.Message)
	case models.BulkString:
		serializedOutput = SerializeBulkStrings(v.Message)
	case models.SimpleError:
		serializedOutput = SerializeErrors(v.Message)
	case []any:
		serializedOutput = SerializeArrays(v)
	case int64:
		serializedOutput = SerializeIntegers(v)
	case int32:
		serializedOutput = SerializeIntegers(int64(v))
	case int16:
		serializedOutput = SerializeIntegers(int64(v))
	case int8:
		serializedOutput = SerializeIntegers(int64(v))
	case int:
		serializedOutput = SerializeIntegers(int64(v))
	case nil:
		serializedOutput = SerializeBulkStringsForNull()
	default:
		return "", errors.New("Message type is not available")
	}

	return serializedOutput, nil
}

func SerializeSimpleStrings(message string) string {
	prefix := constants.SIMPLE_STRINGS_PREFIX
	crlf := constants.CRLF

	return prefix + message + crlf
}

func SerializeBulkStrings(message string) string {
	prefix := constants.BULK_STRINGS_PREFIX
	length := len(message)
	crlf := constants.CRLF
	return fmt.Sprintf("%s%d%s%s%s", prefix, length, crlf, message, crlf)
}

func SerializeBulkStringsForNull() string {
	prefix := constants.BULK_STRINGS_PREFIX
	content := "-1"
	crlf := constants.CRLF

	return prefix + content + crlf
}

func SerializeErrors(message string) string {
	prefix := constants.ERRORS_PREFIX
	crlf := constants.CRLF

	return prefix + message + crlf
}

func SerializeIntegers(message int64) string {
	prefix := constants.INTEGERS_PREFIX

	crlf := constants.CRLF

	return fmt.Sprintf("%s%d%s", prefix, message, crlf)
}

func SerializeArrays(messages []any) string {
	prefix := constants.ARRAYS_PREFIX
	length := len(messages)
	crlf := constants.CRLF
	output := fmt.Sprintf("%s%v%s", prefix, length, crlf)
	for _, message := range messages {
		serializedOutput, _ := Serialize[any](message)
		output += serializedOutput
	}
	return output
}
